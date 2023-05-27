package crawler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cast"

	"github.com/rocky114/craftsman/internal/pkg/path"

	"github.com/rocky114/craftsman/internal/storage"
	"github.com/sirupsen/logrus"

	"github.com/gocolly/colly/v2"
)

type nanjingLigongUniversity struct {
	university
}

func init() {
	collection["4132010288"] = &nanjingLigongUniversity{university{
		name: "南京理工大学",
		code: "4132010288",
	}}
}

func (u *nanjingLigongUniversity) crawl(ctx context.Context) error {
	c := colly.NewCollector(colly.CacheDir(path.GetTmpPath()))

	paramCollector := c.Clone()
	detailCollector := c.Clone()

	c.OnHTML(`table[id=TABLE3] > tbody > tr:first-of-type`, func(element *colly.HTMLElement) {
		var years, provinces, colleges []string
		element.ForEach(`select[id=ctl00_ContentPlaceHolder1_DropDownList1] > option`, func(i int, element *colly.HTMLElement) {
			years = append(years, element.Attr("value"))
		})
		element.ForEach(`select[id=ctl00_ContentPlaceHolder1_DropDownList2] > option`, func(i int, element *colly.HTMLElement) {
			provinces = append(provinces, element.Attr("value"))
		})

		element.ForEach(`select[id=ctl00_ContentPlaceHolder1_DropDownList3] > option`, func(i int, element *colly.HTMLElement) {
			colleges = append(colleges, element.Attr("value"))
		})

		/*if err := detailCollector.Visit("https://zsb.suda.edu.cn/search.aspx?nf=2022&sf=10&xy=1"); err != nil {
			logrus.Errorf("scrape suzhou university err: %v", err)
		}*/

		for _, year := range years {
			if !u.containAdmissionTime(year) {
				return
			}

			for _, province := range provinces {
				for _, college := range colleges {
					url := fmt.Sprintf("https://zsb.suda.edu.cn/search.aspx?nf=%s&sf=%s&xy=%s", year, province, college)
					if err := detailCollector.Visit(url); err != nil {
						logrus.Errorf("scrape suzhou university err: %v", err)
					}
				}
			}
		}
	})

	paramCollector.OnResponse(func(response *colly.Response) {
		if response.StatusCode != http.StatusOK {
			logrus.Errorf("university %s crawl code: %d", u.name, response.StatusCode)
			return
		}

		var params nanjingParamResp
		if err := json.Unmarshal(response.Body, &params); err != nil {
			logrus.Errorf("nanjing scrape unmarshal param err: %v", response.StatusCode)
			return
		}

		for province, items := range params.Data.SsmcNfKlmcSexCampusZslxMap {
			for _, item := range items {
				if !u.containAdmissionTime(item.Nf) {
					continue
				}

				req := map[string]string{
					"ssmc": province,
					"zsnf": item.Nf,
					"klmc": item.Klmc,
					"zslx": item.Zslx,
				}

				if err := detailCollector.Post("https://bkzs.nju.edu.cn/f/ajax_lnfs", req); err != nil {
					logrus.Errorf("nanjing scrape detail err: %v", response.StatusCode)
				}
			}
		}
	})

	detailCollector.OnResponse(func(response *colly.Response) {
		if response.StatusCode != http.StatusOK {
			logrus.Errorf("nanjing scrape detail response status code: %d", response.StatusCode)
			return
		}

		var params nanjingAdmissionScoreResp
		if err := json.Unmarshal(response.Body, &params); err != nil {
			logrus.Errorf("nanjing scrape unmarshal detail err: %v", response.StatusCode)
			return
		}

		for _, item := range params.Data.ZsSsgradeList {
			if err := storage.GetQueries().CreateAdmissionMajor(ctx, storage.CreateAdmissionMajorParams{
				University:    u.name,
				Province:      item.Ssmc,
				AdmissionType: item.Zslx,
				SelectExam:    item.Klmc,
				AdmissionTime: item.Nf,
				MinScore:      cast.ToString(item.MinScore),
			}); err != nil {
				logrus.Errorf("create admission major err: %v", err)
			}
		}
	})

	c.OnResponse(func(response *colly.Response) {
		if response.StatusCode != http.StatusOK {
			logrus.Errorf("nanjing scrape response status code: %d", response.StatusCode)
			return
		}

		if err := json.Unmarshal(response.Body, &nanjingLogin); err != nil {
			logrus.Errorf("nanjing scrape unmarshal login err: %v", response.StatusCode)
			return
		}

		if err := paramCollector.Post("https://bkzs.nju.edu.cn/f/ajax_lnfs_param", nil); err != nil {
			logrus.Errorf("nanjing scrape params err: %v", response.StatusCode)
		}
	})

	return c.Post("https://zs.nuaa.edu.cn/lnlqfs/list.psp", nil)
}

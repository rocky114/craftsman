package scraper

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/rocky114/craftsman/internal/storage"

	"github.com/sirupsen/logrus"

	"github.com/spf13/cast"

	"github.com/gocolly/colly/v2"
)

type nanjingLoginResp struct {
	State     uint   `json:"state"`
	Msg       string `json:"msg"`
	Data      string `json:"data"`
	Jessionid string `json:"jessionid"`
}

type nanjingParamResp struct {
	Data struct {
		SsmcNfKlmcSexCampusZslxMap map[string][]struct {
			Klmc string `json:"klmc"`
			Nf   string `json:"nf"`
			Zslx string `json:"zslx"`
			Ssmc string `json:"ssmc"`
		} `json:"ssmc_nf_klmc_sex_campus_zslx_Map"`
	} `json:"data"`
}

type nanjingAdmissionScoreResp struct {
	Data struct {
		ZsSsgradeList []struct {
			Klmc     string  `json:"klmc"`
			Nf       string  `json:"nf"`
			Ssmc     string  `json:"ssmc"`
			MinScore float32 `json:"minScore"`
		}
	}
}

var nanjingLogin nanjingLoginResp

func ScrapeAdmissionMajorScoreNanjing() error {
	c := colly.NewCollector(colly.CacheDir("./web"))
	c.OnRequest(func(request *colly.Request) {
		request.Headers.Set("X-Requested-Time", cast.ToString(time.Now().UnixMilli()))
		request.Headers.Set("X-Requested-With", "XMLHttpRequest")
	})

	paramCollector := c.Clone()
	paramCollector.OnRequest(func(request *colly.Request) {
		request.Headers.Set("X-Requested-Time", cast.ToString(time.Now().UnixMilli()))
		request.Headers.Set("X-Requested-With", "XMLHttpRequest")
		request.Headers.Set("Csrf-Token", nanjingLogin.Data)
		request.Headers.Set("Cookie", fmt.Sprintf("zhaosheng.nju.session.id=%s", nanjingLogin.Jessionid))
		request.Headers.Set("Referer", "https://bkzs.nju.edu.cn/static/front/nju/basic/html_web/lnfs.html")
	})

	detailCollector := c.Clone()
	detailCollector.OnRequest(func(request *colly.Request) {
		request.Headers.Set("X-Requested-Time", cast.ToString(time.Now().UnixMilli()))
		request.Headers.Set("X-Requested-With", "XMLHttpRequest")
		request.Headers.Set("Csrf-Token", nanjingLogin.Data)
		request.Headers.Set("Cookie", fmt.Sprintf("zhaosheng.nju.session.id=%s", nanjingLogin.Jessionid))
		request.Headers.Set("Referer", "https://bkzs.nju.edu.cn/static/front/nju/basic/html_web/lnfs.html")
	})

	paramCollector.OnResponse(func(response *colly.Response) {
		if response.StatusCode != http.StatusOK {
			logrus.Errorf("nanjing scrape params response status code: %d", response.StatusCode)
			return
		}

		var params nanjingParamResp
		if err := json.Unmarshal(response.Body, &params); err != nil {
			logrus.Errorf("nanjing scrape unmarshal param err: %v", response.StatusCode)
			return
		}

		for province, items := range params.Data.SsmcNfKlmcSexCampusZslxMap {
			for _, item := range items {
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
			if err := storage.GetQueries().CreateAdmissionMajor(context.Background(), storage.CreateAdmissionMajorParams{
				College:       "南京大学",
				Province:      item.Ssmc,
				SubjectType:   item.Klmc,
				AdmissionTime: item.Nf,
				MinScore:      cast.ToInt32(item.MinScore),
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

	return c.Post("https://bkzs.nju.edu.cn/f/ajax_get_csrfToken", nil)
}

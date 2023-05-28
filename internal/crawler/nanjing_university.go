package crawler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/rocky114/craftsman/internal/pkg/path"

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

type nanjingAdmissionPlanResp struct {
	Data struct {
		ZsjhList []struct {
			Nf     string `json:"nf"`     //年份
			Ssmc   string `json:"ssmc"`   //省份
			Zydhmc string `json:"zydhmc"` //专业
			Zylx   string `json:"zylx"`   //类型
			Klmc   string `json:"klmc"`   //选考
			Zsjhs  uint32 `json:"zsjhs"`  //计划数量
		} `json:"zsjhList"`
	} `json:"data"`
}

type nanjingAdmissionScoreResp struct {
	Data struct {
		ZsSsgradeList []struct {
			Klmc     string  `json:"klmc"`
			Zslx     string  `json:"zslx"`
			Nf       string  `json:"nf"`
			Ssmc     string  `json:"ssmc"`
			MinScore float32 `json:"minScore"`
		} `json:"zsssgradelist"`
	} `json:"data"`
}

var nanjingLogin nanjingLoginResp

func init() {
	collection["4132010284"] = &nanjingUniversity{university{
		code: "4132010284",
		name: "南京大学",
	}}
}

type nanjingUniversity struct {
	university
}

func (u *nanjingUniversity) crawl(ctx context.Context) error {
	c := colly.NewCollector(colly.CacheDir(path.GetTmpPath()))
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

	admissionPlanCollector := c.Clone()
	admissionPlanCollector.OnRequest(func(request *colly.Request) {
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

	scoreMap := make(map[string]string)

	paramCollector.OnResponse(func(response *colly.Response) {
		if response.StatusCode != http.StatusOK {
			logrus.Errorf("%s university http code: %d", u.name, response.StatusCode)
			return
		}

		var params nanjingParamResp
		if err := json.Unmarshal(response.Body, &params); err != nil {
			logrus.Errorf("nanjingParamResp err: %v", response.StatusCode)
			return
		}

		// admission score
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
					logrus.Errorf("%s detailCollector err: %v", u.name, err)
				}
			}
		}

		// admission plan
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

				if err := admissionPlanCollector.Post("https://bkzs.nju.edu.cn/f/ajax_zsjh", req); err != nil {
					logrus.Errorf("%s admissionPlanCollector err: %v", u.name, err)
				}
			}
		}
	})

	admissionPlanCollector.OnResponse(func(response *colly.Response) {
		if response.StatusCode != http.StatusOK {
			logrus.Errorf("%s university http code: %d", u.name, response.StatusCode)
			return
		}

		var admissionPlan nanjingAdmissionPlanResp
		if err := json.Unmarshal(response.Body, &admissionPlan); err != nil {
			logrus.Errorf("nanjingAdmissionPlanResp unmarshal err: %v", err)
			return
		}

		for _, item := range admissionPlan.Data.ZsjhList {
			minScore := scoreMap[fmt.Sprintf("%s%s", item.Ssmc, item.Klmc)]

			if err := storage.GetQueries().CreateAdmissionMajor(ctx, storage.CreateAdmissionMajorParams{
				University:      u.name,
				Province:        item.Ssmc,
				Major:           item.Zydhmc,
				AdmissionType:   item.Zylx,
				SelectExam:      item.Klmc,
				AdmissionTime:   item.Nf,
				AdmissionNumber: cast.ToString(item.Zsjhs),
				MinScore:        minScore,
			}); err != nil {
				logrus.Errorf("create admission major err: %v", err)
			}
		}
	})

	detailCollector.OnResponse(func(response *colly.Response) {
		if response.StatusCode != http.StatusOK {
			logrus.Errorf("%s university http code: %d", u.name, response.StatusCode)
			return
		}

		var params nanjingAdmissionScoreResp
		if err := json.Unmarshal(response.Body, &params); err != nil {
			logrus.Errorf("nanjing scrape unmarshal detail err: %v", response.StatusCode)
			return
		}

		for _, item := range params.Data.ZsSsgradeList {
			scoreMap[fmt.Sprintf("%s%s", item.Ssmc, item.Klmc)] = cast.ToString(item.MinScore)
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

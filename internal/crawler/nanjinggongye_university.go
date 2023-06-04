package crawler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/rocky114/craftsman/internal/pkg/path"

	"github.com/rocky114/craftsman/internal/storage"
	"github.com/sirupsen/logrus"

	"github.com/gocolly/colly/v2"
)

type nanjinggongyeUniversity struct {
	university
}

func init() {
	collection["4132010291"] = &nanjinggongyeUniversity{university{
		name: "南京工业大学",
		code: "4132010291",
	}}
}

type nanjinggongyeUniversityResp struct {
	Rows []struct {
		NF   string `json:"NF"`
		SFMC string `json:"SFMC"`
		KL   string `json:"KL"`
		PC   string `json:"PC"`
		LQQK string `json:"LQQK"`
		SX   string `json:"SX"`
		XX   string `json:"XX"`
	} `json:"rows"`
}

func (u *nanjinggongyeUniversity) crawl(ctx context.Context) error {
	c := colly.NewCollector(colly.CacheDir(path.GetTmpPath()))

	c.OnResponse(func(response *colly.Response) {
		if response.StatusCode != http.StatusOK {
			logrus.Errorf("nanjinggongyeUniversity http code: %d", response.StatusCode)
			return
		}

		var resp nanjinggongyeUniversityResp
		if err := json.Unmarshal(response.Body, &resp); err != nil {
			logrus.Errorf("nanjinggongyeUniversity unmarshal err: %v", response.StatusCode)
			return
		}

		for _, item := range resp.Rows {
			if err := storage.GetQueries().CreateAdmissionMajor(ctx, storage.CreateAdmissionMajorParams{
				University:               u.name,
				Province:                 item.SFMC,
				AdmissionType:            item.PC,
				AdmissionTime:            u.admissionTime,
				AdmissionNumber:          item.LQQK,
				SelectExam:               item.KL,
				MinScore:                 item.XX,
				ProvinceControlScoreLine: item.SX,
			}); err != nil {
				logrus.Errorf("create admission major err: %v", err)
			}
		}
	})

	params := map[string]string{
		"limit":           "200",
		"offset":          "0",
		"sortOrder":       "desc",
		"chaxun1nianfen":  u.admissionTime,
		"chaxun1shengfen": "",
	}
	return c.Post("https://zhaosheng.njtech.edu.cn/index/Menu.ashx?acdo=chaxunlnlist", params)
}

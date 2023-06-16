package crawler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/rocky114/craftsman/internal/storage"
	"github.com/sirupsen/logrus"

	"github.com/rocky114/craftsman/internal/pkg/path"

	"github.com/gocolly/colly/v2"
)

type huaiyinshifanUniversity struct {
	university
}

func init() {
	collection["4132010323"] = &huaiyinshifanUniversity{university{
		name: "淮阴师范学院",
		code: "4132010323",
	}}
}

type huaiyinshifanUniversityResp struct {
	Zymc string `json:"zymc"`
	Pc   string `json:"pc"`
	Zgf  string `json:"zgf"`
	Zdf  string `json:"zdf"`
	Pjf  string `json:"pjf"`
	Ss   string `json:"ss"`
	Nf   string `json:"nf"`
}

var huaiyinshifanUniversityTokens = []string{
	"5rGf6IuP",
	"5YyX5Lqs",
	"5aSp5rSl",
	"5rKz5YyX",
	"5bGx6KW/",
	"5YaF6JKZ5Y+k",
	"6L695a6B",
	"5ZCJ5p6X",
	"6buR6b6Z5rGf",
	"5LiK5rW3",
	"5rWZ5rGf",
	"5a6J5b69",
	"56aP5bu6",
	"5rGf6KW/",
	"5bGx5Lic",
	"5rKz5Y2X",
	"5rmW5YyX",
	"5rmW5Y2X",
	"5bm/5Lic",
	"5bm/6KW/",
	"5rW35Y2X",
	"6YeN5bqG",
	"5Zub5bed",
	"6LS15bee",
	"5LqR5Y2X",
	"6KW/6JeP",
	"6ZmV6KW/",
	"55SY6IKD",
	"6Z2S5rW3",
	"5a6B5aSP",
	"5paw55aG",
}

func (u *huaiyinshifanUniversity) crawl(ctx context.Context) error {
	c := colly.NewCollector(colly.CacheDir(path.GetTmpPath()))

	c.OnResponse(func(response *colly.Response) {
		if response.StatusCode != http.StatusOK {
			logrus.Errorf("huaiyinshifanUniversity http code: %d", response.StatusCode)
			return
		}

		var resp []huaiyinshifanUniversityResp
		if err := json.Unmarshal(response.Body, &resp); err != nil {
			logrus.Errorf("huaiyinshifanUniversity unmarshal err: %v", err)
			return
		}

		for _, item := range resp {
			if err := storage.GetQueries().CreateAdmissionMajor(ctx, storage.CreateAdmissionMajorParams{
				University:    u.name,
				Province:      item.Ss,
				Major:         item.Zymc,
				MaxScore:      item.Zgf,
				MinScore:      item.Zdf,
				AverageScore:  item.Pjf,
				AdmissionTime: item.Nf,
				AdmissionType: item.Pc,
			}); err != nil {
				logrus.Errorf("create admission major err: %v", err)
			}
		}
	})

	for _, token := range huaiyinshifanUniversityTokens {
		params := map[string]string{
			"ss":         token,
			"nf":         u.admissionTime,
			"klxz":       "5pmu6YCa57G7",
			"templateid": "3",
			"owner":      "1583930167",
		}
		if err := c.Post("https://zb.hytc.edu.cn/system/resource/importdata/getlnlq.jsp", params); err != nil {
			logrus.Errorf("huaiyinshifanUniversity err: %v", err)
		}
	}

	return nil
}

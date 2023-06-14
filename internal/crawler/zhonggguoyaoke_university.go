package crawler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rocky114/craftsman/internal/types"

	"github.com/rocky114/craftsman/internal/storage"
	"github.com/sirupsen/logrus"

	"github.com/rocky114/craftsman/internal/pkg/path"

	"github.com/gocolly/colly/v2"
)

type zhongguoyaokeUniversity struct {
	university
}

func init() {
	collection["4132010316"] = &zhongguoyaokeUniversity{university{
		name: "中国药科大学",
		code: "4132010316",
	}}
}

type zhongguoyaokeUniversityResp struct {
	Data []struct {
		F1 string `json:"f1"`
		F2 string `json:"f2"`
		F3 string `json:"f3"`
		F4 string `json:"f4"`
		F5 string `json:"f5"`
		F8 string `json:"f8"`
	} `json:"data"`
}

func (u *zhongguoyaokeUniversity) crawl(ctx context.Context) error {
	c := colly.NewCollector(colly.CacheDir(path.GetTmpPath()))

	c.OnResponse(func(response *colly.Response) {
		if response.StatusCode != http.StatusOK {
			logrus.Errorf("zhongguoyaokeUniversity http code: %d", response.StatusCode)
			return
		}

		var resp zhongguoyaokeUniversityResp
		if err := json.Unmarshal(response.Body, &resp); err != nil {
			logrus.Errorf("zhongguoyaokeUniversity unmarshal err: %v", err)
			return
		}

		for _, item := range resp.Data {
			if err := storage.GetQueries().CreateAdmissionMajor(ctx, storage.CreateAdmissionMajorParams{
				University:    u.name,
				Province:      item.F8,
				Major:         item.F1,
				SelectExam:    item.F2,
				MaxScore:      item.F3,
				MinScore:      item.F4,
				AverageScore:  item.F5,
				AdmissionTime: u.admissionTime,
			}); err != nil {
				logrus.Errorf("create admission major err: %v", err)
			}
		}
	})

	for _, province := range types.Provinces {
		params := map[string]string{
			"returnInfos": `[{"name":"f1"},{"name":"f2"},{"name":"f3"},{"name":"f4"},{"name":"f5"},{"name":"f8"}]`,
			"rows":        "50",
			"pageIndex":   "1",
			"columnId":    "10209",
			"siteId":      "225",
			"conditions":  fmt.Sprintf(`[{"orConditions":[{"field":"f8","value":"%s","judge":"="}]},{"orConditions":[{"field":"f7","value":"%s","judge":"="}]}]`, province, u.admissionTime),
		}
		if err := c.Post("http://zb.cpu.edu.cn/_wp3services/generalQuery?queryObj=articles", params); err != nil {
			logrus.Errorf("zhongguoyaokeUniversity err: %v", err)
		}
	}

	return nil
}

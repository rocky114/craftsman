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

type jiangsuligongUniversity struct {
	university
}

func init() {
	collection["4132011463"] = &jiangsuligongUniversity{university{
		name: "江苏理工学院",
		code: "4132011463",
	}}
}

type jiangsuligongUniversityResp struct {
	Data []struct {
		Title string `json:"title"`
		F2    string `json:"f2"`
		F4    string `json:"f4"`
		F5    string `json:"f5"`
		F6    string `json:"f6"`
		F7    string `json:"f7"`
		F8    string `json:"f8"`
		F9    string `json:"f9"`
	} `json:"data"`
}

func (u *jiangsuligongUniversity) crawl(ctx context.Context) error {
	c := colly.NewCollector(colly.CacheDir(path.GetTmpPath()))

	c.OnResponse(func(response *colly.Response) {
		if response.StatusCode != http.StatusOK {
			logrus.Errorf("jiangsuligongUniversity http code: %d", response.StatusCode)
			return
		}

		var resp jiangsuligongUniversityResp
		if err := json.Unmarshal(response.Body, &resp); err != nil {
			logrus.Errorf("jiangsuligongUniversity unmarshal err: %v", err)
			return
		}

		for _, item := range resp.Data {
			if err := storage.GetQueries().CreateAdmissionMajor(ctx, storage.CreateAdmissionMajorParams{
				University:    u.name,
				Province:      item.F9,
				Major:         item.Title,
				SelectExam:    item.F7,
				MaxScore:      item.F4,
				MinScore:      item.F5,
				AverageScore:  item.F6,
				AdmissionTime: item.F8,
				AdmissionType: item.F2,
			}); err != nil {
				logrus.Errorf("create admission major err: %v", err)
			}
		}
	})

	for _, province := range types.Provinces {
		province = "江苏"
		params := map[string]string{
			"returnInfos": `[{"field":"title","name":"title"},{"field":"f2","name":"f2"},{"field":"f4","name":"f4"},{"field":"f5","name":"f5"},{"field":"f6","name":"f6"},{"field":"f7","name":"f7"},{"field":"f8","name":"f8"},{"field":"f9","name":"f9"}]`,
			"rows":        "200",
			"pageIndex":   "1",
			"columnId":    "6021",
			"siteId":      "17",
			"conditions":  fmt.Sprintf(`[{"field":"scope","value":1,"judge":"="},{"field":"f8","value":"%s","judge":"="},{"field":"f9","value":"%s","judge":"like"}]`, u.admissionTime, province),
		}
		if err := c.Post("http://zs.jstu.edu.cn/_wp3services/generalQuery?queryObj=articles", params); err != nil {
			logrus.Errorf("jiangsuligongUniversity err: %v", err)
		}

		break
	}

	return nil
}

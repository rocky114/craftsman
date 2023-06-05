package crawler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/rocky114/craftsman/internal/types"
	"github.com/spf13/cast"

	"github.com/rocky114/craftsman/internal/storage"
	"github.com/sirupsen/logrus"

	"github.com/rocky114/craftsman/internal/pkg/path"

	"github.com/gocolly/colly/v2"
)

type hehaiUniversity struct {
	university
}

func init() {
	collection["4132010294"] = &hehaiUniversity{university{
		name: "河海大学",
		code: "4132010294",
	}}
}

type hehaiUniversityResp struct {
	Data []struct {
		Province     string `json:"province"`
		Type         string `json:"type"`
		Discipline   string `json:"discipline"`
		Major        string `json:"major"`
		FileScore    int    `json:"filescore"`
		Highestscore int    `json:"highestscore"`
		Red          string `json:"red"`
	} `json:"data"`
}

func (u *hehaiUniversity) crawl(ctx context.Context) error {
	c := colly.NewCollector(colly.CacheDir(path.GetTmpPath()))

	c.OnResponse(func(response *colly.Response) {
		if response.StatusCode != http.StatusOK {
			logrus.Errorf("hehaiUniversity http code: %d", response.StatusCode)
			return
		}

		var resp hehaiUniversityResp
		if err := json.Unmarshal(response.Body, &resp); err != nil {
			logrus.Errorf("hehaiUniversity unmarshal err: %v", err)
			return
		}

		for _, item := range resp.Data {
			if item.Red != "" {
				continue
			}

			if err := storage.GetQueries().CreateAdmissionMajor(ctx, storage.CreateAdmissionMajorParams{
				University:    u.name,
				Province:      item.Province,
				Major:         item.Major,
				AdmissionType: item.Type,
				SelectExam:    item.Discipline,
				MaxScore:      cast.ToString(item.Highestscore),
				MinScore:      cast.ToString(item.FileScore),
				AdmissionTime: u.admissionTime,
			}); err != nil {
				logrus.Errorf("create admission major err: %v", err)
			}
		}
	})

	for _, province := range types.Provinces {
		params := map[string]string{"year": u.admissionTime, "province": province}
		if err := c.Post("https://zsw.hhu.edu.cn/api/lsfs/fsList", params); err != nil {
			logrus.Errorf("hehaiUniversity err:%v", err)
		}

		logrus.Infof("hehaiUniversity crawl %s", province)

		time.Sleep(10 * time.Second)
	}

	return nil
}

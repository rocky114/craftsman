package crawler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rocky114/craftsman/internal/storage"
	"github.com/sirupsen/logrus"

	"github.com/rocky114/craftsman/internal/pkg/path"

	"github.com/gocolly/colly/v2"
)

type nanjinghangtianhangkongUniversity struct {
	university
}

func init() {
	collection["4132010287"] = &nanjinghangtianhangkongUniversity{university{
		name: "南京航空航天大学",
		code: "4132010287",
	}}
}

type nanjinghangtianhangkongUniversityResp struct {
	Data []struct {
		Year         string `json:"year"`
		Province     string `json:"province"`
		Type         string `json:"type"`
		Specialty    string `json:"specialty"`
		College      string `json:"college"`
		HighestScore string `json:"highestScore"`
		LowestScore  string `json:"lowestScore"`
		Subject      string `json:"subject"`
	} `json:"data"`
}

func (u *nanjinghangtianhangkongUniversity) crawl(ctx context.Context) error {
	c := colly.NewCollector(colly.CacheDir(path.GetTmpPath()))

	c.OnResponse(func(response *colly.Response) {
		if response.StatusCode != http.StatusOK {
			logrus.Errorf("nanjinghangtianhangkongUniversity http code: %d", response.StatusCode)
			return
		}

		var resp nanjinghangtianhangkongUniversityResp
		if err := json.Unmarshal(response.Body, &resp); err != nil {
			logrus.Errorf("nanjinghangtianhangkongUniversity unmarshal err: %v", err)
			return
		}

		for _, item := range resp.Data {
			if err := storage.GetQueries().CreateAdmissionMajor(ctx, storage.CreateAdmissionMajorParams{
				University:    u.name,
				College:       item.College,
				Province:      item.Province,
				Major:         item.Specialty,
				AdmissionType: item.Type,
				SelectExam:    item.Subject,
				MaxScore:      item.HighestScore,
				MinScore:      item.LowestScore,
				AdmissionTime: item.Year,
			}); err != nil {
				logrus.Errorf("create admission major err: %v", err)
			}
		}
	})

	return c.Visit(fmt.Sprintf("https://zsservice.nuaa.edu.cn/zsw-admin/api/getAdmissionScore?year=%s", u.admissionTime))
}

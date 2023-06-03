package crawler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/rocky114/craftsman/internal/types"

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

// todo: cann't open page
type nanjinggongyeUniversityResp struct {
	Rows []struct {
		Province         string `json:"province"`
		ProfessionalName string `json:"professional_name"`
		ClassName        string `json:"class_name"`
		Year3            string `json:"year3"`
	} `json:"rows"`
}

func (u *nanjinggongyeUniversity) crawl(ctx context.Context) error {
	c := colly.NewCollector(colly.CacheDir(path.GetTmpPath()))

	detailCollector := c.Clone()

	c.OnHTML(`div.gyright h3[id=title]`, func(element *colly.HTMLElement) {
		title := []rune(strings.TrimSpace(element.Text))
		admissionTime := string(title[6:10])
		if !u.containAdmissionTime(admissionTime) {
			return
		}

		for _, province := range types.Provinces {
			addr := fmt.Sprintf("http://zsb.njust.edu.cn/lqScore/initDateWebCon?pageSize=100&rowoffset=0&val1=%s", url.QueryEscape(province))
			if err := detailCollector.Visit(addr); err != nil {
				logrus.Errorf("nanjingLigongUniversity admission score err: %v", err)
			}
		}
	})

	detailCollector.OnResponse(func(response *colly.Response) {
		if response.StatusCode != http.StatusOK {
			logrus.Errorf("nanjingLigongUniversity http code: %d", response.StatusCode)
			return
		}

		var resp nanjingligongUniversityResp
		if err := json.Unmarshal(response.Body, &resp); err != nil {
			logrus.Errorf("nanjingLigongUniversity unmarshal detail err: %v", response.StatusCode)
			return
		}

		for _, item := range resp.Rows {
			if err := storage.GetQueries().CreateAdmissionMajor(ctx, storage.CreateAdmissionMajorParams{
				University:    u.name,
				Province:      item.Province,
				Major:         item.ProfessionalName,
				AdmissionType: item.ClassName,
				AdmissionTime: u.admissionTime,
				MinScore:      item.Year3,
			}); err != nil {
				logrus.Errorf("create admission major err: %v", err)
			}
		}
	})

	return c.Visit("http://zsb.njust.edu.cn/lqjh_fsx")
}

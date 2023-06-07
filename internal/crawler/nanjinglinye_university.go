package crawler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/rocky114/craftsman/internal/storage"
	"github.com/rocky114/craftsman/internal/types"
	"github.com/sirupsen/logrus"

	"github.com/rocky114/craftsman/internal/pkg/path"

	"github.com/gocolly/colly/v2"
)

type nanjinglinyeUniversity struct {
	university
}

func init() {
	collection["4132010298"] = &nanjinglinyeUniversity{university{
		name: "南京林业大学",
		code: "4132010298",
	}}
}

type nanjinglinyeUniversityResp struct {
	Data []struct {
		ProvinceName  string `json:"provinceName"`
		Batch         string `json:"batch"`
		Kind          string `json:"kind"`
		MajorName     string `json:"majorName"`
		MaxScore      string `json:"maxScore"`
		MinScore      string `json:"minScore"`
		ProvinceScore string `json:"provinceScore"`
	} `json:"data"`
}

func (u *nanjinglinyeUniversity) crawl(ctx context.Context) error {
	c := colly.NewCollector(colly.CacheDir(path.GetTmpPath()))

	c.OnResponse(func(response *colly.Response) {
		if response.StatusCode != http.StatusOK {
			logrus.Errorf("nanjinglinyeUniversity http code: %d", response.StatusCode)
			return
		}

		var resp nanjinglinyeUniversityResp
		if err := json.Unmarshal(response.Body, &resp); err != nil {
			logrus.Errorf("nanjinglinyeUniversity unmarshal err: %v", err)
			return
		}

		for _, item := range resp.Data {
			if err := storage.GetQueries().CreateAdmissionMajor(ctx, storage.CreateAdmissionMajorParams{
				University:               u.name,
				Province:                 u.trimProvinceSuffix(item.ProvinceName),
				Major:                    item.MajorName,
				AdmissionType:            item.Batch,
				SelectExam:               item.Kind,
				MaxScore:                 item.MaxScore,
				MinScore:                 item.MinScore,
				AdmissionTime:            u.admissionTime,
				ProvinceControlScoreLine: item.ProvinceScore,
			}); err != nil {
				logrus.Errorf("create admission major err: %v", err)
			}
		}
	})

	for _, province := range types.Provinces {
		addr := fmt.Sprintf("https://lqcx.njfu.edu.cn/enrollment/open/scoreList?year=%s&provinceName=%s&kind=", u.admissionTime, url.QueryEscape(province))
		if err := c.Visit(addr); err != nil {
			logrus.Errorf("nanjinglinyeUniversity err:%v", err)
			continue
		}

		logrus.Infof("nanjinglinyeUniversity crawl %s", province)

		time.Sleep(5 * time.Second)
	}

	return nil
}

func (u *nanjinglinyeUniversity) trimProvinceSuffix(province string) string {
	for _, item := range []string{"壮族自治区", "回族自治区", "维吾尔自治区", "省", "市", "自治区"} {
		province = strings.TrimSuffix(province, item)
	}

	return province
}

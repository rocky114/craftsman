package crawler

import (
	"context"
	"fmt"
	"strings"

	"github.com/rocky114/craftsman/internal/pkg/path"

	"github.com/rocky114/craftsman/internal/storage"
	"github.com/sirupsen/logrus"

	"github.com/gocolly/colly/v2"
)

type suzhouUniversity struct {
	university
}

func init() {
	collection["4132010285"] = &suzhouUniversity{university{
		name: "苏州大学",
		code: "4132010285",
	}}
}

func (u *suzhouUniversity) crawl(ctx context.Context) error {
	c := colly.NewCollector(colly.CacheDir(path.GetTmpPath("suzhouUniversity")))

	detailCollector := c.Clone()

	c.OnHTML(`table[id=TABLE3] > tbody > tr:first-of-type`, func(element *colly.HTMLElement) {
		var years, provinces, colleges []string
		element.ForEach(`select[id=ctl00_ContentPlaceHolder1_DropDownList1] > option`, func(i int, element *colly.HTMLElement) {
			years = append(years, element.Attr("value"))
		})
		element.ForEach(`select[id=ctl00_ContentPlaceHolder1_DropDownList2] > option`, func(i int, element *colly.HTMLElement) {
			provinces = append(provinces, element.Attr("value"))
		})

		element.ForEach(`select[id=ctl00_ContentPlaceHolder1_DropDownList3] > option`, func(i int, element *colly.HTMLElement) {
			colleges = append(colleges, element.Attr("value"))
		})

		for _, year := range years {
			if !u.containAdmissionTime(year) {
				return
			}

			for _, province := range provinces {
				for _, college := range colleges {
					url := fmt.Sprintf("https://zsb.suda.edu.cn/search.aspx?nf=%s&sf=%s&xy=%s", year, province, college)
					if err := detailCollector.Visit(url); err != nil {
						logrus.Errorf("%s crawl %s err: %v", u.name, url, err)
					}
				}
			}
		}
	})

	detailCollector.OnHTML(`table[id=ctl00_ContentPlaceHolder1_GridView1]`, func(element *colly.HTMLElement) {
		element.ForEach(`tr`, func(i int, element *colly.HTMLElement) {
			admissionTime := element.ChildText("td:nth-of-type(1)")
			if admissionTime == "年份" || admissionTime == "" {
				return
			}

			province := element.ChildText("td:nth-of-type(2)")
			major := strings.SplitN(element.ChildText("td:nth-of-type(3)"), "--", 2)
			duration := element.ChildText("td:nth-of-type(4)")
			selectExam := element.ChildText("td:nth-of-type(5)")
			maxScore := element.ChildText("td:nth-of-type(6)")
			minScore := element.ChildText("td:nth-of-type(7)")
			averageScore := element.ChildText("td:nth-of-type(8)")

			if len(major) == 2 {
				selectExam = fmt.Sprintf("%s(%s)", selectExam, strings.Split(major[1], "，")[0])
			}

			if err := storage.GetQueries().CreateAdmissionMajor(context.Background(), storage.CreateAdmissionMajorParams{
				University:    u.name,
				Major:         major[0],
				SelectExam:    selectExam,
				Province:      province,
				AdmissionTime: admissionTime,
				Duration:      duration,
				MaxScore:      maxScore,
				MinScore:      minScore,
				AverageScore:  averageScore,
			}); err != nil {
				logrus.Errorf("create admission major err: %v", err)
			}
		})
	})

	return c.Visit("https://zsb.suda.edu.cn/markHistory.aspx")
}

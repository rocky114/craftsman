package scraper

import (
	"context"
	"fmt"
	"strings"

	"github.com/spf13/cast"

	"github.com/rocky114/craftsman/internal/storage"
	"github.com/sirupsen/logrus"

	"github.com/gocolly/colly/v2"
)

func ScrapeAdmissionMajorScoreSuda() error {
	c := colly.NewCollector(colly.CacheDir("./web"))

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

		/*if err := detailCollector.Visit("https://zsb.suda.edu.cn/search.aspx?nf=2022&sf=10&xy=1"); err != nil {
			logrus.Errorf("scrape suzhou university err: %v", err)
		}*/

		for _, year := range years {
			for _, province := range provinces {
				for _, college := range colleges {
					url := fmt.Sprintf("https://zsb.suda.edu.cn/search.aspx?nf=%s&sf=%s&xy=%s", year, province, college)
					if err := detailCollector.Visit(url); err != nil {
						logrus.Errorf("scrape suzhou university err: %v", err)
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
			duration := cast.ToInt32(element.ChildText("td:nth-of-type(4)"))
			subjectType := element.ChildText("td:nth-of-type(5)")
			maxScore := cast.ToInt32(cast.ToFloat32(element.ChildText("td:nth-of-type(6)")))
			minScore := cast.ToInt32(cast.ToFloat32(element.ChildText("td:nth-of-type(7)")))
			averageScore := cast.ToInt32(cast.ToFloat32(element.ChildText("td:nth-of-type(8)")))

			selectExam := ""
			if len(major) == 2 {
				selectExam = major[1]
			}

			if err := storage.GetQueries().CreateAdmissionMajor(context.Background(), storage.CreateAdmissionMajorParams{
				College:       "苏州大学",
				Major:         major[0],
				SelectExam:    selectExam,
				Province:      province,
				SubjectType:   subjectType,
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

	c.OnRequest(func(request *colly.Request) {
		fmt.Println("visiting", request.URL.String())
	})

	return c.Visit("https://zsb.suda.edu.cn/markHistory.aspx")
}

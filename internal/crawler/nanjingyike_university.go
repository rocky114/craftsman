package crawler

import (
	"context"
	"fmt"
	"strings"

	"github.com/rocky114/craftsman/internal/storage"

	"github.com/sirupsen/logrus"

	"github.com/rocky114/craftsman/internal/pkg/path"

	"github.com/gocolly/colly/v2"
)

type nanjingyikeUniversity struct {
	university
}

func init() {
	collection["4132010312"] = &nanjingyikeUniversity{university{
		name: "南京医科大学",
		code: "4132010312",
	}}
}

func (u *nanjingyikeUniversity) crawl(ctx context.Context) error {
	c := colly.NewCollector(colly.UserAgent(userAgent), colly.CacheDir(path.GetTmpPath("nanjingyikeUniversity")), colly.AllowURLRevisit())

	detailCollector := c.Clone()

	detailCollector.OnHTML("div.wp_articlecontent div.Article_Content", func(element *colly.HTMLElement) {
		element.ForEach("div", func(i int, element *colly.HTMLElement) {
			admissionType, major, selectExam, maxScore, minScore := "", "", "", "", ""
			element.ForEach("table tbody tr", func(i int, element *colly.HTMLElement) {
				if i == 0 {
					return
				}

				isMergeCell := element.ChildAttr("td:nth-of-type(1)", "rowspan") == ""
				if isMergeCell {
					major = element.ChildText("td:nth-of-type(1)")
					maxScore = element.ChildText("td:nth-of-type(2)")
					minScore = element.ChildText("td:nth-of-type(3)")
				} else {
					admissionType = element.ChildText("td:nth-of-type(1)")
					major = element.ChildText("td:nth-of-type(2)")
					selectExam = element.ChildText("td:nth-of-type(3)")
					maxScore = element.ChildText("td:nth-of-type(4)")
					minScore = element.ChildText("td:nth-of-type(5)")
				}

				if err := storage.GetQueries().CreateAdmissionMajor(context.Background(), storage.CreateAdmissionMajorParams{
					University:    u.name,
					AdmissionTime: u.admissionTime,
					Province:      "江苏",
					Major:         major,
					SelectExam:    selectExam,
					MinScore:      minScore,
					MaxScore:      maxScore,
					AdmissionType: admissionType,
				}); err != nil {
					logrus.Errorf("create admission major err: %v", err)
				}
			})
		})
	})

	c.OnHTML(`div[id=newslist] div[id=wp_news_w3] table tbody`, func(element *colly.HTMLElement) {
		element.ForEach("tr", func(i int, element *colly.HTMLElement) {
			title := element.ChildText("td:nth-of-type(2)")
			if strings.Contains(title, fmt.Sprintf("南京医科大学%s年江苏省各专业", u.admissionTime)) {
				addr := element.ChildAttr("td:nth-of-type(2) table tr td a", "href")
				if err := detailCollector.Visit(fmt.Sprintf("https://zs.njmu.edu.cn%s", addr)); err != nil {
					logrus.Errorf("nanjingyikeUniversity err: %v", err)
				}
			}
		})
	})

	return c.Visit("https://zs.njmu.edu.cn/3431/list.htm")
}

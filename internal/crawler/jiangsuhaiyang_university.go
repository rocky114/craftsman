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

type jiangsuhaiyangUniversity struct {
	university
}

func init() {
	collection["4132011641"] = &jiangsuhaiyangUniversity{university{
		name: "江苏海洋大学",
		code: "4132011641",
	}}
}

func (u *jiangsuhaiyangUniversity) crawl(ctx context.Context) error {
	c := colly.NewCollector(colly.UserAgent(userAgent), colly.CacheDir(path.GetTmpPath("jiangsuhaiyangUniversity")), colly.AllowURLRevisit())

	detailCollector := c.Clone()

	detailCollector.OnHTML(`div.v_news_content div.WordSection1 div:nth-of-type(1)`, func(element *colly.HTMLElement) {
		province, admissionType, selectExam, admissionNumber, maxScore, minScore := "", "", "", "", "", ""
		element.ForEach("table tbody tr", func(i int, element *colly.HTMLElement) {
			length := len(element.DOM.Find("td").Nodes)
			isMergeCell := element.ChildAttr("td:nth-of-type(1)", "rowspan") != ""
			if isMergeCell {
				province = element.ChildText("td:nth-of-type(1) p:nth-of-type(1)")
				admissionType = element.ChildText("td:nth-of-type(1) p:nth-of-type(2)")
				selectExam = element.ChildText("td:nth-of-type(2)")
			} else {
				selectExam = element.ChildText("td:nth-of-type(1)")
			}

			admissionNumber = element.ChildText(fmt.Sprintf("td:nth-of-type(%d)", length-2))
			maxScore = element.ChildText(fmt.Sprintf("td:nth-of-type(%d)", length-1))
			minScore = element.ChildText(fmt.Sprintf("td:nth-of-type(%d)", length))

			if err := storage.GetQueries().CreateAdmissionMajor(context.Background(), storage.CreateAdmissionMajorParams{
				University:      u.name,
				AdmissionTime:   u.admissionTime,
				Province:        province,
				AdmissionType:   admissionType,
				SelectExam:      selectExam,
				AdmissionNumber: admissionNumber,
				MinScore:        minScore,
				MaxScore:        maxScore,
			}); err != nil {
				logrus.Errorf("create admission major err: %v", err)
			}
		})
	})

	c.OnHTML(`div.main div.listr div.newslist ul`, func(element *colly.HTMLElement) {
		element.ForEach("li", func(i int, element *colly.HTMLElement) {
			title := element.ChildText("a")
			if strings.Contains(title, fmt.Sprintf("江苏海洋大学%s年各省录取分数", u.admissionTime)) {
				if err := detailCollector.Visit(fmt.Sprintf("https://zsxx.jou.edu.cn/%s", element.ChildAttr("a", "href"))); err != nil {
					logrus.Errorf("jiangsuhaiyangUniversity err: %v", err)
				}
			}
		})
	})

	return c.Visit("https://zsxx.jou.edu.cn/wnlq.htm")
}

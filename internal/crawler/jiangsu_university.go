package crawler

import (
	"context"
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/rocky114/craftsman/internal/pkg/path"

	"github.com/gocolly/colly/v2"
)

type jiangsuUniversity struct {
	university
}

func init() {
	collection["4132010299"] = &jiangsuUniversity{university{
		name: "江苏大学",
		code: "4132010299",
	}}
}

func (u *jiangsuUniversity) crawl(ctx context.Context) error {
	c := colly.NewCollector(colly.CacheDir(path.GetTmpPath()))

	provinceCollector := c.Clone()
	detailCollector := c.Clone()

	c.OnHTML(`div.content_neirong div.left div.pro_list ul`, func(element *colly.HTMLElement) {
		element.ForEach("li", func(i int, element *colly.HTMLElement) {
			addr := element.ChildAttr("a", "href")
			if err := provinceCollector.Visit(fmt.Sprintf("https://zb.ujs.edu.cn/%s", addr)); err != nil {
				logrus.Errorf("jiangsuUniversity err: %v", err)
			}
		})
	})

	provinceCollector.OnHTML(`div.content_neirong div.right div.right_news_content ul`, func(element *colly.HTMLElement) {
		element.ForEach(`li`, func(i int, element *colly.HTMLElement) {
			admissionTime := strings.TrimPrefix(element.ChildText("p.pl"), "江苏大学")[0:4]
			if admissionTime != u.admissionTime {
				return
			}
			addr := element.ChildAttr("p.pl a", "href")

			if err := detailCollector.Visit(fmt.Sprintf("https://zb.ujs.edu.cn/%s", addr)); err != nil {
				logrus.Errorf("jiangsuUniversity err: %v", err)
			}
		})
	})

	detailCollector.OnHTML(`div.content_neirong div.right`, func(element *colly.HTMLElement) {
		title := strings.TrimPrefix(element.ChildText("div.right_title2 p.p2"), fmt.Sprintf("江苏大学%s", u.admissionTime))
		province := string([]rune(title)[2:4])

		element.ForEach("div.right_content2 table tr", func(i int, element *colly.HTMLElement) {
			if i == 0 {
				return
			}

			admissionType, major, selectExam, admissionNumber, maxScore, minScore, averageScore := "", "", "", "", "", "", ""

			if province == "江苏" {
				admissionType = element.ChildText("td:nth-of-type(1)")
				selectExam = element.ChildText("td:nth-of-type(2)")[2:]
				major = element.ChildText("td:nth-of-type(3)")
				admissionNumber = element.ChildText("td:nth-of-type(4)")
				maxScore = element.ChildText("td:nth-of-type(5)")
				minScore = element.ChildText("td:nth-of-type(6)")
				averageScore = element.ChildText("td:nth-of-type(7)")
			} else if province == "广东" || province == "湖南" {
				selectExam = element.ChildText("td:nth-of-type(1)")
				major = element.ChildText("td:nth-of-type(2)")
				admissionNumber = element.ChildText("td:nth-of-type(3)")
				maxScore = element.ChildText("td:nth-of-type(4)")
				minScore = element.ChildText("td:nth-of-type(5)")
				averageScore = element.ChildText("td:nth-of-type(6)")
			} else {
				admissionType = element.ChildText("td:nth-of-type(1)")
				major = element.ChildText("td:nth-of-type(2)")
				admissionNumber = element.ChildText("td:nth-of-type(3)")
				maxScore = element.ChildText("td:nth-of-type(4)")
				minScore = element.ChildText("td:nth-of-type(5)")
				averageScore = element.ChildText("td:nth-of-type(6)")
			}

			fmt.Println(province, admissionType, major, selectExam, admissionNumber, maxScore, minScore, averageScore)
		})
	})

	return c.Visit("https://zb.ujs.edu.cn/lnfs.htm")
}

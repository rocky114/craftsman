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
	c := colly.NewCollector(colly.UserAgent(userAgent), colly.CacheDir(path.GetTmpPath()))

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

	detailCollector.OnHTML(`div.content_neirong`, func(element *colly.HTMLElement) {
		province := element.ChildText("div.site div a:nth-of-type(3)")

		head := map[int]string{}
		element.ForEach("div.right_content2 table tr.firstRow td", func(i int, element *colly.HTMLElement) {
			head[i] = element.Text
		})

		mergeSelectExam, mergeAdmissionType := "", ""
		isMergeCell := false

		element.ForEach("div.right_content2 table tr", func(i int, element *colly.HTMLElement) {
			if element.ChildAttr("td:nth-of-type(1)", "rowspan") != "" {
				isMergeCell = true
			}
		})
		element.ForEach("div.right_content2 table tr", func(i int, element *colly.HTMLElement) {
			if i == 0 {
				return
			}

			childMergeCell := element.ChildAttr("td:nth-of-type(1)", "rowspan") == ""
			admissionType, major, selectExam, admissionNumber, maxScore, minScore, averageScore := "", "", "", "", "", "", ""

			element.ForEach("td", func(i int, element *colly.HTMLElement) {
				if isMergeCell && childMergeCell {
					i = i + 1
				}

				switch head[i] {
				case "批次":
					admissionType = element.Text
				case "专业组":
					selectExam = element.Text[2:]
				case "专业":
					major = element.Text
				case "录取人数":
					admissionNumber = element.Text
				case "最高分":
					maxScore = element.Text
				case "最低分":
					minScore = element.Text
				case "平均分":
					averageScore = element.Text
				case "选考要求":
					selectExam = element.Text
				case "首选科目":
					selectExam = element.Text
				case "选考科目":
					selectExam = element.Text
				}

				if isMergeCell && !childMergeCell {
					mergeSelectExam = selectExam
					mergeAdmissionType = admissionType
				}
			})

			if isMergeCell {
				selectExam = mergeSelectExam
				admissionType = mergeAdmissionType
			}

			if err := storage.GetQueries().CreateAdmissionMajor(context.Background(), storage.CreateAdmissionMajorParams{
				University:      u.name,
				Major:           major,
				Province:        province,
				AdmissionTime:   u.admissionTime,
				MaxScore:        maxScore,
				MinScore:        minScore,
				AverageScore:    averageScore,
				AdmissionNumber: admissionNumber,
				SelectExam:      selectExam,
				AdmissionType:   admissionType,
			}); err != nil {
				logrus.Errorf("create admission major err: %v", err)
			}
		})
	})

	return c.Visit("https://zb.ujs.edu.cn/lnfs.htm")
}

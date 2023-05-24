package crawler

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/rocky114/craftsman/internal/pkg/path"

	"github.com/rocky114/craftsman/internal/storage"
	"github.com/sirupsen/logrus"

	"github.com/gocolly/colly/v2"
)

type dongnanUniversity struct {
	university
}

func init() {
	collection["4132010286"] = &dongnanUniversity{university{
		name: "东南大学",
		code: "4132010286",
	}}
}

func (u *dongnanUniversity) crawl(ctx context.Context) error {
	c := colly.NewCollector(colly.CacheDir(path.GetTmpPath()))
	c.SetRequestTimeout(60 * time.Second)

	listCollector := c.Clone()
	detailCollector := c.Clone()

	c.OnHTML(`map[id=Map]`, func(element *colly.HTMLElement) {
		element.ForEach(`area`, func(i int, element *colly.HTMLElement) {
			url := element.Attr("href")
			if err := listCollector.Visit(url); err != nil {
				logrus.Errorf("scrape %s university err: %v", u.name, err)
			}
		})
	})

	listCollector.OnHTML(`div[id=wp_news_w6]`, func(element *colly.HTMLElement) {
		element.ForEach(`ul li`, func(i int, element *colly.HTMLElement) {
			title := element.ChildText("span:nth-of-type(1) a")
			if !containAdmissionTime(title[0:4]) {
				return
			}

			if !strings.Contains(title, "专业分数线") {
				return
			}

			url := element.ChildAttr("span:nth-of-type(1) a", "href")
			if err := detailCollector.Visit(fmt.Sprintf("https://zsb.seu.edu.cn/%s", url)); err != nil {
				logrus.Errorf("scrape %s university err: %v", u.name, err)
			}
		})
	})

	detailCollector.OnHTML(`div[id=wp_news_w6]`, func(element *colly.HTMLElement) {
		element.ForEach(`ul li`, func(i int, element *colly.HTMLElement) {
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

	return c.Visit("http://zsb.seu.edu.cn/")
}

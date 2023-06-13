package crawler

import (
	"context"
	"time"

	"github.com/rocky114/craftsman/internal/types"

	"github.com/rocky114/craftsman/internal/storage"
	"github.com/sirupsen/logrus"

	"github.com/rocky114/craftsman/internal/pkg/path"

	"github.com/gocolly/colly/v2"
)

type nanjingnongyeUniversity struct {
	university
}

func init() {
	collection["4132010307"] = &nanjingnongyeUniversity{university{
		name: "南京农业大学",
		code: "4132010307",
	}}
}

func (u *nanjingnongyeUniversity) crawl(ctx context.Context) error {
	c := colly.NewCollector(colly.UserAgent(userAgent), colly.CacheDir(path.GetTmpPath("nanjingnongyeUniversity")), colly.AllowURLRevisit())

	c.OnHTML(`div.cxjg table tbody`, func(element *colly.HTMLElement) {
		element.ForEach("tr", func(i int, element *colly.HTMLElement) {
			if i == 0 {
				return
			}

			if err := storage.GetQueries().CreateAdmissionMajor(context.Background(), storage.CreateAdmissionMajorParams{
				University:    u.name,
				AdmissionTime: u.admissionTime,
				Province:      types.Provinces[currentIndex],
				Major:         element.ChildText("td:nth-of-type(1)"),
				SelectExam:    element.ChildText("td:nth-of-type(2)"),
				MinScore:      element.ChildText("td:nth-of-type(4)"),
				MaxScore:      element.ChildText("td:nth-of-type(3)"),
				AverageScore:  element.ChildText("td:nth-of-type(5)"),
			}); err != nil {
				logrus.Errorf("create admission major err: %v", err)
			}
		})
	})

	for i, province := range types.Provinces {
		currentIndex = i
		if err := c.Post("https://zsxx.njau.edu.cn/lnlqfs.jsp?wbtreeid=1024", map[string]string{"nf": u.admissionTime, "sf": province}); err != nil {
			logrus.Errorf("nanjingnongyeUniversity err: %v", err)
		}

		time.Sleep(3 * time.Second)
	}

	return c.Visit("https://zsxx.jou.edu.cn/wnlq.htm")
}

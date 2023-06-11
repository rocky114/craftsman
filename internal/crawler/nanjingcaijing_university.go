package crawler

import (
	"context"

	"github.com/rocky114/craftsman/internal/storage"
	"github.com/sirupsen/logrus"

	"github.com/rocky114/craftsman/internal/pkg/path"

	"github.com/spf13/cast"

	"github.com/gocolly/colly/v2"
)

type nanjingcaijingUniversity struct {
	university
}

func init() {
	collection["4132010327"] = &nanjingcaijingUniversity{university{
		name: "南京财经大学",
		code: "4132010327",
	}}
}

func (u *nanjingcaijingUniversity) crawl(ctx context.Context) error {
	c := colly.NewCollector(colly.UserAgent(userAgent), colly.CacheDir(path.GetTmpPath("nanjingcaijingUniversity")), colly.AllowURLRevisit())

	c.OnHTML(`div.table_box table tbody`, func(element *colly.HTMLElement) {
		element.ForEach("tr", func(i int, element *colly.HTMLElement) {
			if err := storage.GetQueries().CreateAdmissionMajor(context.Background(), storage.CreateAdmissionMajorParams{
				University:               u.name,
				AdmissionTime:            u.admissionTime,
				Province:                 element.ChildText("td:nth-of-type(2)"),
				AdmissionType:            element.ChildText("td:nth-of-type(3)"),
				SelectExam:               element.ChildText("td:nth-of-type(4)"),
				Major:                    element.ChildText("td:nth-of-type(5)"),
				AdmissionNumber:          element.ChildText("td:nth-of-type(6)"),
				ProvinceControlScoreLine: element.ChildText("td:nth-of-type(7)"),
				MinScore:                 element.ChildText("td:nth-of-type(8)"),
				AverageScore:             element.ChildText("td:nth-of-type(9)"),
				MaxScore:                 element.ChildText("td:nth-of-type(10)"),
			}); err != nil {
				logrus.Errorf("create admission major err: %v", err)
			}
		})
	})

	return c.Post("http://bkzs.nufe.edu.cn/ldcx.jsp?wbtreeid=1023", map[string]string{"nf": cast.ToString(u.admissionTime)})
}

package crawler

import (
	"context"

	"github.com/rocky114/craftsman/internal/storage"

	"github.com/sirupsen/logrus"

	"github.com/rocky114/craftsman/internal/pkg/path"

	"github.com/gocolly/colly/v2"
)

type nanjingxinxigongchengUniversity struct {
	university
}

func init() {
	collection["4132010300"] = &nanjingxinxigongchengUniversity{university{
		name: "南京信息工程大学",
		code: "4132010300",
	}}
}

func (u *nanjingxinxigongchengUniversity) crawl(ctx context.Context) error {
	c := colly.NewCollector(colly.UserAgent(userAgent), colly.CacheDir(path.GetTmpPath()))

	c.OnHTML(`div.con_wrap table.seaech_list`, func(element *colly.HTMLElement) {
		element.ForEach("tr", func(i int, element *colly.HTMLElement) {
			if i == 0 {
				return
			}

			admissionTime := element.ChildText("td:nth-of-type(2)")
			province := element.ChildText("td:nth-of-type(3)")
			college := element.ChildText("td:nth-of-type(4)")
			major := element.ChildText("td:nth-of-type(5)")
			selectExam := element.ChildText("td:nth-of-type(6)")
			admissionType := element.ChildText("td:nth-of-type(7)")
			admissionNumber := element.ChildText("td:nth-of-type(8)")
			provinceControlScoreLine := element.ChildText("td:nth-of-type(9)")
			minScore := element.ChildText("td:nth-of-type(10)")
			maxScore := element.ChildText("td:nth-of-type(11)")
			averageScore := element.ChildText("td:nth-of-type(12)")

			if err := storage.GetQueries().CreateAdmissionMajor(ctx, storage.CreateAdmissionMajorParams{
				University:               u.name,
				College:                  college,
				Province:                 province,
				Major:                    major,
				AdmissionType:            admissionType,
				SelectExam:               selectExam,
				AdmissionTime:            admissionTime,
				AdmissionNumber:          admissionNumber,
				ProvinceControlScoreLine: provinceControlScoreLine,
				MinScore:                 minScore,
				MaxScore:                 maxScore,
				AverageScore:             averageScore,
			}); err != nil {
				logrus.Errorf("create admission major err: %v", err)
			}
		})
	})

	return c.Post("https://zs.nuist.edu.cn/wnfs.jsp?wbtreeid=1011", map[string]string{"nf": u.admissionTime})
}

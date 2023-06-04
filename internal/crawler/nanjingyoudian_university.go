package crawler

import (
	"context"
	"fmt"
	"strings"

	"github.com/rocky114/craftsman/internal/storage"

	"github.com/xuri/excelize/v2"

	"github.com/sirupsen/logrus"

	"github.com/rocky114/craftsman/internal/pkg/path"

	"github.com/gocolly/colly/v2"
)

type nanjingyoudianUniversity struct {
	university
}

func init() {
	collection["4132010293"] = &nanjingyoudianUniversity{university{
		name: "南京邮电大学",
		code: "4132010293",
	}}
}

func (u *nanjingyoudianUniversity) crawl(ctx context.Context) error {
	c := colly.NewCollector(colly.CacheDir(path.GetTmpPath()))

	detailCollector := c.Clone()
	excelCollector := c.Clone()

	c.OnHTML("div[id=wp_news_w51] ul.news_list", func(element *colly.HTMLElement) {
		element.ForEach("li", func(i int, element *colly.HTMLElement) {
			admissionTime := element.ChildText("span.news_meta")
			if strings.TrimSpace(admissionTime) == u.admissionTime {
				uri := strings.TrimSpace(element.ChildAttr("span.news_title a", "href"))
				if err := detailCollector.Visit(fmt.Sprintf("https://zs.njupt.edu.cn/%s", uri)); err != nil {
					logrus.Errorf("nanjingyoudianUniversity err: %v", err)
				}
			}
		})
	})

	detailCollector.OnHTML("div.wrapper div.wp_articlecontent", func(element *colly.HTMLElement) {
		uri := element.ChildAttr("p a", "href")
		if err := excelCollector.Visit(fmt.Sprintf("https://zs.njupt.edu.cn/%s", uri)); err != nil {
			logrus.Errorf("nanjingyoudianUniversity err: %v", err)
		}
	})

	excelCollector.OnResponse(func(response *colly.Response) {
		filename := fmt.Sprintf("%s%s", path.GetTmpPath(), response.FileName())
		if err := response.Save(filename); err != nil {
			logrus.Errorf("nanjingyoudianUniversity err: %v", err)
			return
		}

		f, err := excelize.OpenFile(filename)
		if err != nil {
			logrus.Errorf("nanjingyoudianUniversity err: %v", err)
			return
		}

		defer func() {
			if err = f.Close(); err != nil {
				logrus.Errorf("nanjingyoudianUniversity err: %v", err)
			}
		}()

		for _, sheet := range f.GetSheetList() {
			rows, err := f.GetRows(sheet)
			if err != nil {
				logrus.Errorf("nanjingyoudianUniversity err: %v", err)
				return
			}

			admissionType, selectExam := "", ""
			for i, row := range rows {
				if i < 2 {
					continue
				}

				if row[0] != "" {
					admissionType = row[0]
				}
				if row[1] != "" {
					selectExam = row[1]
				}

				if err = storage.GetQueries().CreateAdmissionMajor(ctx, storage.CreateAdmissionMajorParams{
					University:      u.name,
					Province:        sheet,
					Major:           row[2],
					AdmissionType:   admissionType,
					AdmissionNumber: row[3],
					SelectExam:      selectExam,
					MaxScore:        row[4],
					MinScore:        row[5],
					AverageScore:    row[6],
					AdmissionTime:   u.admissionTime,
				}); err != nil {
					logrus.Errorf("nanjingyoudianUniversity err: %v", err)
				}
			}
		}
	})

	return c.Visit("https://zs.njupt.edu.cn")
}

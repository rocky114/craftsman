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

type changzhouUniversity struct {
	university
}

func init() {
	collection["4132010292"] = &changzhouUniversity{university{
		name: "常州大学",
		code: "4132010292",
	}}
}

func (u *changzhouUniversity) crawl(ctx context.Context) error {
	c := colly.NewCollector(colly.CacheDir(path.GetLogPath()))

	listCollector := c.Clone()
	detailCollector := c.Clone()
	excelCollector := c.Clone()

	var excelFiles []string
	c.OnHTML("div ul.wp_listcolumn", func(element *colly.HTMLElement) {
		element.ForEach("li", func(i int, element *colly.HTMLElement) {
			if err := listCollector.Visit("https://cdzs.cczu.edu.cn" + element.ChildAttr("a", "href")); err != nil {
				logrus.Errorf("changzhouUniversity err: %v", err)
			}
		})
	})

	listCollector.OnHTML("div ul.wp_article_list", func(element *colly.HTMLElement) {
		element.ForEach("li", func(i int, element *colly.HTMLElement) {
			title := element.ChildText("div.pr_fields span.Article_Title")
			if strings.TrimSpace(title) == u.admissionTime {
				uri := element.ChildAttr("div.pr_fields span.Article_Title a", "href")
				if err := detailCollector.Visit("https://cdzs.cczu.edu.cn" + uri); err != nil {
					logrus.Errorf("changzhouUniversity err: %v", err)
				}
			}
		})
	})

	detailCollector.OnHTML("div.wp_articlecontent", func(element *colly.HTMLElement) {
		if err := excelCollector.Visit("https://cdzs.cczu.edu.cn" + element.ChildAttr("a", "href")); err != nil {
			logrus.Errorf("changzhouUniversity err: %v", err)
		}
	})

	excelCollector.OnResponse(func(response *colly.Response) {
		filename := fmt.Sprintf("%s%s", path.GetTmpPath(), response.FileName())
		if err := response.Save(filename); err != nil {
			logrus.Errorf("changzhouUniversity err: %v", err)
		} else {
			excelFiles = append(excelFiles, filename)
		}
	})

	c.OnScraped(func(response *colly.Response) {
		for _, excelFile := range excelFiles {
			u.CreateAdmissionMajor(ctx, excelFile)
		}
	})

	return c.Visit("https://cdzs.cczu.edu.cn//lnfswsjw/list.htm")
}

func (u *changzhouUniversity) CreateAdmissionMajor(ctx context.Context, file string) {
	f, err := excelize.OpenFile(file)
	if err != nil {
		logrus.Errorf("CreateAdmissionMajor err: %v", err)
		return
	}

	defer func() {
		if err = f.Close(); err != nil {
			logrus.Errorf("CreateAdmissionMajor err: %v", err)
		}
	}()

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		logrus.Errorf("CreateAdmissionMajor err: %v", err)
		return
	}

	for i, row := range rows {
		if i == 0 {
			continue
		}

		if err = storage.GetQueries().CreateAdmissionMajor(ctx, storage.CreateAdmissionMajorParams{
			University:      u.name,
			Province:        row[0],
			Major:           row[3],
			AdmissionType:   row[2],
			AdmissionNumber: row[4],
			SelectExam:      row[1],
			MaxScore:        row[5],
			MinScore:        row[6],
			AdmissionTime:   u.admissionTime,
		}); err != nil {
			logrus.Errorf("create admission major err: %v", err)
		}
	}
}

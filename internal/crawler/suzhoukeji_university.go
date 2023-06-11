package crawler

import (
	"context"
	"fmt"

	"github.com/rocky114/craftsman/internal/storage"

	"github.com/rocky114/craftsman/internal/pkg/path"

	"github.com/sirupsen/logrus"

	"github.com/gocolly/colly/v2"
)

type suzhoukejiUniversity struct {
	university
}

var requests []struct {
	province string
	uri      string
}

func init() {
	collection["4132010332"] = &suzhoukejiUniversity{university{
		name: "苏州科技大学",
		code: "4132010332",
	}}
}

func init() {
	requests = []struct {
		province string
		uri      string
	}{
		{
			province: "北京",
			uri:      "&sStyle=%C6%D5%B8%DF&sZone=%B1%B1%BE%A9",
		},
		{
			province: "天津",
			uri:      "&sStyle=%C6%D5%B8%DF&sZone=%CC%EC%BD%F2",
		},
		{
			province: "河北",
			uri:      "&sStyle=%C6%D5%B8%DF&sZone=%BA%D3%B1%B1",
		}, {
			province: "山西",
			uri:      "&sStyle=%C6%D5%B8%DF&sZone=%C9%BD%CE%F7",
		},
		{
			province: "辽宁",
			uri:      "&sStyle=%C6%D5%B8%DF&sZone=%C1%C9%C4%FE",
		},
		{
			province: "吉林",
			uri:      "&sStyle=%C6%D5%B8%DF&sZone=%BC%AA%C1%D6",
		},
		{
			province: "黑龙江",
			uri:      "&sStyle=%C6%D5%B8%DF&sZone=%BA%DA%C1%FA%BD%AD",
		},
		{
			province: "上海",
			uri:      "&sStyle=%C6%D5%B8%DF&sZone=%C9%CF%BA%A3",
		},
		{
			province: "江苏",
			uri:      "&sStyle=%C6%D5%B8%DF&sZone=%BD%AD%CB%D5",
		},
		{
			province: "浙江",
			uri:      "&sStyle=%C6%D5%B8%DF&sZone=%D5%E3%BD%AD",
		},
		{
			province: "安徽",
			uri:      "&sStyle=%C6%D5%B8%DF&sZone=%B0%B2%BB%D5",
		},
		{
			province: "福建",
			uri:      "&sStyle=%C6%D5%B8%DF&sZone=%B8%A3%BD%A8",
		},
		{
			province: "江西",
			uri:      "&sStyle=%C6%D5%B8%DF&sZone=%BD%AD%CE%F7",
		},
		{
			province: "山东",
			uri:      "&sStyle=%C6%D5%B8%DF&sZone=%C9%BD%B6%AB",
		},
		{
			province: "河南",
			uri:      "&sStyle=%C6%D5%B8%DF&sZone=%BA%D3%C4%CF",
		},
		{
			province: "湖北",
			uri:      "&sStyle=%C6%D5%B8%DF&sZone=%BA%FE%B1%B1",
		},
		{
			province: "湖南",
			uri:      "&sStyle=%C6%D5%B8%DF&sZone=%BA%FE%C4%CF",
		},
		{
			province: "广东",
			uri:      "&sStyle=%C6%D5%B8%DF&sZone=%B9%E3%B6%AB",
		},
		{
			province: "广西",
			uri:      "&sStyle=%C6%D5%B8%DF&sZone=%B9%E3%CE%F7",
		},
		{
			province: "海南",
			uri:      "&sStyle=%C6%D5%B8%DF&sZone=%BA%A3%C4%CF",
		},
		{
			province: "重庆",
			uri:      "&sStyle=%C6%D5%B8%DF&sZone=%D6%D8%C7%EC",
		},
		{
			province: "四川",
			uri:      "&sStyle=%C6%D5%B8%DF&sZone=%CB%C4%B4%A8",
		},
		{
			province: "贵州",
			uri:      "&sStyle=%C6%D5%B8%DF&sZone=%B9%F3%D6%DD",
		},
		{
			province: "云南",
			uri:      "&sStyle=%C6%D5%B8%DF&sZone=%D4%C6%C4%CF",
		},
		{
			province: "陕西",
			uri:      "&sStyle=%C6%D5%B8%DF&sZone=%C9%C2%CE%F7",
		},
		{
			province: "甘肃",
			uri:      "&sStyle=%C6%D5%B8%DF&sZone=%B8%CA%CB%E0",
		},
		{
			province: "青海",
			uri:      "&sStyle=%C6%D5%B8%DF&sZone=%C7%E0%BA%A3",
		},
		{
			province: "宁夏",
			uri:      "&sStyle=%C6%D5%B8%DF&sZone=%C4%FE%CF%C4",
		},
		{
			province: "新疆",
			uri:      "&sStyle=%C6%D5%B8%DF&sZone=%D0%C2%BD%AE",
		},
	}
}

func (u *suzhoukejiUniversity) crawl(ctx context.Context) error {
	c := colly.NewCollector(colly.UserAgent(userAgent), colly.CacheDir(path.GetTmpPath("suzhoukejiUniversity")))

	c.OnHTML(`table > tbody`, func(element *colly.HTMLElement) {
		element.ForEach("tr", func(i int, element *colly.HTMLElement) {
			title := element.ChildText("td:nth-of-type(1)")
			if title == "合计" || title == "序号" || title == "" {
				return
			}

			if err := storage.GetQueries().CreateAdmissionMajor(context.Background(), storage.CreateAdmissionMajorParams{
				University:               u.name,
				AdmissionType:            element.ChildText("td:nth-of-type(2)"),
				SelectExam:               element.ChildText("td:nth-of-type(3)"),
				Major:                    element.ChildText("td:nth-of-type(4)"),
				AdmissionNumber:          element.ChildText("td:nth-of-type(5)"),
				Province:                 requests[currentIndex].province,
				AdmissionTime:            u.admissionTime,
				MaxScore:                 element.ChildText("td:nth-of-type(6)"),
				MinScore:                 element.ChildText("td:nth-of-type(7)"),
				ProvinceControlScoreLine: element.ChildText("td:nth-of-type(8)"),
			}); err != nil {
				logrus.Errorf("create admission major err: %v", err)
			}
		})
	})

	for i, item := range requests {
		currentIndex = i

		addr := fmt.Sprintf("https://zsb.usts.edu.cn/news/relqfs.asp?sYear=%s", u.admissionTime) + item.uri
		if err := c.Visit(addr); err != nil {
			logrus.Errorf("suzhoukejiUniversity params: %s, err: %v", addr, err)
		}
	}

	return nil
}

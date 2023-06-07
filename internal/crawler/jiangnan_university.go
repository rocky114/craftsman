package crawler

import (
	"context"
	"fmt"

	"github.com/rocky114/craftsman/internal/pkg/path"

	"github.com/gocolly/colly/v2"
)

type jiangnanUniversity struct {
	university
}

func init() {
	collection["4132010295"] = &jiangnanUniversity{university{
		name: "江南大学",
		code: "4132010295",
	}}
}

type jiangnanUniversityResp struct {
	Data []struct {
		Province     string `json:"province"`
		Type         string `json:"type"`
		Discipline   string `json:"discipline"`
		Major        string `json:"major"`
		FileScore    int    `json:"filescore"`
		Highestscore int    `json:"highestscore"`
		Red          string `json:"red"`
	} `json:"data"`
}

func (u *jiangnanUniversity) crawl(ctx context.Context) error {
	c := colly.NewCollector(colly.CacheDir(path.GetTmpPath()))

	yearCollector := c.Clone()
	provinceCollector := c.Clone()
	selectExamCollector := c.Clone()
	admissionTypeCollector := c.Clone()
	detailCollector := c.Clone()

	c.OnResponse(func(response *colly.Response) {
		yearCollector.PostRaw("http://admission3.jiangnan.edu.cn:3001/api/front/historyScore/getQuery", []byte{})
		provinceCollector.PostRaw("http://admission3.jiangnan.edu.cn:3001/api/front/historyScore/getNotArtProvinceQuery?year=2022", []byte{})

		params := `{"provinceId":10,"year":2022}`
		selectExamCollector.PostRaw("http://admission3.jiangnan.edu.cn:3001/api/front/subject/getQueryByYearAndProvinceIdNew", []byte(params))

		params = `{"provinceId":10,"year":2022,"subjectId":3}`
		admissionTypeCollector.PostRaw("http://admission3.jiangnan.edu.cn:3001/api/front/recruitmentPlan/getQueryByYearAndProvinceIdWithOutArt", []byte(params))

		params = `{"provinceId":10,"recruitmentCategoryId":1,"subjectId":3,"year":2022}`
		detailCollector.PostRaw("http://admission3.jiangnan.edu.cn:3001/api/front/historyScore/getList", []byte(params))
	})

	yearCollector.OnResponse(func(response *colly.Response) {
		fmt.Println(string(response.Body))
	})

	provinceCollector.OnResponse(func(response *colly.Response) {
		fmt.Println(string(response.Body))
	})

	selectExamCollector.OnResponse(func(response *colly.Response) {
		fmt.Println(string(response.Body))
	})

	admissionTypeCollector.OnResponse(func(response *colly.Response) {
		fmt.Println(string(response.Body))
	})

	detailCollector.OnResponse(func(response *colly.Response) {
		fmt.Println(string(response.Body))
	})

	return c.Visit("http://admission3.jiangnan.edu.cn:3001/historyScore/nonArt")
}

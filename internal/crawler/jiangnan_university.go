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

var tokens = []string{
	"co1zLbSIrhyATJTOwEn9xJ4Jeq+xPYMsQiSyL8A6qCo/BuTU0zFgGDvhpAsBpgnJnPa8+A0eb+1uYdtIdGHOagZzCff39RaL7x6CSvNVFiHchWQtFf8iEYB7RvhAjH6LWHJFj+939pAePBgpbr/pOsTkwuQXbbEY5SbaGnMuiOE=",
	"dLzXgWVohH1EB76W0MgjH/SDmViecGLc1D/vIuk7RKIz10m29THxmLaFH2hXPxL1MlhMaRNzVqQ0muQ+Ofr0COccKBlfSlWZobnIpt5Lz7rc1yAvzBb9N087zBCfqxW9R5OawYql95VrPMJPdu47iX2S55uaZIBb3Hda8ge5yzk=",
	"kGJo/47FOVjxNinWtkWolnYM0RzHBHutOB9jzCn1pLSy7qjFagLEr/xt0qKhLL1qqlU+yW+1wsSNPLrAe+zSEleJc1X+WKdgRQOYqYdhpJ2mQmIjGP1SRXcsp58rwPxP9VPKkaCAbHdHECxUSRVZOalj3GEqTd6Ol0KHwph24TQ=",
	"ChUocB9taVQ+LG1okdeOCrAgLXkdrT2h0P2SzXjCG+rR+Sl4VV0uHofn/6mBWfq6BkQX7A9h0dvGCoSnm/ZYJ/CPaw0nq+YpcTLuTk7HoKrupt7pPrgtGYP1W5fGmpeGNIrriWrWkrR/AQZEZcpUpqK/acVRprBCt1SgtjJIITE=",
	"AJaZ1m2v902Rq9Od3q03EwL7pN2fpouyYlSCTr7s8lcNVvlLzqisp0z+SKfo8AT1zQVGSNyEhInR1OIoQyJBCEitqMp3k46rEpE5rRtqZpA2OMNMEk5fxJ6VFbNKrqRhkRcz4ypkkADffqa//d/c3zZHk2mOMTM4FbIF7eoEYhI=",
	"C5cecnLFYeFC40Zs6/zCWC1rKHSuwFTEoC+99mzMIHZXMRUk068Gugr+f1xHKDcXVVyrGr183lAKPVlHC0GTgSVjZQAKVc97zQh9Ob7LCxnavpRMoA0q9Q35RUxboImxDwQ1Cz/VBZHDXdhONaR/mKYS/UnvvHYxZWvumcuQSbU=",
	"M0RQ+Mtn0+W94yYP/AqK5/05S6t2DXqFrSDmYPsIvzli0zY8jK/0QdmSs8QnAJHv1vNfEd/L0cipnK3+YSDh/mnhwSoOffHl11c07RmjO9UpzFY/i8caje6xLF5AEsDxuVio7eHWcBidcN1mwlh1lBWZOzXQqqP2ihbXvlzeOss=",
	"GTFsm92ZDGnq21oR5vqUxKiXYmktagpGViKXrKQTesCCHLNmLWzr0k+YawqiSp57m4WveZWkttnuT8cXht85WbF+BmIGJYe5DHRTawPOMnCOg3abG3hEfhU25aDCAJIHI/0J/EuwUcrV0o/6C0iVvN+cbpmVshxmaYQIitH2uKg=",
	"FENpUksrcQdwCtXFq0wRVukSYLDkQb5yw861B1Y62qZj8NBbUtvtFX9V29tm9HXW2rGqfJpIFqaR0L6MK2MRA0EIICueO4uVHh+a6ClHOOrgR39BurE+p/Qpj/2eFKnF5ZgHLpJEaMgBySgFNBGWgeLk+rZi9Tuk9ABwa6uPjR8=",
	"gwHQPZ4a9g/9Ke7HYHzyIbIxGbISHQ6CPvn6SitbpzBS1mjZCZ5I63hEvUHa8UZk2wVGTJSlleQ1/UlkiwIDgKU8gRS+kvDjsRLqtU9lAFA0Xtl1uEHMTgP4ImQOtWop4Iale3KXXdk/J8rzUsBJlxS7duCc/+q7AHPIrc40OGw=",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
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

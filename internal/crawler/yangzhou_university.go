package crawler

import (
	"context"
	"fmt"

	"github.com/gocolly/colly/v2"
)

type yangzhouUniversity struct {
	university
}

func init() {
	collection["4132011117"] = &yangzhouUniversity{university{
		name: "扬州大学",
		code: "4132011117",
	}}
}

// todo: cookie

func (u *yangzhouUniversity) crawl(ctx context.Context) error {
	c := colly.NewCollector(colly.UserAgent(userAgent), colly.AllowURLRevisit())

	c.OnHTML(`div.right ul.newsList`, func(element *colly.HTMLElement) {
		element.ForEach("li", func(i int, element *colly.HTMLElement) {
			title := element.Text
			if fmt.Sprintf("%s年扬州大学江苏省录取信息", u.admissionTime) == title {
				fmt.Println(title)
			}

			if fmt.Sprintf("%s年扬州大学江苏省外录取信息", u.admissionTime) == title {
				fmt.Println(title)
			}
		})

		/*element.ForEach("tr", func(i int, element *colly.HTMLElement) {
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
		})*/
	})

	c.OnRequest(func(request *colly.Request) {
		request.Headers.Add("Host", "zhaoban.yzu.edu.cn")
		//request.Headers.Add("Referer", "https://zhaoban.yzu.edu.cn/bkcx/lncx.htm")
		request.Headers.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
		request.Headers.Add("Cookie", "JRthf9qDHS8VO=5Ff9LK33buBH9mZKz7RwH6YpRyWblWW1xetE7YS_h0XgngSG.xDAySO2YHr1Vt.dlNzhuLBH0xL1Ycha4X1aMpa; _ga=GA1.1.993161105.1686474751; JSESSIONID=9E70CC5FDCAF81B81CD85A90E490E741; _ga_2BJJGL0MED=GS1.1.1686482458.3.1.1686482459.0.0.0; JRthf9qDHS8VP=5REKXDC1m.f7qqqDEjOkwGaS.2Hds1bcRyFBnmAhPd7GDdJ77wNHzgG8wTuSztsjI6oYY1HI3lvr2Wcx.CvzzaEesu5HBXv9XQr5tKzfpUqar3i.QpwiDAfUmm3GKFLxASol8lLAGSEYGWm7BUewWPNaSzP9bCO3bbj_21ynPjlwIgCvXngaVrHBUiw8yiXUbWGwyywkOBxuO5Hl0KS18ALn4OKKf8_WtbvstX6GKu2wJQ2itlVAE4mEOEBPDro9QGn7zF8XhyL27ENyUyC_J0mbfRfGKYQhxO1W3du8PEn8A")
	})

	c.OnResponse(func(response *colly.Response) {
		fmt.Println("000---000")
	})

	return c.Visit("https://zhaoban.yzu.edu.cn/bkcx/lncx.htm")
}

package crawler

import (
	"context"
	"fmt"
	"strings"

	"github.com/rocky114/craftsman/internal/storage"

	"github.com/sirupsen/logrus"

	"github.com/rocky114/craftsman/internal/pkg/path"

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

func (u *yangzhouUniversity) crawl(ctx context.Context) error {
	cookie := "JRthf9qDHS8VO=5J5.RiK.gXGcF0iwHgZl2NQiMi1jI6l8yq59s_gw5hHJzUHqy3MjRdxgnJi8tC_tS97JYjypx.bE.FqR1EM2Dwq; _ga=GA1.1.1592715626.1687073902; _ga_2BJJGL0MED=GS1.1.1687073901.1.0.1687073901.0.0.0; JRthf9qDHS8VP=5RE_0MK1izD0qqqDEuUkWoqtZIFIOTPcZPpvNQJwtPy7cC3QBdljKCQ6O7jax1h4Px949cC4vsM5xdWw5nCjbbybmOezpMLAu.IPmJajVu16aJy0dpegN2EzgOc03nRF9o.tC2N8FUbdsEH9b1I8.IPWrf_3InSm1M7JvBLFEMpdaT5HEZrpwBevXG6RQv7hZAUKfV.EpGHvBNadFYKOm540AZfv0Z4UUPrnaBZCJW2XPyTsAv0IsXa7zVMlU2xavbGZ6m7_yaYakxpDOxTti7PkEIpF4V4z5ujDrapmxYn.A"
	c := colly.NewCollector(colly.UserAgent(userAgent), colly.CacheDir(path.GetTmpPath("yangzhouUniversity")))
	c.OnRequest(func(request *colly.Request) {
		request.Headers.Add("Host", "zhaoban.yzu.edu.cn")
		request.Headers.Add("Cookie", cookie)
	})

	detailCollector := c.Clone()

	detailCollector.OnRequest(func(request *colly.Request) {
		request.Headers.Add("Host", "zhaoban.yzu.edu.cn")
		request.Headers.Add("Cookie", cookie)
	})
	detailCollector.OnResponse(func(response *colly.Response) {
		fmt.Println(string(response.Body), "----")
	})
	c.OnHTML(`div.ny div.right ul.newsList`, func(element *colly.HTMLElement) {
		element.ForEach("li", func(i int, element *colly.HTMLElement) {
			title := element.ChildAttr("a", "title")
			if fmt.Sprintf("%s年扬州大学江苏省录取信息", u.admissionTime) == title {
				addr := strings.TrimPrefix(element.ChildAttr("a", "href"), "../")

				fmt.Println(fmt.Sprintf("https://zhaoban.yzu.edu.cn/%s", addr))
				if err := detailCollector.Visit(fmt.Sprintf("https://zhaoban.yzu.edu.cn/%s", addr)); err != nil {
					logrus.Errorf("yangzhouUniversity err: %v", err)
				}
			}

			if fmt.Sprintf("%s年扬州大学江苏省外录取信息", u.admissionTime) == title {
				fmt.Println(title)
			}
		})
	})

	detailCollector.OnHTML("div.ny div.right div.v_news_content table tbody", func(element *colly.HTMLElement) {
		province, admissionType, major, selectExam, admissionNumber, maxScore, minScore, averageScore := "江苏", "", "", "", "", "", "", ""

		fmt.Println("===")
		element.ForEach("tr", func(i int, element *colly.HTMLElement) {
			if i == 0 {
				return
			}

			length := element.DOM.Find("td").Length()
			admissionType = element.ChildText("td:nth-of-type(1)")
			selectExam = element.ChildText("td:nth-of-type(2)")
			major = element.ChildText("td:nth-of-type(4)")
			admissionNumber = element.ChildText(fmt.Sprintf("td:nth-of-type(%d)", length-4))
			maxScore = element.ChildText(fmt.Sprintf("td:nth-of-type(%d)", length-3))
			minScore = element.ChildText(fmt.Sprintf("td:nth-of-type(%d)", length-2))
			averageScore = element.ChildText(fmt.Sprintf("td:nth-of-type(%d)", length-1))

			if err := storage.GetQueries().CreateAdmissionMajor(context.Background(), storage.CreateAdmissionMajorParams{
				University:      u.name,
				AdmissionTime:   u.admissionTime,
				Major:           major,
				Province:        province,
				AdmissionType:   admissionType,
				SelectExam:      selectExam,
				AdmissionNumber: admissionNumber,
				MinScore:        minScore,
				MaxScore:        maxScore,
				AverageScore:    averageScore,
			}); err != nil {
				logrus.Errorf("create admission major err: %v", err)
			}
		})
	})

	return c.Visit("https://zhaoban.yzu.edu.cn/bkcx/lncx.htm")
}

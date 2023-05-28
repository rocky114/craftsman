package crawler

import (
	"context"
	"time"

	"github.com/rocky114/craftsman/internal/pkg/path"

	"github.com/gocolly/colly/v2"
)

type jiangsukejiUniversity struct {
	university
}

func init() {
	collection["4132010289"] = &jiangsukejiUniversity{university{
		name: "江苏科技大学",
		code: "4132010289",
	}}
}

func (u *jiangsukejiUniversity) crawl(ctx context.Context) error {
	c := colly.NewCollector(colly.CacheDir(path.GetTmpPath("jiangsukejiUniversity")))
	c.SetRequestTimeout(60 * time.Second)

	c.OnHTML(`map[id=Map]`, func(element *colly.HTMLElement) {

	})

	return c.Visit("http://zs.just.edu.cn/1603/list.htm")
}

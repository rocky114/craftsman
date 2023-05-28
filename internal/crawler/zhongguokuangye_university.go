package crawler

import (
	"context"
	"time"

	"github.com/rocky114/craftsman/internal/pkg/path"

	"github.com/gocolly/colly/v2"
)

type zhongguokuangyeUniversity struct {
	university
}

func init() {
	collection["4132010290"] = &zhongguokuangyeUniversity{university{
		name: "中国矿业大学",
		code: "4132010290",
	}}
}

// todo: 暂无2022数据
func (u *zhongguokuangyeUniversity) crawl(ctx context.Context) error {
	c := colly.NewCollector(colly.CacheDir(path.GetTmpPath("jiangsukejiUniversity")))
	c.SetRequestTimeout(60 * time.Second)

	c.OnHTML(`map[id=Map]`, func(element *colly.HTMLElement) {

	})

	return c.Visit("https://zs.cumt.edu.cn/lnfs_16011/list1.htm")
}

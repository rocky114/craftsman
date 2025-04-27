package utils

const (
	PageSize = 10
)

type Pagination[T any] struct {
	List       []T   `json:"list"`        // 当前页数据列表
	TotalCount int64 `json:"total_count"` // 总记录数
	Page       int   `json:"page"`        // 当前页码
	PageSize   int   `json:"page_size"`   // 每页条数
}

func Offset(page int) int {
	if page < 1 {
		page = 1
	}

	return (page - 1) * PageSize
}

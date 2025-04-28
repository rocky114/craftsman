package dto

import (
	"github.com/rocky114/craftman/internal/database/sqlc"
)

type ScoreDistribution struct {
	ID int32 `db:"id"`
	// 年份
	Year string `db:"year"`
	// 省份
	Province string `db:"province"`
	// 物理等科目类, 历史等科目类
	SubjectCategory string `db:"subject_category"`
	// 分数段
	ScoreRange string `db:"score_range"`
	// 同分人数
	SameScoreCount string `db:"same_score_count"`
	// 累计人数
	CumulativeCount string `db:"cumulative_count"`
}

func TransformListScoreDistributionsResponse(items []sqlc.ScoreDistribution) []ScoreDistribution {
	ret := make([]ScoreDistribution, 0, len(items))
	for _, item := range items {
		ret = append(ret, ScoreDistribution{
			ID:              item.ID,
			Year:            item.Year,
			Province:        item.Province,
			SubjectCategory: item.SubjectCategory,
			ScoreRange:      item.ScoreRange,
			SameScoreCount:  item.SameScoreCount,
			CumulativeCount: item.CumulativeCount,
		})
	}

	return ret
}

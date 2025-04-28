package dto

import (
	"github.com/rocky114/craftman/internal/database/sqlc"
)

type AdmissionSummary struct {
	ID uint32 `db:"id"`
	// 录取年份
	Year string `db:"year"`
	// 省份 江苏
	Province string `db:"province"`
	// 高校名称
	UniversityName string `db:"university_name"`
	// 类型
	AdmissionType string `db:"admission_type"`
	// 科类
	SubjectCategory string `db:"subject_category"`
	// 全校最高分
	HighestScore string `db:"highest_score"`
	// 最高分位次
	HighestScoreRank string `db:"highest_score_rank"`
	// 全校最低分
	LowestScore string `db:"lowest_score"`
	// 最低分位次
	LowestScoreRank string `db:"lowest_score_rank"`
}

func TransformListAdmissionSummariesResponse(items []sqlc.AdmissionSummary) []AdmissionSummary {
	ret := make([]AdmissionSummary, 0, len(items))
	for _, item := range items {
		ret = append(ret, AdmissionSummary{
			ID:               item.ID,
			Year:             item.Year,
			UniversityName:   item.UniversityName,
			AdmissionType:    item.AdmissionType,
			SubjectCategory:  item.SubjectCategory,
			HighestScore:     item.HighestScore,
			HighestScoreRank: item.HighestScoreRank,
			LowestScore:      item.LowestScore,
			LowestScoreRank:  item.LowestScoreRank,
		})
	}

	return ret
}

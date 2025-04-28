package dto

import (
	"github.com/rocky114/craftman/internal/database/sqlc"
)

type AdmissionSummary struct {
	ID uint32 `json:"id"`
	// 录取年份
	Year string `json:"year"`
	// 省份 江苏
	Province string `json:"province"`
	// 高校名称
	UniversityName string `json:"university_name"`
	// 类型
	AdmissionType string `json:"admission_type"`
	// 科类
	SubjectCategory string `json:"subject_category"`
	// 全校最高分
	HighestScore string `json:"highest_score"`
	// 最高分位次
	HighestScoreRank string `json:"highest_score_rank"`
	// 全校最低分
	LowestScore string `json:"lowest_score"`
	// 最低分位次
	LowestScoreRank string `json:"lowest_score_rank"`
}

func TransformListAdmissionSummariesResponse(items []sqlc.AdmissionSummary) []AdmissionSummary {
	ret := make([]AdmissionSummary, 0, len(items))
	for _, item := range items {
		ret = append(ret, AdmissionSummary{
			ID:               item.ID,
			Year:             item.Year,
			Province:         item.Province,
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

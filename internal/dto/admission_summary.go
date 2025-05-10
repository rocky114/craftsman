package dto

import (
	"github.com/rocky114/craftman/internal/database/sqlc"
	"github.com/rocky114/craftman/internal/utils"
)

type AdmissionSummaryResponse struct {
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

func ToAdmissionSummaryResponses(items []sqlc.AdmissionSummary) []AdmissionSummaryResponse {
	ret := make([]AdmissionSummaryResponse, 0, len(items))
	for _, item := range items {
		ret = append(ret, AdmissionSummaryResponse{
			ID:               item.ID,
			Year:             item.Year,
			Province:         item.Province,
			UniversityName:   item.UniversityName,
			AdmissionType:    item.AdmissionType,
			SubjectCategory:  item.SubjectCategory,
			HighestScore:     item.HighestScore,
			HighestScoreRank: utils.Ternary[string](item.HighestScoreRank != "", item.HighestScoreRank, "无"),
			LowestScore:      item.LowestScore,
			LowestScoreRank:  utils.Ternary[string](item.LowestScoreRank != "", item.LowestScoreRank, "无"),
		})
	}

	return ret
}

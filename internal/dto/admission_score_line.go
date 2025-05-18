package dto

import (
	"github.com/rocky114/craftman/internal/database/sqlc"
	"github.com/rocky114/craftman/internal/utils"
)

type AdmissionScoreLineResponse struct {
	ID uint32 `json:"id"`
	// 录取年份（如2024）
	Year string `json:"year"`
	// 省份 江苏
	Province string `json:"province"`
	// 关联院校表
	UniversityName string `json:"university_name"`
	// 录取批次
	AdmissionBatch string `json:"admission_batch"`
	// 类型: 普通类,艺术类,国家专项,高校专项,中外合作,飞行技术,预科
	AdmissionType string `json:"admission_type"`
	// 科类: 历史+不限
	SubjectCategory string `json:"subject_category"`
	// 专业组
	MajorGroup string `json:"major_group"`
	// 最低分
	LowestScore string `json:"lowest_score"`
	// 排名 200000名次
	LowestScoreRank string `json:"lowest_score_rank"`
}

func ToAdmissionScoreLineResponses(items []sqlc.AdmissionScoreLine) []AdmissionScoreLineResponse {
	ret := make([]AdmissionScoreLineResponse, 0, len(items))
	for _, item := range items {
		ret = append(ret, AdmissionScoreLineResponse{
			ID:              item.ID,
			Year:            item.Year,
			Province:        item.Province,
			UniversityName:  item.UniversityName,
			AdmissionBatch:  item.AdmissionBatch,
			AdmissionType:   item.AdmissionType,
			SubjectCategory: item.SubjectCategory,
			MajorGroup:      item.MajorGroup,
			LowestScore:     item.LowestScore,
			LowestScoreRank: utils.Ternary[string](item.LowestScoreRank != "", item.LowestScoreRank, "无"),
		})
	}

	return ret
}

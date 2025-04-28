package dto

import (
	"github.com/rocky114/craftman/internal/database/sqlc"
)

type AdmissionScoreResponse struct {
	ID uint32 `json:"id"`
	// 录取年份（如2024）
	Year string `json:"year"`
	// 关联院校表
	UniversityName string `json:"university_name"`
	// 省份 江苏
	Province string `json:"province"`
	// 类型: 普通类,艺术类,国家专项,高校专项,中外合作,飞行技术,预科
	AdmissionType string `json:"admission_type"`
	// 科类文本
	SubjectCategoryTxt string `json:"subject_category_txt"`
	// 专业名称 计算机
	MajorName string `json:"major_name"`
	// 最高分
	HighestScore string `json:"highest_score"`
	// 排名 200000名次
	HighestScoreRank string `json:"highest_score_rank"`
	// 最低分
	LowestScore string `json:"lowest_score"`
	// 排名 200000名次
	LowestScoreRank string `json:"lowest_score_rank"`
}

func ToAdmissionScoreResponses(items []sqlc.AdmissionScore) []AdmissionScoreResponse {
	ret := make([]AdmissionScoreResponse, 0, len(items))
	for _, item := range items {
		ret = append(ret, AdmissionScoreResponse{
			ID:                 item.ID,
			Year:               item.Year,
			Province:           item.Province,
			UniversityName:     item.UniversityName,
			AdmissionType:      item.AdmissionType,
			SubjectCategoryTxt: item.SubjectCategoryTxt,
			MajorName:          item.MajorName,
			HighestScore:       item.HighestScore,
			HighestScoreRank:   item.HighestScoreRank,
			LowestScore:        item.LowestScore,
			LowestScoreRank:    item.LowestScoreRank,
		})
	}

	return ret
}

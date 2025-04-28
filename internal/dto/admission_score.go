package dto

import (
	"github.com/rocky114/craftman/internal/database/sqlc"
)

type AdmissionScore struct {
	ID uint32 `db:"id"`
	// 录取年份（如2024）
	Year string `db:"year"`
	// 关联院校表
	UniversityName string `db:"university_name"`
	// 省份 江苏
	Province string `db:"province"`
	// 类型: 普通类,艺术类,国家专项,高校专项,中外合作,飞行技术,预科
	AdmissionType string `db:"admission_type"`
	// 科类文本
	SubjectCategoryTxt string `db:"subject_category_txt"`
	// 专业名称 计算机
	MajorName string `db:"major_name"`
	// 最高分
	HighestScore string `db:"highest_score"`
	// 排名 200000名次
	HighestScoreRank string `db:"highest_score_rank"`
	// 最低分
	LowestScore string `db:"lowest_score"`
	// 排名 200000名次
	LowestScoreRank string `db:"lowest_score_rank"`
}

func TransformListAdmissionScoresResponse(items []sqlc.AdmissionScore) []AdmissionScore {
	ret := make([]AdmissionScore, 0, len(items))
	for _, item := range items {
		ret = append(ret, AdmissionScore{
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

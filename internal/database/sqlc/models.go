// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package sqlc

import (
	"database/sql"
	"time"
)

// 高校录取查询条件表
type AdmissionQueryCondition struct {
	// 主键ID
	ID int32 `db:"id"`
	// 学校名称
	UniversityName string `db:"university_name"`
	// 查询URL
	Url string `db:"url"`
	// 查询年份
	Year string `db:"year"`
	// 省份
	Province string `db:"province"`
	// 录取类型(统招计划/专项计划等)
	AdmissionType string `db:"admission_type"`
	// 创建时间
	CreateTime sql.NullTime `db:"create_time"`
}

// 高校专业录取分数
type AdmissionScore struct {
	ID uint32 `db:"id"`
	// 录取年份（如2024）
	Year string `db:"year"`
	// 省份 江苏
	Province string `db:"province"`
	// 关联院校表
	UniversityName string `db:"university_name"`
	// 类型: 普通类,艺术类,国家专项,高校专项,中外合作,飞行技术,预科
	AdmissionType string `db:"admission_type"`
	// 科类: 历史+不限
	SubjectCategory string `db:"subject_category"`
	// 专业名称 计算机
	MajorName string `db:"major_name"`
	// 最高分
	HighestScore string `db:"highest_score"`
	// 排名 200000名次
	HighestScoreRank string `db:"highest_score_rank"`
	// 最低分
	LowestScore string `db:"lowest_score"`
	// 排名 200000名次
	LowestScoreRank string    `db:"lowest_score_rank"`
	CreateTime      time.Time `db:"create_time"`
}

// 考试院公布的投档线
type AdmissionScoreLine struct {
	ID uint32 `db:"id"`
	// 录取年份（如2024）
	Year string `db:"year"`
	// 省份 江苏
	Province string `db:"province"`
	// 关联院校表
	UniversityName string `db:"university_name"`
	// 录取批次
	AdmissionBatch string `db:"admission_batch"`
	// 类型: 普通类,艺术类,国家专项,高校专项,中外合作,飞行技术,预科
	AdmissionType string `db:"admission_type"`
	// 定向区域
	AdmissionRegion string `db:"admission_region"`
	// 科类: 历史+不限
	SubjectCategory string `db:"subject_category"`
	// 专业组
	MajorGroup string `db:"major_group"`
	// 最低分
	LowestScore string `db:"lowest_score"`
	// 排名 200000名次
	LowestScoreRank string    `db:"lowest_score_rank"`
	CreateTime      time.Time `db:"create_time"`
}

// 高校录取分数汇总表
type AdmissionSummary struct {
	ID uint32 `db:"id"`
	// 录取年份
	Year string `db:"year"`
	// 省份
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
	LowestScoreRank string    `db:"lowest_score_rank"`
	CreateTime      time.Time `db:"create_time"`
}

// 一段一分表
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

// 学校基础信息表
type University struct {
	ID uint32 `db:"id"`
	// 学校名称
	Name string `db:"name"`
	// 省份
	Province string `db:"province"`
	// 招生网址
	AdmissionWebsite string    `db:"admission_website"`
	CreateTime       time.Time `db:"create_time"`
	UpdateTime       time.Time `db:"update_time"`
}

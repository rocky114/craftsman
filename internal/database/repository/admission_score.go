package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/rocky114/craftman/internal/database/sqlc"
)

type AdmissionScoreQueryParams struct {
	Year            string `json:"year"`
	UniversityName  string `json:"university_name"`
	Province        string `json:"province"`
	AdmissionType   string `json:"admission_type"`
	SubjectCategory string `json:"subject_category"`
	Limit           int    `json:"limit"`
	Offset          int    `json:"offset"`
}

func (q *Repository) buildAdmissionScoreQuery(baseQuery string, arg AdmissionScoreQueryParams) (string, []interface{}) {
	var queryBuilder strings.Builder
	queryBuilder.WriteString(baseQuery)

	args := make([]interface{}, 0)
	conditions := make([]string, 0)

	if arg.Year != "" {
		conditions = append(conditions, "year = ?")
		args = append(args, arg.Year)
	}
	if arg.UniversityName != "" {
		conditions = append(conditions, "university_name LIKE ?")
		args = append(args, "%"+arg.UniversityName+"%")
	}
	if arg.Province != "" {
		conditions = append(conditions, "province = ?")
		args = append(args, arg.Province)
	}
	if arg.AdmissionType != "" {
		conditions = append(conditions, "admission_type = ?")
		args = append(args, arg.AdmissionType)
	}
	if arg.SubjectCategory != "" {
		conditions = append(conditions, "subject_category = ?")
		args = append(args, arg.SubjectCategory)
	}

	if len(conditions) > 0 {
		queryBuilder.WriteString(" WHERE " + strings.Join(conditions, " AND "))
	}

	return queryBuilder.String(), args
}

func (q *Repository) ListAdmissionScores(ctx context.Context, arg AdmissionScoreQueryParams) ([]sqlc.AdmissionScore, error) {
	baseQuery := "SELECT id, year, university_name, province, admission_type, subject_category, subject_category_txt, major_name, enrollment_quota, min_admission_score, highest_score, highest_score_rank, lowest_score, lowest_score_rank, create_time FROM admission_score"
	query, args := q.buildAdmissionScoreQuery(baseQuery, arg)

	query += " ORDER BY id ASC"
	if arg.Limit > 0 {
		query += " LIMIT ?"
		args = append(args, arg.Limit)
	}
	if arg.Offset > 0 {
		query += " OFFSET ?"
		args = append(args, arg.Offset)
	}

	var items []sqlc.AdmissionScore
	if err := q.db.SelectContext(ctx, &items, query, args...); err != nil {
		return nil, fmt.Errorf("ListAdmissionScores failed: %w", err)
	}
	return items, nil
}

func (q *Repository) CountAdmissionScores(ctx context.Context, arg AdmissionScoreQueryParams) (int64, error) {
	baseQuery := "SELECT COUNT(*) AS total_count FROM admission_score"
	query, args := q.buildAdmissionScoreQuery(baseQuery, arg)

	var total int64
	if err := q.db.GetContext(ctx, &total, query, args...); err != nil {
		return 0, fmt.Errorf("CountAdmissionScores failed: %w", err)
	}
	return total, nil
}

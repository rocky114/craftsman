package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/rocky114/craftman/internal/database/sqlc"
)

type AdmissionSummaryQueryParams struct {
	Year            string `json:"year"`
	UniversityName  string `json:"university_name"`
	AdmissionType   string `json:"admission_type"`
	SubjectCategory string `json:"subject_category"`
	Limit           int    `json:"limit"`
	Offset          int    `json:"offset"`
}

func (q *Repository) buildAdmissionSummaryQuery(baseQuery string, arg AdmissionSummaryQueryParams) (string, []interface{}) {
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

func (q *Repository) ListAdmissionSummaries(ctx context.Context, arg AdmissionSummaryQueryParams) ([]sqlc.AdmissionSummary, error) {
	baseQuery := "SELECT id, year, province, university_name, admission_type, subject_category, highest_score, highest_score_rank, lowest_score, lowest_score_rank, create_time FROM admission_summary"
	query, args := q.buildAdmissionSummaryQuery(baseQuery, arg)

	query += " ORDER BY lowest_score desc"
	if arg.Limit > 0 {
		query += " LIMIT ?"
		args = append(args, arg.Limit)
	}
	if arg.Offset > 0 {
		query += " OFFSET ?"
		args = append(args, arg.Offset)
	}

	var items []sqlc.AdmissionSummary
	if err := q.db.SelectContext(ctx, &items, query, args...); err != nil {
		return nil, fmt.Errorf("ListAdmissionSummaries failed: %w", err)
	}
	return items, nil
}

func (q *Repository) CountAdmissionSummaries(ctx context.Context, arg AdmissionSummaryQueryParams) (int64, error) {
	baseQuery := "SELECT COUNT(*) AS total_count FROM admission_summary"
	query, args := q.buildAdmissionSummaryQuery(baseQuery, arg)

	var total int64
	if err := q.db.GetContext(ctx, &total, query, args...); err != nil {
		return 0, fmt.Errorf("CountAdmissionSummaries failed: %w", err)
	}
	return total, nil
}

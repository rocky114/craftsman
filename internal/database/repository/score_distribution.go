package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/rocky114/craftman/internal/database/sqlc"
)

type ScoreDistributionQueryParams struct {
	Year            string `json:"year"`
	Province        string `json:"province"`
	SubjectCategory string `json:"subject_category"`
	ScoreRange      string `json:"score_range"`
	Limit           int    `json:"limit"`
	Offset          int    `json:"offset"`
}

func (q *Repository) buildScoreDistributionQuery(baseQuery string, arg ScoreDistributionQueryParams) (string, []interface{}) {
	var queryBuilder strings.Builder
	queryBuilder.WriteString(baseQuery)

	args := make([]interface{}, 0)
	conditions := make([]string, 0)

	if arg.Year != "" {
		conditions = append(conditions, "year = ?")
		args = append(args, arg.Year)
	}
	if arg.Province != "" {
		conditions = append(conditions, "province = ?")
		args = append(args, arg.Province)
	}
	if arg.SubjectCategory != "" {
		conditions = append(conditions, "subject_category = ?")
		args = append(args, arg.SubjectCategory)
	}
	if arg.ScoreRange != "" {
		conditions = append(conditions, "score_range = ?")
		args = append(args, arg.ScoreRange)
	}

	if len(conditions) > 0 {
		queryBuilder.WriteString(" WHERE " + strings.Join(conditions, " AND "))
	}

	return queryBuilder.String(), args
}

func (q *Repository) ListScoreDistributions(ctx context.Context, arg ScoreDistributionQueryParams) ([]sqlc.ScoreDistribution, error) {
	baseQuery := "SELECT id, year, province, subject_category, score_range, same_score_count, cumulative_count FROM score_distribution"
	query, args := q.buildScoreDistributionQuery(baseQuery, arg)

	query += " ORDER BY id ASC"
	if arg.Limit > 0 {
		query += " LIMIT ?"
		args = append(args, arg.Limit)
	}
	if arg.Offset > 0 {
		query += " OFFSET ?"
		args = append(args, arg.Offset)
	}

	var items []sqlc.ScoreDistribution
	if err := q.db.SelectContext(ctx, &items, query, args...); err != nil {
		return nil, fmt.Errorf("ListScoreDistributions failed: %w", err)
	}
	return items, nil
}

func (q *Repository) CountScoreDistributions(ctx context.Context, arg ScoreDistributionQueryParams) (int64, error) {
	baseQuery := "SELECT COUNT(*) AS total_count FROM score_distribution"
	query, args := q.buildScoreDistributionQuery(baseQuery, arg)

	var total int64
	if err := q.db.GetContext(ctx, &total, query, args...); err != nil {
		return 0, fmt.Errorf("CountScoreDistributions failed: %w", err)
	}
	return total, nil
}

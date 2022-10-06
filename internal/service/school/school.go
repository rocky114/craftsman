package school

import (
	"context"

	"github.com/rocky114/craftsman/internal/storage"
)

type ListSchoolResponse struct {
	Count int64            `json:"count"`
	Rows  []storage.School `json:"rows"`
}

func ListSchool(req storage.ListSchoolsParams) (*ListSchoolResponse, error) {
	schools, err := storage.GetQueries().ListSchools(context.Background(), req)
	if err != nil {
		return nil, err
	}
	count, err := storage.GetQueries().CountSchools(context.Background())
	if err != nil {
		return nil, err
	}

	return &ListSchoolResponse{Count: count, Rows: schools}, nil
}

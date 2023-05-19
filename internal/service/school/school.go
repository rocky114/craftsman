package school

import (
	"context"

	"github.com/rocky114/craftsman/internal/storage"
)

type ListUniversitiesResponse struct {
	Total int64                `json:"total"`
	Items []storage.University `json:"items"`
}

func ListSchool(req storage.ListUniversitiesParams) (*ListUniversitiesResponse, error) {
	schools, err := storage.GetQueries().ListUniversities(context.Background(), req)
	if err != nil {
		return nil, err
	}
	count, err := storage.GetQueries().CountUniversities(context.Background())
	if err != nil {
		return nil, err
	}

	return &ListUniversitiesResponse{Total: count, Items: schools}, nil
}

package school

import (
	"context"

	"github.com/rocky114/craftsman/internal/storage"
)

func ListSchool(req storage.ListSchoolParams) ([]storage.School, error) {
	return storage.GetQueries().ListSchool(context.Background(), req)
}

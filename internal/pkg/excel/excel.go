package excel

import (
	"fmt"
	"log"

	"github.com/rocky114/craftsman/internal/storage"
	"github.com/xuri/excelize/v2"
)

func GetSchools(filename string) []storage.CreateSchoolParams {
	f, err := excelize.OpenFile(filename)
	if err != nil {
		log.Fatalf("open excel err: %v", err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	rows, err := f.GetRows("sheet1")
	if err != nil {
		log.Fatalf("read excel row err: %v", err)
	}

	schools := make([]storage.CreateSchoolParams, 0)
	for i, columns := range rows {
		if i < 4 || len(columns) <= 2 {
			continue
		}

		var remark string
		if len(columns) == 7 {
			remark = columns[6]
		}
		schools = append(schools, storage.CreateSchoolParams{
			Name:       columns[1],
			Code:       columns[2],
			Department: columns[3],
			Location:   columns[4],
			Level:      columns[5],
			Remark:     remark,
		})
	}

	return schools
}

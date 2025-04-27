package dto

import "github.com/rocky114/craftman/internal/database/sqlc"

type UniversityResponse struct {
	ID uint32 `json:"id"`
	// 学校名称
	Name string `json:"name"`
	// 省份
	Province string `json:"province"`
	// 招生网址
	AdmissionWebsite string `json:"admission_website"`
}

func ToUniversityResponse(items []sqlc.University) []UniversityResponse {
	ret := make([]UniversityResponse, 0, len(items))
	for _, item := range items {
		ret = append(ret, UniversityResponse{
			ID:               item.ID,
			Name:             item.Name,
			Province:         item.Province,
			AdmissionWebsite: item.AdmissionWebsite,
		})
	}

	return ret
}

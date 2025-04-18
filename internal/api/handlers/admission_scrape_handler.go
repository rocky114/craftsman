package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/rocky114/craftman/internal/app/config"
	"github.com/rocky114/craftman/internal/database"
	"github.com/rocky114/craftman/internal/database/sqlc"
	"github.com/rocky114/craftman/internal/utils"
	"net/http"
)

type AdmissionScrapeHandler struct {
	repo *database.Repository
	cfg  *config.Config
}

func NewAdmissionScrapeHandler(q *database.Repository, cfg *config.Config) *AdmissionScrapeHandler {
	return &AdmissionScrapeHandler{repo: q, cfg: cfg}
}

func (h *AdmissionScrapeHandler) CreateAdmissionScore(c echo.Context) error {
	var admissionScore struct {
		Name string `query:"name"`
		Year string `query:"year"`
	}

	if err := c.Bind(&admissionScore); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	condition, err := h.repo.GetQueryConditionByYearAndName(c.Request().Context(), sqlc.GetQueryConditionByYearAndNameParams{
		Year:           admissionScore.Year,
		UniversityName: admissionScore.Name,
	})
	if err != nil {
		return utils.Error(c, http.StatusNotFound, fmt.Sprintf("failed to query condition: %s", err.Error()))
	}

	respAdmission, err := utils.FetchAdmissionData(h.cfg.Scraper.URL, utils.AdmissionRequest{
		URL:            condition.Url,
		Year:           condition.Year,
		Province:       condition.Province,
		AdmissionType:  condition.AdmissionType,
		UniversityName: condition.UniversityName,
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if respAdmission.Status == utils.AdmissionRespStatusErr {
		return echo.NewHTTPError(http.StatusInternalServerError, respAdmission.Message)
	}

	err = h.repo.WithTransaction(c.Request().Context(), func(q *sqlc.Queries) error {
		for _, item := range respAdmission.Data {
			if err = q.CreateAdmissionScore(c.Request().Context(), sqlc.CreateAdmissionScoreParams{
				UniversityName:    condition.UniversityName,
				Province:          item.Province,
				Year:              item.Year,
				AdmissionType:     item.AdmissionType,
				AcademicCategory:  item.AcademicCategory,
				MajorName:         item.MajorName,
				EnrollmentQuota:   item.EnrollmentQuota,
				MinAdmissionScore: item.MinAdmissionScore,
				HighestScore:      item.HighestScore,
				HighestScoreRank:  item.HighestScoreRank,
				LowestScore:       item.LowestScore,
				LowestScoreRank:   item.LowestScoreRank,
			}); err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return utils.Success(c)
}

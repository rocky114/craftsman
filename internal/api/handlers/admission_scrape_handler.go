package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/rocky114/craftman/internal/database/sqlc"
	"github.com/rocky114/craftman/internal/utils"
	"net/http"
)

type AdmissionScrapeHandler struct {
	queries *sqlc.Queries
}

func NewAdmissionScrapeHandler(q *sqlc.Queries) *AdmissionScrapeHandler {
	return &AdmissionScrapeHandler{queries: q}
}

func (h *AdmissionScrapeHandler) CreateAdmissionScore(c echo.Context) error {
	var admissionScore struct {
		Name string `query:"name"`
		Year string `query:"year"`
	}

	if err := c.Bind(&admissionScore); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	condition, err := h.queries.GetQueryConditionByYearAndName(c.Request().Context(), sqlc.GetQueryConditionByYearAndNameParams{
		Year:           admissionScore.Year,
		UniversityName: admissionScore.Name,
	})
	if err != nil {
		return utils.Error(c, http.StatusNotFound, "query condition not found")
	}

	respAdmission, err := utils.FetchAdmissionData("", utils.AdmissionRequest{
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

	err = h.queries.CreateAdmissionScore(c.Request().Context(), sqlc.CreateAdmissionScoreParams{
		Province: admissionScore.Name,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return utils.Success(c)
}

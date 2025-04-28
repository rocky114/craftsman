package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/rocky114/craftman/internal/app/config"
	"github.com/rocky114/craftman/internal/database"
	"github.com/rocky114/craftman/internal/database/repository"
	"github.com/rocky114/craftman/internal/dto"
	"github.com/rocky114/craftman/internal/utils"
	"net/http"
)

type AdmissionSummaryHandler struct {
	repo *database.Database
	cfg  *config.Config
}

func NewAdmissionSummaryHandler(q *database.Database, cfg *config.Config) *AdmissionSummaryHandler {
	return &AdmissionSummaryHandler{repo: q, cfg: cfg}
}

func (h *AdmissionSummaryHandler) ListAdmissionSummaries(c echo.Context) error {
	var req struct {
		Page int `json:"page" form:"page" query:"page"` // 当前页码（从1开始）
	}

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	items, err := h.repo.ListAdmissionSummaries(c.Request().Context(), repository.AdmissionSummaryQueryParams{
		Limit:  utils.PageSize,
		Offset: utils.Offset(req.Page),
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	totalCount, err := h.repo.CountAdmissionSummaries(c.Request().Context(), repository.AdmissionSummaryQueryParams{})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	ret := utils.Pagination[dto.AdmissionSummary]{
		List:       dto.TransformListAdmissionSummariesResponse(items),
		TotalCount: totalCount,
		Page:       req.Page,
		PageSize:   utils.PageSize,
	}

	return c.JSON(http.StatusOK, ret)
}

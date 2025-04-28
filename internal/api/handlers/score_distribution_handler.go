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

type ScoreDistributionHandler struct {
	repo *database.Database
	cfg  *config.Config
}

func NewScoreDistributionHandler(q *database.Database, cfg *config.Config) *ScoreDistributionHandler {
	return &ScoreDistributionHandler{repo: q, cfg: cfg}
}

func (h *ScoreDistributionHandler) ListScoreDistributions(c echo.Context) error {
	var req struct {
		Page int `json:"page" form:"page" query:"page"` // 当前页码（从1开始）
	}

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	items, err := h.repo.ListScoreDistributions(c.Request().Context(), repository.ScoreDistributionQueryParams{
		Limit:  utils.PageSize,
		Offset: utils.Offset(req.Page),
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	totalCount, err := h.repo.CountScoreDistributions(c.Request().Context(), repository.ScoreDistributionQueryParams{})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	ret := utils.Pagination[dto.ScoreDistributionResponse]{
		List:       dto.ToScoreDistributionResponses(items),
		TotalCount: totalCount,
		Page:       req.Page,
		PageSize:   utils.PageSize,
	}

	return c.JSON(http.StatusOK, ret)
}

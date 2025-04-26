package handlers

import (
	"github.com/rocky114/craftman/internal/app/config"
	"github.com/rocky114/craftman/internal/database"
	"github.com/rocky114/craftman/internal/database/sqlc"
	"github.com/rocky114/craftman/internal/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UniversityHandler struct {
	repo *database.Database
	cfg  *config.Config
}

func NewUniversityHandler(q *database.Database, cfg *config.Config) *UniversityHandler {
	return &UniversityHandler{repo: q, cfg: cfg}
}

func (h *UniversityHandler) CreateUniversity(c echo.Context) error {
	var university struct {
		Name             string `json:"name"`
		Province         string `json:"province"`
		AdmissionWebsite string `json:"admission_website"`
	}

	if err := c.Bind(&university); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := h.repo.CreateUniversity(c.Request().Context(), sqlc.CreateUniversityParams{
		Name:             university.Name,
		Province:         university.Province,
		AdmissionWebsite: university.AdmissionWebsite,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return utils.Success(c)
}

func (h *UniversityHandler) GetUniversity(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid user ID")
	}

	item, err := h.repo.GetUniversity(c.Request().Context(), uint32(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}

	return utils.SuccessWithData(c, item)
}

func (h *UniversityHandler) ListUniversities(c echo.Context) error {
	items, err := h.repo.ListUniversities(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return utils.SuccessWithData(c, items)
}

func (h *UniversityHandler) DeleteUniversity(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid user ID")
	}

	err = h.repo.DeleteUniversity(c.Request().Context(), uint32(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return utils.Success(c)
}

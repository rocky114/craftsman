package handlers

import (
	"github.com/rocky114/craftman/internal/database/sqlc"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UniversityHandler struct {
	queries *sqlc.Queries
}

func NewUniversityHandler(q *sqlc.Queries) *UniversityHandler {
	return &UniversityHandler{queries: q}
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

	err := h.queries.CreateUniversity(c.Request().Context(), sqlc.CreateUniversityParams{
		Name:             university.Name,
		Province:         university.Province,
		AdmissionWebsite: university.AdmissionWebsite,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	//id, _ := result.LastInsertId()
	return c.JSON(http.StatusCreated, map[string]int64{"id": 0})
}

func (h *UniversityHandler) GetUniversity(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid user ID")
	}

	item, err := h.queries.GetUniversity(c.Request().Context(), uint32(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}

	return c.JSON(http.StatusOK, item)
}

func (h *UniversityHandler) ListUniversities(c echo.Context) error {
	items, err := h.queries.ListUniversities(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, items)
}

func (h *UniversityHandler) DeleteUniversity(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid user ID")
	}

	err = h.queries.DeleteUniversity(c.Request().Context(), uint32(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

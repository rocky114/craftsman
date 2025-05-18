package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/rocky114/craftman/internal/app/config"
	"github.com/rocky114/craftman/internal/database"
	"github.com/rocky114/craftman/internal/database/repository"
	"github.com/rocky114/craftman/internal/database/sqlc"
	"github.com/rocky114/craftman/internal/dto"
	"github.com/rocky114/craftman/internal/utils"
	"net/http"
	"strings"
)

type AdmissionScoreLineHandler struct {
	repo *database.Database
	cfg  *config.Config
}

func NewAdmissionScoreLineHandler(q *database.Database, cfg *config.Config) *AdmissionScoreLineHandler {
	return &AdmissionScoreLineHandler{repo: q, cfg: cfg}
}

func (h *AdmissionScoreLineHandler) CreateAdmissionScoreLine(c echo.Context) error {
	var admissionScoreLineReq struct {
		Url            string `json:"url"`
		UniversityName string `json:"universityName"`
		Year           int    `json:"year"`
	}

	var admissionScoreResp struct {
		Total int `json:"total"`
	}

	if err := c.Bind(&admissionScoreLineReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	respAdmission, err := utils.FetchAdmissionScoreLineData(h.cfg.Scraper.URL, utils.AdmissionScoreLineRequest{
		URL: admissionScoreLineReq.Url,
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if respAdmission.Status == utils.AdmissionRespStatusErr {
		return echo.NewHTTPError(http.StatusInternalServerError, respAdmission.Message)
	}

	err = h.repo.WithTransaction(c.Request().Context(), func(q *sqlc.Queries) error {
		for _, item := range respAdmission.Data {
			if err = q.CreateAdmissionScoreLine(c.Request().Context(), sqlc.CreateAdmissionScoreLineParams{
				UniversityName:  admissionScoreLineReq.UniversityName,
				Province:        item.Province,
				Year:            item.Year,
				AdmissionBatch:  item.AdmissionBatch,
				AdmissionType:   item.AdmissionType,
				SubjectCategory: item.SubjectCategory,
				MajorGroup:      item.MajorGroup,
				LowestScore:     strings.Split(item.LowestScore, ".")[0],
				LowestScoreRank: item.LowestScoreRank,
			}); err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	admissionScoreResp.Total = len(respAdmission.Data)

	return utils.SuccessWithData(c, admissionScoreResp)
}

func (h *AdmissionScoreLineHandler) ListAdmissionScoreLines(c echo.Context) error {
	var req struct {
		Page            int    `json:"page" form:"page" query:"page"` // 当前页码（从1开始）
		UniversityName  string `json:"university_name" form:"university_name" query:"university_name"`
		AdmissionBatch  string `json:"admission_batch" form:"admission_batch" query:"admission_batch"`
		AdmissionType   string `json:"admission_type" form:"admission_type" query:"admission_type"`
		SubjectCategory string `json:"subject_category" form:"subject_category" query:"subject_category"`
	}

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	items, err := h.repo.ListAdmissionScoreLines(c.Request().Context(), repository.AdmissionScoreLineQueryParams{
		UniversityName:  req.UniversityName,
		AdmissionBatch:  req.AdmissionBatch,
		AdmissionType:   req.AdmissionType,
		SubjectCategory: req.SubjectCategory,
		Limit:           utils.PageSize,
		Offset:          utils.Offset(req.Page),
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	totalCount, err := h.repo.CountAdmissionScoreLines(c.Request().Context(), repository.AdmissionScoreLineQueryParams{
		AdmissionType:   req.AdmissionType,
		UniversityName:  req.UniversityName,
		SubjectCategory: req.SubjectCategory,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	ret := utils.Pagination[dto.AdmissionScoreLineResponse]{
		List:       dto.ToAdmissionScoreLineResponses(items),
		TotalCount: totalCount,
		Page:       req.Page,
		PageSize:   utils.PageSize,
	}

	return c.JSON(http.StatusOK, ret)
}

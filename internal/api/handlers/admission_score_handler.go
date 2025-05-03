package handlers

import (
	"fmt"
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

type AdmissionScoreHandler struct {
	repo *database.Database
	cfg  *config.Config
}

func NewAdmissionScrapeHandler(q *database.Database, cfg *config.Config) *AdmissionScoreHandler {
	return &AdmissionScoreHandler{repo: q, cfg: cfg}
}

func (h *AdmissionScoreHandler) CreateAdmissionScore(c echo.Context) error {
	var admissionScoreReq struct {
		UniversityName string `json:"university_name"`
		Year           string `json:"year"`
	}

	var admissionScoreResp struct {
		Total int `json:"total"`
	}

	if err := c.Bind(&admissionScoreReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	admissionQueryCondition, err := h.repo.GetQueryConditionByYearAndName(c.Request().Context(), sqlc.GetQueryConditionByYearAndNameParams{
		Year:           admissionScoreReq.Year,
		UniversityName: admissionScoreReq.UniversityName,
	})
	if err != nil {
		return utils.Error(c, http.StatusNotFound, fmt.Sprintf("failed to query condition: %s", err.Error()))
	}

	respAdmission, err := utils.FetchAdmissionScoreData(h.cfg.Scraper.URL, utils.AdmissionRequest{
		URL:            admissionQueryCondition.Url,
		Year:           admissionQueryCondition.Year,
		Province:       admissionQueryCondition.Province,
		AdmissionType:  admissionQueryCondition.AdmissionType,
		UniversityName: admissionQueryCondition.UniversityName,
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
				UniversityName:    admissionQueryCondition.UniversityName,
				Province:          item.Province,
				Year:              item.Year,
				AdmissionType:     item.AdmissionType,
				SubjectCategory:   item.AcademicCategory,
				MajorName:         item.MajorName,
				EnrollmentQuota:   item.EnrollmentQuota,
				MinAdmissionScore: strings.Split(item.MinAdmissionScore, ".")[0],
				HighestScore:      strings.Split(item.HighestScore, ".")[0],
				HighestScoreRank:  item.HighestScoreRank,
				LowestScore:       strings.Split(item.LowestScore, ".")[0],
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

	admissionScoreResp.Total = len(respAdmission.Data)

	return utils.SuccessWithData(c, admissionScoreResp)
}

func (h *AdmissionScoreHandler) ListAdmissionScores(c echo.Context) error {
	var req struct {
		Page            int    `json:"page" form:"page" query:"page"` // 当前页码（从1开始）
		UniversityName  string `json:"university_name" form:"university_name" query:"university_name"`
		AdmissionType   string `json:"admission_type" form:"admission_type" query:"admission_type"`
		SubjectCategory string `json:"subject_category" form:"subject_category" query:"subject_category"`
	}

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	items, err := h.repo.ListAdmissionScores(c.Request().Context(), repository.AdmissionScoreQueryParams{
		UniversityName:  req.UniversityName,
		AdmissionType:   req.AdmissionType,
		SubjectCategory: req.SubjectCategory,
		Limit:           utils.PageSize,
		Offset:          utils.Offset(req.Page),
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	totalCount, err := h.repo.CountAdmissionScores(c.Request().Context(), repository.AdmissionScoreQueryParams{
		AdmissionType:   req.AdmissionType,
		UniversityName:  req.UniversityName,
		SubjectCategory: req.SubjectCategory,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	ret := utils.Pagination[dto.AdmissionScoreResponse]{
		List:       dto.ToAdmissionScoreResponses(items),
		TotalCount: totalCount,
		Page:       req.Page,
		PageSize:   utils.PageSize,
	}

	return c.JSON(http.StatusOK, ret)
}

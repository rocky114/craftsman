package api

import (
	"github.com/labstack/echo/v4"
	"github.com/rocky114/craftman/internal/api/handlers"
	"github.com/rocky114/craftman/internal/app/config"
	"github.com/rocky114/craftman/internal/database"
	"net/http"
)

func RegisterRoutes(e *echo.Echo, repo *database.Database, cfg *config.Config) {
	universityHandler := handlers.NewUniversityHandler(repo, cfg)
	admissionScoreHandler := handlers.NewAdmissionScrapeHandler(repo, cfg)
	admissionSummaryHandler := handlers.NewAdmissionSummaryHandler(repo, cfg)
	scoreDistributionHandler := handlers.NewScoreDistributionHandler(repo, cfg)

	// check
	{
		e.GET("/api/health", func(c echo.Context) error {
			return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
		})
	}

	// 大学路由
	universityGroup := e.Group("/api/universities")
	{
		universityGroup.POST("", universityHandler.CreateUniversity)
		universityGroup.GET("", universityHandler.ListUniversities)
		universityGroup.GET("/:id", universityHandler.GetUniversity)
		universityGroup.DELETE("/:id", universityHandler.DeleteUniversity)
	}

	// 招生分数路由
	admissionScoreGroup := e.Group("/api/admission_scores")
	{
		admissionScoreGroup.POST("", admissionScoreHandler.CreateAdmissionScore)
		admissionScoreGroup.GET("", admissionScoreHandler.ListAdmissionScores)
	}

	// 招生汇总路由
	admissionSummaryGroup := e.Group("/api/admission_summaries")
	{
		admissionSummaryGroup.GET("", admissionSummaryHandler.ListAdmissionSummaries)
	}

	// 分数段分布路由
	scoreDistributionGroup := e.Group("/api/score_distributions")
	{
		scoreDistributionGroup.GET("", scoreDistributionHandler.ListScoreDistributions)
	}
}

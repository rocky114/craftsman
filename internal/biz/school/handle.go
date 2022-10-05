package school

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rocky114/craftsman/internal/response"
	"github.com/rocky114/craftsman/internal/service/school"
	"github.com/rocky114/craftsman/internal/storage"
	"github.com/sirupsen/logrus"
)

func ListSchool(c *gin.Context) {
	var req storage.ListSchoolParams
	var err error

	if err := c.ShouldBindJSON(&req); err != nil {
		logrus.Errorf("ListSchool err: %v", err)
		c.JSON(http.StatusBadRequest, response.NewFail(response.ErrInvalidParam))
		return
	}

	schools, err := school.ListSchool(req)
	if err != nil {
		logrus.Errorf("ListSchool err: %v", err)
		c.JSON(http.StatusOK, response.NewFail(response.ErrUnknown))
		return
	}

	c.JSON(http.StatusOK, response.NewSuccess(schools))
}

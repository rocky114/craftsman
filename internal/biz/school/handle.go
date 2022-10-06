package school

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rocky114/craftsman/internal/pkg/common"
	"github.com/rocky114/craftsman/internal/response"
	"github.com/rocky114/craftsman/internal/service/school"
	"github.com/rocky114/craftsman/internal/storage"
	"github.com/sirupsen/logrus"
)

type ListSchoolRequest struct {
	Page int32 `json:"page"`
	Size int32 `json:"size"`
}

func ListSchool(c *gin.Context) {
	var req ListSchoolRequest
	var err error

	if err := c.ShouldBindJSON(&req); err != nil {
		logrus.Errorf("ListSchool err: %v", err)
		c.JSON(http.StatusBadRequest, response.NewFail(response.ErrInvalidParam))
		return
	}

	limit, offset := common.GetLimitAndOffset(req.Page, req.Size)
	res, err := school.ListSchool(storage.ListSchoolsParams{Limit: limit, Offset: offset})
	if err != nil {
		logrus.Errorf("ListSchool err: %v", err)
		c.JSON(http.StatusOK, response.NewFail(response.ErrUnknown))
		return
	}

	c.JSON(http.StatusOK, response.NewSuccess(res))
}

package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type SuccessResponse struct {
	Status string      `json:"status"` // success/error
	Data   interface{} `json:"data"`   // 实际数据
}

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

var (
	EmptyArray  = []struct{}{}
	OKStatus    = "OK"
	ErrorStatus = "ERROR"
)

func Success(c echo.Context) error {
	return c.JSON(http.StatusOK, SuccessResponse{
		Status: OKStatus,
		Data:   EmptyArray,
	})
}

func SuccessWithData(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, SuccessResponse{
		Status: OKStatus,
		Data:   data,
	})
}

func Error(c echo.Context, code int, message string) error {
	return c.JSON(code, ErrorResponse{
		Status:  ErrorStatus,
		Message: message,
	})
}

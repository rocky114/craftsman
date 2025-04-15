package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Status  string      `json:"status"`  // success/error
	Data    interface{} `json:"data"`    // 实际数据
	Message string      `json:"message"` // 提示信息
}

var (
	EmptyArray = []struct{}{}
)

func Success(c echo.Context) error {
	return c.JSON(http.StatusOK, Response{
		Status:  "success",
		Data:    EmptyArray,
		Message: "",
	})
}

func SuccessWithData(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, Response{
		Status:  "success",
		Data:    data,
		Message: "",
	})
}

func SuccessWithMessage(c echo.Context, data interface{}, message string) error {
	return c.JSON(http.StatusOK, Response{
		Status:  "success",
		Data:    data,
		Message: message,
	})
}

func Error(c echo.Context, code int, message string) error {
	return c.JSON(code, Response{
		Status:  "error",
		Data:    EmptyArray,
		Message: message,
	})
}

func ErrorWithData(c echo.Context, code int, data interface{}, message string) error {
	return c.JSON(code, Response{
		Status:  "error",
		Data:    data,
		Message: message,
	})
}

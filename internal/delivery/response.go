package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   *ErrorInfo  `json:"error,omitempty"`
	Meta    *Meta       `json:"meta,omitempty"`
}

type ErrorInfo struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type Meta struct {
	Page       int `json:"page,omitempty"`
	PerPage    int `json:"per_page,omitempty"`
	Total      int `json:"total,omitempty"`
	TotalPages int `json:"total_pages,omitempty"`
}

func OK(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Data:    data,
	})
}

func Fail(ctx *gin.Context, status int, code, message string) {
	ctx.JSON(status, Response{
		Success: false,
		Error:   &ErrorInfo{Code: code, Message: message},
	})
}

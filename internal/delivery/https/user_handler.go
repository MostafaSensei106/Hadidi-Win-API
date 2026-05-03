package https

import (
	"github.com/MostafaSensei106/Hadidi-Win-API/internal/domain"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserUsecase domain.UserUsecase
}

func NewUserHandler(r *gin.Engine, us domain.UserUsecase) {
	handler := &UserHandler{
		UserUsecase: us,
	}

	v1 := r.Group("/v1")
	{
	}
}

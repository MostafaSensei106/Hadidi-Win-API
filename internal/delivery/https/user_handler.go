package https

import (
	"net/http"

	"github.com/MostafaSensei106/Hadidi-Win-API/internal/delivery"
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
		v1.GET("/users", handler.GetUsers)
		v1.GET("/users/:id", handler.GetUser)
		v1.POST("/users", handler.CreateUser)
		v1.PUT("/users/:id", handler.UpdateUser)
		v1.DELETE("/users/:id", handler.DeleteUser)
		v1.GET("/users/search", handler.SearchUsers)
	}

}

func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.UserUsecase.GetUsers(c)
	if err != nil {
		delivery.Fail(c, http.StatusNoContent, "Users Not Found", err.Error())
		return
	}
	c.JSON(http.StatusOK, delivery.Response{
		Success: true,
		Data:    users,
	})
}

func (h *UserHandler) GetUser(c *gin.Context) {}

func (h *UserHandler) CreateUser(c *gin.Context) {}

func (h *UserHandler) UpdateUser(c *gin.Context) {}

func (h *UserHandler) DeleteUser(c *gin.Context) {}

func (h *UserHandler) SearchUsers(c *gin.Context) {}

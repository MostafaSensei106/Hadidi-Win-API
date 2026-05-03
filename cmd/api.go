package cmd

import (
	"net/http"

	"github.com/MostafaSensei106/Hadidi-Win-API/internal/constants/api"
	"github.com/gin-gonic/gin"
)

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
	City      string `json:"city"`
	Job       string `json:"job"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Avatar    string `json:"avatar"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

func (a *application) mount() error {
	r := gin.Default()

	{
		v1 := r.Group("v1")

		v1.GET(api.Health, func(ctx *gin.Context) {
			ctx.String(
				http.StatusOK, "PONG FORM V1",
			)
		})
	}

	{
		v2 := r.Group("v2")

		v2.GET(api.User+"/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")

			if id == "0" {
				Fail(ctx, http.StatusNotFound, "User Not Found", "no user with that ID")
				return

			}
			OK(ctx, User{ID: 1, Name: "Alice", Email: "alice@example.com"})

		})
	}
	{
		v3 := r.Group("v3")

		v3.GET(api.Health, func(ctx *gin.Context) {
			ctx.String(
				http.StatusOK, "PONG FORM V3",
			)

		})
	}

	return r.Run(
		a.config.port,
	)

}

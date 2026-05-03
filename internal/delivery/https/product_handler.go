package https

import (
	"net/http"

	"github.com/MostafaSensei106/Hadidi-Win-API/internal/delivery"
	"github.com/MostafaSensei106/Hadidi-Win-API/internal/domain"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	ProductUsecase domain.ProductUsecase
}

func NewProductHandler(r *gin.Engine, us domain.ProductUsecase) {
	handler := &ProductHandler{
		ProductUsecase: us,
	}

	v1 := r.Group("/v1")
	{
		v1.GET("/products", handler.GetProducts)
		v1.GET("/products/:id", handler.GetProduct)
		v1.POST("/products", handler.CreateProduct)
		v1.PUT("/products/:id", handler.UpdateProduct)
		v1.DELETE("/products/:id", handler.DeleteProduct)
		v1.GET("/products/search", handler.SearchProducts)
	}
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	produts, err := h.ProductUsecase.GetProducts(c)
	if err != nil {
		delivery.Fail(c, http.StatusNotFound, "Products Not Found", err.Error())
		return
	}
	c.JSON(
		http.StatusOK, delivery.Response{
			Success: true,
			Data:    produts,
		},
	)
}

func (h *ProductHandler) GetProduct(c *gin.Context) {
	id := c.Param("id")
	product, err := h.ProductUsecase.GetProduct(c.Request.Context(), id)
	if err != nil {
		delivery.Fail(c, http.StatusNotFound, "Product Not Found", err.Error())
		return
	}
	c.JSON(

		http.StatusOK, delivery.Response{
			Success: true,
			Data:    product,
		},
	)

}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	body := domain.Product{}
	product, err := h.ProductUsecase.CreateProduct(c, &body)
	if err != nil {
		delivery.Fail(c, http.StatusBadRequest, "Product Not Created", err.Error())
		return
	}
	c.JSON(http.StatusCreated, delivery.Response{
		Success: true,
		Data:    product,
	})

}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	delivery.Fail(c, http.StatusBadRequest, "bad_request", "bad request")

}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	delivery.Fail(c, http.StatusBadRequest, "bad_request", "bad request")

}

func (h *ProductHandler) SearchProducts(c *gin.Context) {
	delivery.Fail(c, http.StatusBadRequest, "bad_request", "bad request")

}

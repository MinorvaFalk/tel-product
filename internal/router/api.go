package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (r *Router) MapApiHandler() {
	v1 := r.e.Group("/api/v1")

	v1.GET("/status", func(c echo.Context) error { return c.String(http.StatusOK, "OK") })

	product := v1.Group("/products")
	product.GET("/", r.h.PaginateListProduct)
	product.GET("/:id", r.h.GetProduct)
	product.POST("/", r.h.CreateProduct)
	product.PUT("/", r.h.UpdateProduct)
	product.PATCH("/", r.h.PatchProduct)
	product.DELETE("/:id", r.h.DeleteProduct)
}

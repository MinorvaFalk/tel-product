package api

import "github.com/labstack/echo/v4"

type Handler interface {
	PaginateListProduct(c echo.Context) error
	GetProduct(c echo.Context) error
	CreateProduct(c echo.Context) error
	UpdateProduct(c echo.Context) error
	PatchProduct(c echo.Context) error
	DeleteProduct(c echo.Context) error
}

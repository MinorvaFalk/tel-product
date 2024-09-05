package apihandler

import (
	"net/http"
	"tel/product/internal/api"
	"tel/product/internal/model"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	uc api.Usecase
}

func New(uc api.Usecase) *Handler {
	return &Handler{uc: uc}
}

func (h *Handler) PaginateListProduct(c echo.Context) error {
	req := new(model.Pagination)
	if err := c.Bind(req); err != nil {
		return err
	}

	res, err := h.uc.PaginateListProduct(c.Request().Context(), req)
	if err != nil {
		return nil
	}

	return c.JSON(http.StatusOK, model.HTTPResponse{
		Status:     http.StatusOK,
		Message:    "SUCCESS",
		Data:       res,
		Pagination: req,
	})
}

func (h *Handler) GetProduct(c echo.Context) error {
	req := new(model.ProductIDRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}

	res, err := h.uc.GetProduct(c.Request().Context(), req.ID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.HTTPResponse{
		Status:  http.StatusOK,
		Message: "SUCCESS",
		Data:    res,
	})
}

func (h *Handler) CreateProduct(c echo.Context) error {
	req := new(model.ProductsCreateRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}

	res, err := h.uc.CreateProduct(c.Request().Context(), req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.HTTPResponse{
		Status:  http.StatusOK,
		Message: "SUCCESS",
		Data:    res,
	})
}

func (h *Handler) UpdateProduct(c echo.Context) error {
	req := new(model.ProductsUpdateRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}

	res, err := h.uc.UpdateProduct(c.Request().Context(), req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.HTTPResponse{
		Status:  http.StatusOK,
		Message: "SUCCESS",
		Data:    res,
	})
}

func (h *Handler) PatchProduct(c echo.Context) error {
	req := new(model.ProductsPatchRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}

	res, err := h.uc.PatchProduct(c.Request().Context(), req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.HTTPResponse{
		Status:  http.StatusOK,
		Message: "SUCCESS",
		Data:    res,
	})
}

func (h *Handler) DeleteProduct(c echo.Context) error {
	req := new(model.ProductIDRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}

	res, err := h.uc.DeleteProduct(c.Request().Context(), req.ID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.HTTPResponse{
		Status:  http.StatusOK,
		Message: "SUCCESS",
		Data:    res,
	})
}

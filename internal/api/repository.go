package api

import (
	"context"
	"tel/product/internal/entity"
	"tel/product/internal/model"
)

type Repository interface {
	PaginateListProducts(ctx context.Context, paginate *model.Pagination) ([]*entity.Products, error)
	GetProduct(ctx context.Context, id string) (*entity.Products, error)
	CreateProduct(ctx context.Context, data *entity.Products) error
	UpdateProduct(ctx context.Context, data *entity.Products) error
	PatchProduct(ctx context.Context, id string, req map[string]any) (*entity.Products, error)
	DeleteProduct(ctx context.Context, id string) (*entity.Products, error)
}

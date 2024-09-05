package api

import (
	"context"
	"tel/product/internal/model"
)

type Usecase interface {
	PaginateListProduct(ctx context.Context, paginate *model.Pagination) ([]*model.Products, error)
	GetProduct(ctx context.Context, id string) (*model.Products, error)
	CreateProduct(ctx context.Context, req *model.ProductsCreateRequest) (*model.Products, error)
	UpdateProduct(ctx context.Context, req *model.ProductsUpdateRequest) (*model.Products, error)
	PatchProduct(ctx context.Context, req *model.ProductsPatchRequest) (*model.Products, error)
	DeleteProduct(ctx context.Context, id string) (*model.Products, error)
}

package apiusecase

import (
	"context"
	"tel/product/internal/api"
	"tel/product/internal/entity"
	"tel/product/internal/model"
	"tel/product/pkg/exception"

	"github.com/go-viper/mapstructure/v2"
	"github.com/google/uuid"
)

type Usecase struct {
	repo api.Repository
}

func New(repo api.Repository) *Usecase {
	return &Usecase{repo: repo}
}

func (u *Usecase) PaginateListProduct(ctx context.Context, paginate *model.Pagination) ([]*model.Products, error) {
	var data []*model.Products

	res, err := u.repo.PaginateListProducts(ctx, paginate)
	if err != nil {
		return nil, err
	}

	for _, r := range res {
		data = append(data, &model.Products{
			ID:          r.ID.String(),
			Name:        r.Name,
			Description: r.Description,
			Price:       r.Price,
			Variety:     r.Variety,
			Rating:      r.Rating,
			Stock:       r.Stock,
		})
	}

	return data, nil
}

func (u *Usecase) GetProduct(ctx context.Context, id string) (*model.Products, error) {
	res, err := u.repo.GetProduct(ctx, id)
	if err != nil {
		return nil, err
	}

	data := model.Products{
		ID:          res.ID.String(),
		Name:        res.Name,
		Description: res.Description,
		Price:       res.Price,
		Variety:     res.Variety,
		Rating:      res.Rating,
		Stock:       res.Stock,
	}

	return &data, nil
}

func (u *Usecase) CreateProduct(ctx context.Context, req *model.ProductsCreateRequest) (*model.Products, error) {
	data := &entity.Products{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Variety:     req.Variety,
		Stock:       req.Stock,
	}

	if err := u.repo.CreateProduct(ctx, data); err != nil {
		return nil, err
	}

	res := model.Products{
		ID:          data.ID.String(),
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Rating:      data.Rating,
		Variety:     req.Variety,
		Stock:       req.Stock,
	}

	return &res, nil
}

func (u *Usecase) UpdateProduct(ctx context.Context, req *model.ProductsUpdateRequest) (*model.Products, error) {
	data := &entity.Products{
		ID:          uuid.MustParse(req.ID),
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Rating:      req.Rating,
		Variety:     req.Variety,
		Stock:       req.Stock,
	}

	if err := u.repo.UpdateProduct(ctx, data); err != nil {
		return nil, err
	}

	res := model.Products{
		ID:          req.ID,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Rating:      req.Rating,
		Variety:     req.Variety,
		Stock:       req.Stock,
	}

	return &res, nil
}

func (u *Usecase) PatchProduct(ctx context.Context, req *model.ProductsPatchRequest) (*model.Products, error) {
	var data map[string]any
	if err := mapstructure.Decode(req, &data); err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, exception.NewValidatonError("no updates provided", nil)
	}

	res, err := u.repo.PatchProduct(ctx, req.ID, data)
	if err != nil {
		return nil, err
	}

	products := model.Products{
		ID:          res.ID.String(),
		Name:        res.Name,
		Description: res.Description,
		Price:       res.Price,
		Variety:     res.Variety,
		Rating:      res.Rating,
		Stock:       res.Stock,
	}

	return &products, nil
}

func (u *Usecase) DeleteProduct(ctx context.Context, id string) (*model.Products, error) {
	res, err := u.repo.DeleteProduct(ctx, id)
	if err != nil {
		return nil, err
	}

	data := model.Products{
		ID:          res.ID.String(),
		Name:        res.Name,
		Description: res.Description,
		Price:       res.Price,
		Variety:     res.Variety,
		Rating:      res.Rating,
		Stock:       res.Stock,
	}

	return &data, nil
}

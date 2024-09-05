package apirepository

import (
	"context"
	"tel/product/internal/api"
	"tel/product/internal/entity"
	"tel/product/internal/model"
	"tel/product/pkg/exception"
	"tel/product/pkg/utils"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) withTrx(tx *gorm.DB) *Repository {
	return &Repository{
		db: tx,
	}
}

func (r *Repository) Trx(ctx context.Context, fn func(repo api.Repository) error) error {
	tx := r.db.Begin()
	if err := tx.Error; err != nil {
		return exception.NewDBQueryError(err)
	}

	repo := r.withTrx(tx)
	err := fn(repo)
	if err != nil {
		tx.Rollback()
		return exception.NewDBQueryError(err)
	}

	return tx.Commit().Error
}

func (r *Repository) PaginateListProducts(ctx context.Context, paginate *model.Pagination) ([]*entity.Products, error) {
	var data []*entity.Products

	res := r.db.Scopes(utils.Paginate(&entity.Products{}, paginate, r.db)).Find(&data)
	if err := res.Error; err != nil {
		return nil, exception.NewDBQueryError(err)
	}

	return data, nil
}

func (r *Repository) GetProduct(ctx context.Context, id string) (*entity.Products, error) {
	var data entity.Products
	res := r.db.Where("id = ?", id).First(&data)
	if err := res.Error; err != nil {
		return nil, exception.NewDBQueryError(err)
	}

	return &data, nil
}

func (r *Repository) CreateProduct(ctx context.Context, data *entity.Products) error {
	res := r.db.Create(data)
	if err := res.Error; err != nil {
		return exception.NewDBQueryError(err)
	}

	if res.RowsAffected == 0 {
		return exception.NewDBQueryError(nil, "failed to insert data")
	}

	return nil
}

func (r *Repository) UpdateProduct(ctx context.Context, data *entity.Products) error {
	res := r.db.Save(data)
	if err := res.Error; err != nil {
		return exception.NewDBQueryError(err)
	}

	if res.RowsAffected == 0 {
		return exception.NewDBQueryError(nil, "failed to update data")
	}

	return nil
}

func (r *Repository) PatchProduct(ctx context.Context, id string, req map[string]any) (*entity.Products, error) {
	var data entity.Products

	res := r.db.Model(&data).Where("id = ?", id).Updates(req)
	if err := res.Error; err != nil {
		return nil, exception.NewDBQueryError(err)
	}

	return &data, nil
}

func (r *Repository) DeleteProduct(ctx context.Context, id string) (*entity.Products, error) {
	var data entity.Products

	res := r.db.Clauses(clause.Returning{}).Where("id = ?", id).Delete(&data)
	if err := res.Error; err != nil {
		return nil, exception.NewDBQueryError(err)
	}

	return &data, nil
}

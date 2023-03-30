package usecase

import (
	"shop365-products-api/internal/entity"
	"shop365-products-api/internal/usecase/repo"
)

type CategoryUC struct {
	cr *repo.CategoryRepo
}

func NewCategoryUC(cr *repo.CategoryRepo) *CategoryUC {
	return &CategoryUC{
		cr: cr,
	}
}

func (cu *CategoryUC) GetAll() (*[]entity.Category, error) {
	categories, err := cu.cr.GetAll()

	if err != nil {
		return nil, err
	}

	return categories, nil
}

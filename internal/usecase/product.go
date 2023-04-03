package usecase

import (
	"shop365-products-api/internal/entity"
	"shop365-products-api/internal/usecase/repo"
)

type ProductUC struct {
	pr *repo.ProductRepo
}

func NewProductUC(pr *repo.ProductRepo) *ProductUC {
	return &ProductUC{
		pr: pr,
	}
}

func (uc *ProductUC) GetByID(productID, shardID int64) (entity.Product, error) {
	return uc.pr.GetByID(productID, shardID)
}

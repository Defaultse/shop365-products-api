package usecase

import (
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

func (uc *ProductUC) GetAllProducts() {
	uc.pr.GetAll()
}

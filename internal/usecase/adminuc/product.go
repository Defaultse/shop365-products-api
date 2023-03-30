package adminuc

import (
	"shop365-products-api/internal/dto/admindto"
	"shop365-products-api/internal/usecase/repo/adminrepo"
)

type AdminProductUC struct {
	pr *adminrepo.AdminProductRepo
}

func NewProductUC(pr *adminrepo.AdminProductRepo) *AdminProductUC {
	return &AdminProductUC{
		pr: pr,
	}
}

func (uc *AdminProductUC) CreateNew(product *admindto.Product) error {
	return uc.pr.Create(product)
}

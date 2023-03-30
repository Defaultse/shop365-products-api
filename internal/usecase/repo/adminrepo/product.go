package adminrepo

import (
	"fmt"
	"shop365-products-api/internal/dto/admindto"
	"shop365-products-api/internal/entity"
	"shop365-products-api/internal/entity/adminentity"

	"gorm.io/gorm"
)

type AdminProductRepo struct {
	postgres *gorm.DB
}

func NewAdminProductRepo(postgres *gorm.DB) *AdminProductRepo {
	return &AdminProductRepo{
		postgres: postgres,
	}
}

func (pr *AdminProductRepo) Create(p *admindto.Product) error {
	product := &adminentity.AdminProduct{
		Product: entity.Product{
			Name:        p.Name,
			Description: p.Description,
			CategoryID:  p.CategoryID,
			Price:       p.Price,
		},
	}

	if err := pr.postgres.Create(product).Error; err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

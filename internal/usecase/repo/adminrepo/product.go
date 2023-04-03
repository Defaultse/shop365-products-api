package adminrepo

import (
	"fmt"
	"shop365-products-api/internal/dto/admindto"
	"shop365-products-api/internal/entity"
	"shop365-products-api/internal/entity/adminentity"
	"shop365-products-api/pkg/postgres"
)

type AdminProductRepo struct {
	postgres postgres.Postgres
}

func NewAdminProductRepo(postgres *postgres.Postgres) *AdminProductRepo {
	return &AdminProductRepo{
		postgres: *postgres,
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

	if err := pr.postgres.ShardMap[postgres.ShardNum(postgres.GetShardIDFromHash(product.Name))].Create(product).Error; err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

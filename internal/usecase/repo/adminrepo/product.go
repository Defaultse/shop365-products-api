package adminrepo

import (
	"fmt"
	"hash/fnv"
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

	selectedShard := hash(product.Name)%postgres.ShardQuantity + 1

	fmt.Println(selectedShard)

	if err := pr.postgres.ShardMap[postgres.ShardNum(selectedShard)].Create(product).Error; err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

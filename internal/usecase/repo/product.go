package repo

import (
	"shop365-products-api/internal/entity"
	"shop365-products-api/pkg/postgres"
)

type ProductRepo struct {
	postgres *postgres.Postgres
}

func NewProductRepo(postgres *postgres.Postgres) *ProductRepo {
	return &ProductRepo{
		postgres: postgres,
	}
}

func (pr *ProductRepo) GetByID(productID, shardID int64) (entity.Product, error) {
	product := entity.Product{}

	err := pr.postgres.ShardMap[postgres.ShardNum(shardID)].First(&product, productID).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

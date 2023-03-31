package repo

import (
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

func (pr *ProductRepo) GetAll() {

}

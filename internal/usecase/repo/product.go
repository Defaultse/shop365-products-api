package repo

import "gorm.io/gorm"

type ProductRepo struct {
	postgres *gorm.DB
}

func NewProductRepo(postgres *gorm.DB) *ProductRepo {
	return &ProductRepo{
		postgres: postgres,
	}
}

func (pr *ProductRepo) GetAll() {

}

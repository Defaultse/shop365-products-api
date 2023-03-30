package entity

type Product struct {
	ID          int     `gorm:"column:id;primaryKey"`
	Name        string  `gorm:"column:name"`
	Description string  `gorm:"column:description"`
	CategoryID  string  `gorm:"column:category_id"`
	Rating      float32 `gorm:"column:rating"`
	Price       float64 `gorm:"column:price"`
}

func (Product) TableName() string {
	return "products"
}

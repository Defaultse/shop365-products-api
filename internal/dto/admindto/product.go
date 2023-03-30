package admindto

type Product struct {
	Name        string  `form:"name" validate:"required,min=2,max=100"`
	Description string  `form:"description" validate:"required,min=2,max=100"`
	CategoryID  string  `form:"categoryID" validate:"required"`
	Price       float64 `form:"price"`
}

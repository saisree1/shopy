package domain

type Product struct {
	Id       string  `json:"id" binding:"required,startswith=p"`
	Name     string  `json:"name"  binding:"required"`
	Price    float64 `json:"price" binding:"required"`
	Category string  `json:"category" binding:"required,oneof=Regular Budget Premium"`
	Quantity int     `json:"quantity" binding:"required"`
}

package models

type Product struct {
	Id          int
	ProductName string
	Price       float64
	Description string
	CategoryId  int
}

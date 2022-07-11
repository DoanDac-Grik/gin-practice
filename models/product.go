package models

type Product struct {
	Id       string   `json:"id" bson:"id"`
	Name     string   `json:"name" bson:"name" binding:"required"`
	Price    float64  `json:"price" bson:"price"`
	Discount float64  `json:"discount" bson:"discount" binding:"required,discountvalidator"`
	Images   []string `json:"images" bson:"images"`
}

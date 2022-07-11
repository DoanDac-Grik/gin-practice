package validators

import (
	"gin-practice/models"

	"github.com/go-playground/validator/v10"
)

func DiscountValidator(fl validator.FieldLevel) bool {
	var discount = fl.Field().Float()
	var price = fl.Parent().Interface().(models.Product).Price
	if discount > price/2 {
		return false
	}
	return true
}

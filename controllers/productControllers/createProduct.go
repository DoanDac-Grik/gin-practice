package productcontrollers

import (
	"context"
	"gin-practice/models"
	"gin-practice/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateProduct(ctx *gin.Context) {
	var product models.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	product.Id = uuid.New().String()
	var db = services.ConnectMongodb()
	coll := db.Database("gin-practice").Collection("products")
	_, err := coll.InsertOne(context.TODO(), product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}
	ctx.JSON(200, product)
}

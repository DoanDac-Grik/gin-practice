package productcontrollers

import (
	"context"
	"gin-practice/models"
	"gin-practice/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetProduct(ctx *gin.Context) {
	var productId = ctx.Param("id")
	var product models.Product
	var db = services.ConnectMongodb()
	coll := db.Database("gin-practice").Collection("products")
	if err := coll.FindOne(context.TODO(), bson.M{"id": productId}).Decode(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(200, product)
}

func GetAllProduct(ctx *gin.Context) {
	var products []models.Product

	var db = services.ConnectMongodb()
	coll := db.Database("gin-practice").Collection("products")

	cursor, err := coll.Find(context.TODO(), bson.D{{}})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	for cursor.Next(context.TODO()) {
		var product models.Product
		err := cursor.Decode(&product)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		products = append(products, product)
	}

	ctx.JSON(200, products)

	cursor.Close(ctx)
}

func SearchProduct(ctx *gin.Context) {
	var keyword = ctx.DefaultQuery("name", "")
	var products []models.Product

	var db = services.ConnectMongodb()
	coll := db.Database("gin-practice").Collection("products")
	filter := bson.M{"name": bson.M{"$regex": keyword, "$options": "ig"}}

	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := cursor.All(ctx, &products); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(200, products)
}

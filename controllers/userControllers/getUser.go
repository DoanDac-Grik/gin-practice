package usercontrollers

import (
	"context"
	"gin-practice/models"
	"gin-practice/services"

	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUserById(ctx *gin.Context) {
	var user models.User
	var userId = ctx.Param("id")

	var db = services.ConnectMongodb()
	coll := db.Database("gin-practice").Collection("users")
	if err := coll.FindOne(context.TODO(), bson.M{"id": userId}).Decode(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(200, user)
}

func GetAllUsers(ctx *gin.Context) {
	var users []models.User

	var db = services.ConnectMongodb()
	coll := db.Database("gin-practice").Collection("users")

	cursor, err := coll.Find(context.TODO(), bson.D{{}})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	for cursor.Next(context.TODO()) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		users = append(users, user)
	}
	ctx.JSON(200, users)
	cursor.Close(ctx)
}

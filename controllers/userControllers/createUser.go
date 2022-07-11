package usercontrollers

import (
	"context"

	"gin-practice/models"
	"gin-practice/services"
	"gin-practice/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadGateway, err.Error())
		return
	}

	user.Id = uuid.New().String()

	//Hash password before save to MongoDB
	rawPwd := user.Password
	hashPwd, _ := utils.HashPassword(rawPwd)
	user.Password = hashPwd

	var db = services.ConnectMongodb()
	coll := db.Database("gin-practice").Collection("users")

	_, err := coll.InsertOne(context.TODO(), user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}
	ctx.JSON(200, user)
}

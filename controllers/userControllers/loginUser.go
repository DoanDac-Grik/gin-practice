package usercontrollers

import (
	"context"
	"fmt"
	"gin-practice/models"
	"gin-practice/services"
	"gin-practice/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func LoginUser(ctx *gin.Context) {
	var reqUser models.User
	var dbUser models.User

	if err := ctx.ShouldBindJSON(&reqUser); err != nil {
		log.Fatal(err)
	}

	var db = services.ConnectMongodb()
	coll := db.Database("gin-practice").Collection("users")

	if err := coll.FindOne(context.TODO(), bson.M{"username": reqUser.Username}).Decode(&dbUser); err != nil {
		ctx.JSON(http.StatusBadRequest, "User does not existed!!!")
		fmt.Print(dbUser)
		return
	}

	//Check password correctly
	if utils.CheckPasswordHash(dbUser.Password, reqUser.Password) {
		ctx.JSON(http.StatusBadRequest, "Wrong Password!!!")
		return
	}

	var token = services.GenerateToken(dbUser)

	ctx.JSON(200, gin.H{
		"authToken": token,
	})

}

package main

import (
	productcontrollers "gin-practice/controllers/productControllers"
	usercontrollers "gin-practice/controllers/userControllers"
	"gin-practice/middlewares"
	"gin-practice/services"
	"gin-practice/validators"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func main() {
	router := gin.Default()
	port := ":3000"

	//ping pong
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.String(200, "pong")
	})

	//Conect to mongdb
	var db = services.ConnectMongodb()
	//Disconect to db when close
	defer func() {
		services.DisconnectMongodb(db)
	}()

	//Validator
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("discountvalidator", validators.DiscountValidator)
	}

	//Product
	var productRouter = router.Group("/product")

	productRouter.Use(middlewares.DummyMiddleware(), middlewares.Auth())

	productRouter.POST("/", productcontrollers.CreateProduct)
	productRouter.GET("/all", productcontrollers.GetAllProduct)
	productRouter.GET("/:id", productcontrollers.GetProduct)
	productRouter.GET("/search", productcontrollers.SearchProduct)

	//User
	var userRouter = router.Group("user")
	userRouter.POST("/", usercontrollers.CreateUser)
	userRouter.GET("/:username", usercontrollers.GetUserById)
	userRouter.GET("/all", usercontrollers.GetAllUsers)
	userRouter.POST("/login", usercontrollers.LoginUser)

	router.Run(port)
}

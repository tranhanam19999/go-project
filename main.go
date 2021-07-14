package main

import (
	controller "echo-project/controller"
	helper "echo-project/helper"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var clientGloball *mongo.Client

func main() {
	e := echo.New()
	//e.Use;(middleware.Logger())
	helper.ConnectToMongo()
	e.Use(middleware.Recover())
	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	clientGloball = helper.GetMongoClient()
	e.GET("/", controller.Home)
	e.GET("/user/all", controller.GetUsers)
	e.POST("/user/login", controller.LogIn)
	e.POST("/user", controller.CreateUser)
	e.PUT("/user/:id", controller.UpdateUser)
	e.DELETE("/user", controller.DeleteUser)
	// e.POST("/users", service.CreateUser)
	// e.GET("/users/:id", service.GetUser)
	// e.DELETE("/users/:id", service.DeleteUser)

	// Start server at localhost:1323
	//e.Logger.Fatal()
	e.Start(":80")
}

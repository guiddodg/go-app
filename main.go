package main

import (
	"github.com/gin-gonic/gin"
	"github.com/guiddodg/go-jwt/http/controller"
	"github.com/guiddodg/go-jwt/http/middleware"
	"github.com/guiddodg/go-jwt/inicializer"
)

func init() {
	inicializer.LoadEnv()
	inicializer.ConnectToDB()
	inicializer.FixtureLoad()
}
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/signup", controller.SignUp)
	r.POST("/login", controller.Login)
	r.GET("/protected", middleware.RequireAuth, controller.Protected)

	if err := r.Run(); err != nil {
		panic(err)
	}
}

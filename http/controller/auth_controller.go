package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/guiddodg/go-jwt/http/request"
	"github.com/guiddodg/go-jwt/internal/application"
	"net/http"
)

func SignUp(c *gin.Context) {
	var body request.AuthRequest
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read the body request",
		})
		return
	}

	authService := application.NewAuthService()
	if err := authService.Register(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user created successfully",
	})
}

func Login(c *gin.Context) {
	var body request.AuthRequest
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read the body request",
		})
		return
	}
	authService := application.NewAuthService()
	token, err := authService.Login(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	//for local development use secure: false
	c.SetCookie("Authorization", token, 3600, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func Protected(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

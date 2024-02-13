package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/guiddodg/go-jwt/inicializer"
	"github.com/guiddodg/go-jwt/internal/domain/model"
	"net/http"
	"os"
	"time"
)

func RequireAuth(c *gin.Context) {
	tokenString, tokenErr := c.Cookie("Authorization")
	if tokenErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		c.Abort()
		return
	}
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if time.Now().Unix() > int64(claims["exp"].(float64)) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		var user model.User
		inicializer.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		user.Password = "*******"

		c.Set("user", user)
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}

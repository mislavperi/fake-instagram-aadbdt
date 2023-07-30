package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/domain/models"
)

func JwtTokenCheck(c *gin.Context) {
	var claims models.Claims
	token, err := jwt.ParseWithClaims(c.GetHeader("Authorization"), &claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return "asgfasgas", nil
	})
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	if _, ok := token.Claims.(models.Claims); ok && token.Valid {
		c.Next()
	}
	c.AbortWithStatus(http.StatusUnauthorized)
}

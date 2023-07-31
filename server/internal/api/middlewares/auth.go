package middlewares

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/mislavperi/fake-instagram-aadbdt/server/internal/domain/models"
)

func JwtTokenCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var accessClaim models.Claims
		var refreshClaim models.Claims

		clientAccessToken := strings.Split(ctx.GetHeader("Authorization"), " ")[1]

		accessToken, err := jwt.ParseWithClaims(clientAccessToken, &accessClaim, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte("f83edb0a3b4e9547fd6fbd981513bce0d604472c547daaeed8907a78c5793671"), nil
		})
		if err != nil || !accessToken.Valid {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		accessClaims, ok := accessToken.Claims.(*models.Claims)
		if !ok {
			fmt.Println(ok)
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if time.Until(time.Unix(accessClaims.ExpiresAt, 0)) > 15*time.Second {
			ctx.Next()
		}

		refreshAccessToken := strings.Split(ctx.GetHeader("Refresh"), " ")[1]

		refreshToken, err := jwt.ParseWithClaims(refreshAccessToken, &refreshClaim, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte("f83edb0a3b4e9547fd6fbd981513bce0d604472c547daaeed8907a78c5793671"), nil
		})
		if err != nil || !refreshToken.Valid {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		refreshClaims, ok := refreshToken.Claims.(*models.Claims)
		if !ok {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if time.Until(time.Unix(refreshClaims.ExpiresAt, 0)) > 15*time.Second {
			accessExpirationTime := time.Now().Add(5 * time.Minute).Unix()
			accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, models.Claims{
				Identifier: accessClaims.Identifier,
				Type:       "access",
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: accessExpirationTime,
				},
			}).SignedString([]byte("f83edb0a3b4e9547fd6fbd981513bce0d604472c547daaeed8907a78c5793671"))
			if err != nil {
				ctx.AbortWithStatus(http.StatusInternalServerError)
			}

			ctx.SetCookie("accessToken", accessToken, 3600, "/", "localhost", false, true)

			refreshExpirationTime := time.Now().Add(5 * time.Minute).Unix()
			refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, models.Claims{
				Identifier: refreshClaims.Identifier,
				Type:       "refresh",
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: refreshExpirationTime,
				},
			}).SignedString([]byte("f83edb0a3b4e9547fd6fbd981513bce0d604472c547daaeed8907a78c5793671"))
			if err != nil {
				ctx.AbortWithStatus(http.StatusInternalServerError)
			}
			ctx.SetCookie("refreshToken", refreshToken, 172800, "/", "localhost", false, true)
			ctx.Next()
		}

		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

}

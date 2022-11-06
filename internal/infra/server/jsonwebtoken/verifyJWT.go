package jsonwebtoken

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	domain "github.com/badcode256/example_go_mongodb_jwt_hexagonal/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func VerifyJWT() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var secretKey = []byte(os.Getenv("SECRET_JWT"))
		claims := &domain.MyCustomClaims{}
		authorization := ctx.Request.Header.Get("Authorization")
		tokenArr := strings.Split(authorization, "Bearer")

		if len(tokenArr) != 2 {

			ctx.JSON(http.StatusForbidden, &domain.Response{Message: "Invalid token"})
			ctx.Abort()
			return
		}

		token := strings.TrimSpace(tokenArr[1])

		result, _ := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})
		fmt.Printf("%+v\n", claims)
		if !result.Valid {
			ctx.JSON(http.StatusForbidden, &domain.Response{Message: "Invalid token"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

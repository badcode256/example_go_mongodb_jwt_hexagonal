package user

import (
	"net/http"

	domain "github.com/badcode256/example_go_mongodb_jwt_hexagonal/internal/domain"

	"github.com/badcode256/example_go_mongodb_jwt_hexagonal/internal/infra/server/jsonwebtoken"
	"github.com/badcode256/example_go_mongodb_jwt_hexagonal/internal/service"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func FindHandler(userService service.UserService) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var req domain.Login

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, &domain.Response{Message: "bad_request"})
			return
		}

		user, exist := userService.FindUser(req.Email)
		if !exist {
			ctx.JSON(http.StatusInternalServerError, &domain.Response{Message: "invalid email"})
			return
		}

		passwordHashReq := []byte(req.Password)
		passwordHashDb := []byte(user.Password)

		err := bcrypt.CompareHashAndPassword(passwordHashDb, passwordHashReq)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &domain.Response{Message: "invalid credentials"})
			return
		}

		tokenString, err := jsonwebtoken.GeneratedToken(user)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &domain.Response{Message: err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{"token": tokenString})
	}
}

package user

import (
	"net/http"

	domain "github.com/badcode256/example_go_mongodb_jwt_hexagonal/internal/domain"
	"github.com/badcode256/example_go_mongodb_jwt_hexagonal/internal/service"
	"github.com/gin-gonic/gin"
)

func DeleteHandler(userService service.UserService) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		err := userService.DeleteUser(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &domain.Response{Message: err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
	}
}

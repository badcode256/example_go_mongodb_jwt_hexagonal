package jsonwebtoken

import (
	"fmt"
	"os"
	"time"

	domain "github.com/badcode256/example_go_mongodb_jwt_hexagonal/internal/domain"
	"github.com/golang-jwt/jwt/v4"
)

func GeneratedToken(userResponse domain.UsersResponse) (string, error) {
	var secretKey = []byte(os.Getenv("SECRET_JWT"))
	fmt.Printf("%+v\n", "pass secret")
	payload := jwt.MapClaims{
		"email": userResponse.Email,
		"_id":   userResponse.Id.Hex(),
		"exp":   time.Now().Add(time.Minute * 5).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	fmt.Printf("%+v\n", "pass token")
	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return tokenString, err
	}
	return tokenString, nil

}

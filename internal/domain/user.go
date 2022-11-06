package domain

import (
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Response represents the api response
type Response struct {
	Message string `json:"message"`
}

// InsertUser input  user
type IUser struct {
	User_name string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"createdAt"`
}

// UpdateUser input  user
type UUser struct {
	Id        string `json:"id"`
	User_name string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

//List User
type Users struct {
	Id        primitive.ObjectID `bson:"_id" json:"_id"`
	User_name string             `json:"username"`
	Email     string             `json:"email"`
	CreatedAt string             `json:"createdAt"`
	UpdatedAt string             `json:"updatedAt"`
}

//List UserResponse
type UsersResponse struct {
	Id        primitive.ObjectID `bson:"_id" json:"_id"`
	User_name string             `json:"username"`
	Password  string             `json:"password"`
	Email     string             `json:"email"`
	CreatedAt string             `json:"createdAt"`
	UpdatedAt string             `json:"updatedAt"`
}

//Login request
type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type MyCustomClaims struct {
	Email string             `json:"email"`
	ID    primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	jwt.StandardClaims
}

// User repository implementations

type UserRepository interface {
	CreateUser(user IUser) error
	UpdateUser(user UUser) error
	DeleteUser(id string) error
	FindUser(username string) (userResponse UsersResponse, exist bool)
	ListUsers() (*[]Users, error)
}

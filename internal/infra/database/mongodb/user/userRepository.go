package user

import (
	"context"
	"fmt"
	"time"

	domain "github.com/badcode256/example_go_mongodb_jwt_hexagonal/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	db *mongo.Client
}

func NewUserRepository(db *mongo.Client) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(user domain.IUser) error {

	bytes, _ := bcrypt.GenerateFromPassword([]byte(string(user.Password)), 8)

	user.Password = string(bytes)
	currentTime := time.Now()
	dateNow := fmt.Sprintf("%d-%d-%d %d:%d:%d", currentTime.Year(), currentTime.Month(), currentTime.Day(), currentTime.Hour(), currentTime.Hour(), currentTime.Second())
	user.CreatedAt = dateNow
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := r.db.Database("dbBusiness")
	collection := database.Collection("users")

	_, err := collection.InsertOne(ctx, user)

	if err != nil {
		return fmt.Errorf("error create document user: %v", err)
	}

	return nil

}

func (r *UserRepository) UpdateUser(user domain.UUser) error {
	currentTime := time.Now()
	dateNow := fmt.Sprintf("%d-%d-%d %d:%d:%d", currentTime.Year(), currentTime.Month(), currentTime.Day(), currentTime.Hour(), currentTime.Hour(), currentTime.Second())

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := r.db.Database("dbBusiness")
	collection := database.Collection("users")
	objID, _ := primitive.ObjectIDFromHex(user.Id)
	filter := bson.D{{"_id", objID}}
	update := bson.D{{"$set", bson.D{{"username", user.User_name}, {"email", user.Email}, {"password", user.Password}, {"updatedAt", dateNow}}}}

	_, err := collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return fmt.Errorf("error update document user: %v", err)
	}

	return nil

}

func (r *UserRepository) DeleteUser(id string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := r.db.Database("dbBusiness")
	collection := database.Collection("users")
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objID}
	_, err := collection.DeleteOne(ctx, filter)

	if err != nil {
		return fmt.Errorf("error delete document user: %v", err)
	}

	return nil

}
func (r *UserRepository) FindUser(email string) (userResponse domain.UsersResponse, exist bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := r.db.Database("dbBusiness")
	collection := database.Collection("users")
	var result domain.UsersResponse
	filter := bson.M{"email": email}
	err := collection.FindOne(ctx, filter).Decode(&result)

	if err != nil {
		return result, false
	}

	return result, true

}
func (r *UserRepository) ListUsers() (*[]domain.Users, error) {
	var users []domain.Users
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := r.db.Database("dbBusiness")
	collection := database.Collection("users")

	filter := bson.M{}
	rows, err := collection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}
	for rows.Next(context.TODO()) {
		var user domain.Users
		err2 := rows.Decode(&user)
		if err2 != nil {
			return nil, err2
		}
		users = append(users, user)
	}

	return &users, nil
}

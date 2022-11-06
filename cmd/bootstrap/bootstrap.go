package bootstrap

import (
	"context"
	"os"

	"log"

	"github.com/badcode256/example_go_mongodb_jwt_hexagonal/internal/infra/database/mongoDb/user"
	"github.com/badcode256/example_go_mongodb_jwt_hexagonal/internal/infra/server"
	"github.com/badcode256/example_go_mongodb_jwt_hexagonal/internal/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var uriDb = os.Getenv("MONGO_URI")

func Start() error {

	db, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uriDb))

	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	err = db.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	userRepository := user.NewUserRepository(db)
	userService := service.NewUserService(userRepository)

	server := server.New(context.Background(), "localhost", 3000, userService)

	return server.Run()
}

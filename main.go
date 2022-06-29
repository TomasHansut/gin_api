package main

import (
	"context"
	"fmt"
	"log"

	"github.com/TomasHansut/gin_api/controllers"
	"github.com/TomasHansut/gin_api/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server         *gin.Engine
	userservice    services.UserService
	usercontroller controllers.UserController
	ctx            context.Context
	usercollection *mongo.Collection
	mongoclient    *mongo.Client
	err            error
)

func init() {
	// Create empty context
	ctx = context.TODO()
	// Create mongo connection
	mongoconn := options.Client().ApplyURI("mongodb://localhost:27017")
	// Mongo client connected to DB and check for erros
	mongoclient, err = mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatalln(err)
	}
	// Ping DB to test connection and check for errors
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("mongo connection was established")
	// Get user db
	usercollection = mongoclient.Database("userdb").Collection("users")
	userservice = services.NewUserService(usercollection, ctx)
	usercontroller = controllers.New(userservice)
	server = gin.Default()
}

func main() {
	// Desconnect from DB when program shutdown
	defer mongoclient.Disconnect(ctx)

	basepath := server.Group("/v1")
	usercontroller.RegisterUserRoutes(basepath)

	log.Fatal(server.Run(":9090"))
}

package main

import (
	"fmt"
	"github.com/gabrielsouzacoder/clean-new/api/routes"
	"github.com/gabrielsouzacoder/clean-new/infrastructure/repository"
	"github.com/gabrielsouzacoder/clean-new/usecase/todo"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

func main() {
	fmt.Println("[Server] Initializing ...")
	loadEnvironment()
	todoRepo := selectDatabase()
	todoService := todo.NewService(todoRepo)

	server := NewServer()
	server.Run(todoService)
}

func loadEnvironment() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Println("[Warning] The .env file could not be loaded")
	}
}

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   "8080",
		server: gin.Default(),
	}
}

func (s *Server) Run(todo *todo.Service) {
	router := routes.ConfigRoutes(s.server, todo)

	log.Printf("Server running at port: %v", s.port)
	log.Fatal(router.Run(":" + s.port))
}

func selectDatabase() todo.Repository {
	dbType := os.Getenv("DB_TYPE")

	var todoRepo todo.Repository

	if dbType == "mongo" {
		todoRepo = repository.NewMongoDbRepository(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	} else {
		todoRepo = repository.NewInMemoryDatabase()
	}
	return todoRepo
}

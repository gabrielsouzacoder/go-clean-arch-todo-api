package main

import (
	"fmt"
	"github.com/gabrielsouzacoder/clean-new/api/handler"
	"github.com/gabrielsouzacoder/clean-new/infrastructure/repository"
	"github.com/gabrielsouzacoder/clean-new/usecase/todo"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/urfave/negroni"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	fmt.Println("[Server] Initializing ...")

	loadEnvironment()

	todoRepo := selectDatabase()

	todoService := todo.NewService(todoRepo)

	routers := mux.NewRouter()

	n := negroni.New(
		negroni.NewLogger(),
	)

	handler.MakeTodoHandlers(routers, *n, todoService)

	http.Handle("/", routers)

	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":8080",
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func loadEnvironment() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Println("[Warning] The .env file could not be loaded")
	}
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

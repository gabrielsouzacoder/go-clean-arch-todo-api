package handler

import (
	"encoding/json"
	"github.com/gabrielsouzacoder/clean-new/api/presenter"
	"github.com/gabrielsouzacoder/clean-new/usecase/todo"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"log"
	"net/http"
)

func listTodos(service todo.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, _ := service.ListTodos()

		w.Header().Set("Content-Type", "application/json")

		if data == nil {
			err := json.NewEncoder(w).Encode(make([]string, 0))
			if err != nil {
				return
			}
			return
		}

		var toJ []*presenter.Todo
		for _, d := range data {
			toJ = append(toJ, &presenter.Todo{
				ID:       d.ID,
				Description: d.Description,
				Status: d.Status,
			})
		}

		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, err := w.Write([]byte("error to presenter data"))
			if err != nil {
				return
			}
		}
	})
}

func createTodos(service todo.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			Description    string `json:"description"`
		}

		err := json.NewDecoder(r.Body).Decode(&input)

		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, err := w.Write([]byte("error to decode json"))
			if err != nil {
				return
			}
			return
		}

		_, err2 := service.CreateTodo(input.Description)

		if err2 != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, err2 := w.Write([]byte("Error to create an todo"))
			if err2 != nil {
				return
			}
			return
		}

		w.WriteHeader(http.StatusCreated)
		return
	})
}

func MakeTodoHandlers(r *mux.Router, n negroni.Negroni, service todo.UseCase) {
	r.Handle("/v1/todo", n.With(
		negroni.Wrap(listTodos(service)),
	)).Methods("GET", "OPTIONS").Name("listTodos")

	r.Handle("/v1/todo", n.With(
		negroni.Wrap(createTodos(service)),
	)).Methods("POST", "OPTIONS").Name("createTodos")
}
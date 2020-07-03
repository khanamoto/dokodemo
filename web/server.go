package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/khanamoto/dokodemo/service"
)

type Server interface {
	Handler() http.Handler
}

func NewServer(app service.Dokodemo) Server {
	return &server{app: app} // ServeHTTP(ResponseWriter, *Request) の形で返る
}

type server struct {
	app service.Dokodemo
}

func (s server) Handler() http.Handler {
	allowedOrigins := handlers.AllowedOrigins([]string{"http://localhost:3000"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With"})

	router := mux.NewRouter()

	router.HandleFunc("/hello", s.willSignupHandler).Methods("GET")
	router.HandleFunc("/todos", s.todoIndex).Methods("GET")

	return handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(router)
}

func (s *server) willSignupHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World\n")
}


type Todo struct {
	Name string `json:"name"`
	Completed bool `json:"completed"`
}

type Todos []Todo

func (s *server) todoIndex(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Name: "https://golang.org"},
		Todo{Name: "http://go.shibu.jp/effective_go.html"},
	}

	json.NewEncoder(w).Encode(todos)
}
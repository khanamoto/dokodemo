package web

import (
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/khanamoto/dokodemo/service"
)

type Server interface {
	Handler() http.Handler
}

const sessionKey = "DOKODEMO_SESSION"

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

	router.HandleFunc("/signup", s.signupHandler).Methods("POST")

	return handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(router)
}

func (s *server) signupHandler(w http.ResponseWriter, r *http.Request) {
	name, password, userName, email := r.FormValue("name"), r.FormValue("password"), r.FormValue("userName"), r.FormValue("email")
	if err := s.app.CreateNewUser(name, userName, email, password); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user, err := s.app.FindUserByName(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	expiresAt := time.Now().Add(24 * time.Hour)
	token, err := s.app.CreateNewToken(user.ID, expiresAt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    sessionKey,
		Value:   token,
		Expires: expiresAt,
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

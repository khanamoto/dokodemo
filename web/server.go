package web

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/khanamoto/dokodemo/model"
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
	router.HandleFunc("/signin", s.signinHandler).Methods("POST")
	router.HandleFunc("/signout", s.signoutHandler).Methods("POST")
	router.HandleFunc("/group", s.addStudyGroupHandler).Methods("POST")
	// router.HandleFunc("/subgroup", s.createSubStudyGroupHandler).Methods("POST")

	return handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(router)
}

func (s *server) signupHandler(w http.ResponseWriter, r *http.Request) {
	userDataSet := &model.User{
		Name:     r.FormValue("name"),
		UserName: r.FormValue("userName"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}
	if err := validateBaseUser(userDataSet); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	name, userName, email, password := userDataSet.Name, userDataSet.UserName, userDataSet.Email, userDataSet.Password
	if err := s.app.CreateNewUser(name, userName, email, password); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user, err := s.app.FindUserByUserName(userName)
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

func (s *server) signinHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: ユーザーがログイン済みか findUser でチェックする

	userDataset := &model.User{
		UserName: r.FormValue("userName"),
		Password: r.FormValue("password"),
	}
	if err := validateUserNameAndPassword(userDataset); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userName, password := userDataset.UserName, userDataset.Password
	if ok, err := s.app.LoginUser(userName, password); err != nil || !ok {
		http.Error(w, "user not found or invalid password", http.StatusBadRequest)
		return
	}

	user, err := s.app.FindUserByUserName(userName)
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

	fmt.Println("login success")

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (s *server) signoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    sessionKey,
		Value:   "",
		Expires: time.Unix(0, 0),
	})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (s *server) addStudyGroupHandler(w http.ResponseWriter, r *http.Request) {
	studyGroupDataSet := &model.StudyGroup{
		Name: r.FormValue("name"),
		URL:  r.FormValue("url"),
	}
	if err := validateAll(studyGroupDataSet); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userDataSet := &model.User{
		UserName: r.FormValue("userName"),
	}
	if err := validateUserName(userDataSet); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	name, url, userName := studyGroupDataSet.Name, studyGroupDataSet.URL, userDataSet.UserName
	studyGroup, err := s.app.CreateStudyGroup(name, url)
	if err != nil {
		http.Error(w, "failed to create study group", http.StatusBadRequest)
		return
	}
	if _, err := s.app.CreateMembership(studyGroup.ID, userName); err != nil {
		http.Error(w, "failed to create membership", http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// func (s *server) findStudyGroup(r *http.Request) (studyGroup *model.StudyGroup) {
//
// }

// func (s *server) createSubStudyGroupHandler(w http.ResponseWriter, r *http.Request) {
// 	studyGroup := s.app.FindStudyGroupByID(r.)
// 	if studyGroup == nil {
// 		http.Error(w, "please create study group", http.StatusBadRequest)
// 		return
// 	}
//
// 	name := r.FormValue("name")
// 	if err := s.app.CreateNewStudyGroup(studyGroup.ID, name); err != nil {
// 		http.Error(w, "failed to create sub study group", http.StatusBadRequest)
// 		return
// 	}
//
// 	http.Redirect(w, r, "/", http.StatusSeeOther)
// }

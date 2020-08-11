package web

import (
	"fmt"
	"net/http"
	"strconv"
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

	// router.HandleFunc("/", s.indexHandler).Methods("GET")

	router.HandleFunc("/signup", s.signupHandler).Methods("POST")
	router.HandleFunc("/signin", s.signinHandler).Methods("POST")
	router.HandleFunc("/signout", s.signoutHandler).Methods("POST")

	// router.HandleFunc("/registrations", s.registrationHandler).Methods("POST")

	router.HandleFunc("/organizations", s.addOrganizationHandler).Methods("POST")

	router.HandleFunc("/organizations/{id:[0-9]+}/departments", s.addDepartmentHandler)

	// TODO: departmentsなしでも作れるようにする
	router.HandleFunc("/departments/{id:[0-9]+}/groups", s.addStudyGroupHandler).Methods("POST")

	router.HandleFunc("/events", s.addEventHandler).Methods("POST")
	router.HandleFunc("/groups/{id:[0-9]+}/events", s.addEventHandler).Methods("POST")

	return handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(router)
}

func (s *server) findUser(r *http.Request) (user *model.User) {
	cookie, err := r.Cookie(sessionKey)
	if err == nil && cookie.Value != "" {
		user, _ = s.app.FindUserByToken(cookie.Value)
	}
	return
}

// func (s *server) indexHandler(w http.ResponseWriter, r *http.Request) {
//
// }

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

func (s *server) addOrganizationHandler(w http.ResponseWriter, r *http.Request) {
	organizationDataSet := &model.Organization{
		Name: r.FormValue("name"),
		URL:  r.FormValue("url"),
	}
	if err := validateAll(organizationDataSet); err != nil {
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

	name, url, userName := organizationDataSet.Name, organizationDataSet.URL, userDataSet.UserName
	organization, err := s.app.CreateOrganization(name, url)
	if err != nil {
		http.Error(w, "failed to create organization", http.StatusBadRequest)
		return
	}
	if _, err := s.app.CreateBelonging(organization.ID, userName); err != nil {
		http.Error(w, "failed to create belonging", http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (s *server) addDepartmentHandler(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	organizationID, err := strconv.ParseUint(v["id"], 10, 64)
	if err != nil {
		http.Error(w, "invalid department id", http.StatusBadRequest)
		return
	}
	// TODO: organizationIDが存在するものか確認（ブラウザ以外でpostされたときに弾くために）

	departmentDataSet := &model.Department{
		Name: r.FormValue("name"),
		URL:  r.FormValue("url"),
	}
	if err := validateAll(departmentDataSet); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// TODO: フォーム入力されたuserNameリストをバリデーションして、変数にセットする
	//      userNameはグループに所属していること（管理者でなくてもいい）
	//      userNameは存在していること（service層で確認）

	name, url := departmentDataSet.Name, departmentDataSet.URL
	userNames := r.Form["userName"]

	department, err := s.app.CreateDepartment(organizationID, name, url)
	if err != nil {
		http.Error(w, "failed to create department", http.StatusBadRequest)
		return
	}
	if _, err := s.app.CreateStaff(department.ID, userNames); err != nil {
		http.Error(w, "failed to create staff", http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (s *server) addStudyGroupHandler(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	departmentID, err := strconv.ParseUint(v["id"], 10, 64)
	if err != nil {
		http.Error(w, "invalid department id", http.StatusBadRequest)
		return
	}
	// TODO: departmentIDが存在するものか確認（ブラウザ以外でpostされたときに弾くために）

	studyGroupDataSet := &model.StudyGroup{
		Name: r.FormValue("name"),
		URL:  r.FormValue("url"),
	}
	if err := validateAll(studyGroupDataSet); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// TODO: フォーム入力されたuserNameリストをバリデーションして、変数にセットする
	//      userNameはグループに所属していること（管理者でなくてもいい）
	//      userNameは存在していること（service層で確認）

	name, url := studyGroupDataSet.Name, studyGroupDataSet.URL
	userNames := r.Form["userName"]

	studyGroup, err := s.app.CreateStudyGroup(departmentID, name, url)
	if err != nil {
		http.Error(w, "failed to create study group", http.StatusBadRequest)
		return
	}
	if _, err := s.app.CreateMembership(studyGroup.ID, userNames); err != nil {
		http.Error(w, "failed to create membership", http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (s *server) addEventHandler(w http.ResponseWriter, r *http.Request) {
	user := s.findUser(r)
	if user == nil {
		http.Error(w, "please login", http.StatusBadRequest)
		return
	}

	name, description, place := r.FormValue("name"), r.FormValue("description"), r.FormValue("place")
	eventDate := stringToTime(r.FormValue("eventYear"), r.FormValue("eventMonth"), r.FormValue("eventDay"), r.FormValue("eventHour"), r.FormValue("eventMin"))
	eventDataSet := &model.Event{
		Name:        name,
		EventDate:   eventDate,
		Description: description,
		Place:       place,
	}
	if err := validateAll(eventDataSet); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	event, err := s.app.CreateEvent(name, eventDate, description, place)
	if err != nil {
		http.Error(w, "failed to create event", http.StatusBadRequest)
		return
	}
	if _, err := s.app.CreateAdministrator(user.ID, event.ID); err != nil {
		http.Error(w, "failed to create administrator", http.StatusBadRequest)
		return
	}

	v := mux.Vars(r)
	if v["id"] != "" {
		studyGroupID, err := strconv.ParseUint(v["id"], 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if _, err := s.app.CreateOwnership(studyGroupID, event.ID); err != nil {
			http.Error(w, "failed to create ownership", http.StatusBadRequest)
			return
		}
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

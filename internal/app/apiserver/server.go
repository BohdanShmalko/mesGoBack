package apiserver

import (
	"github.com/BohdanShmalko/mesGoBack/internal/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type server struct {
	logger *logrus.Logger
	router *mux.Router
	store store.Store
}


func newServer(store store.Store) *server {
	s := &server{
		store: store,
		router: mux.NewRouter(),
		logger: logrus.New(),
	}

	s.routerConfigure()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) routerConfigure() {
	s.router.HandleFunc("/users", s.HandleUsersFunc()).Methods("POST")
}

func (s *server) HandleUsersFunc() func (w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

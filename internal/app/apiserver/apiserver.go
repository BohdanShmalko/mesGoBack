package apiserver

import (
	"github.com/BohdanShmalko/mesGoBack/internal/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type ApiServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store *store.Store
}

func New(config *Config) *ApiServer {
	return &ApiServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *ApiServer) Start() error {
	if err := s.loggerConfigure(); err != nil {
		return err
	}
	s.routerConfigure()

	if err := s.storeConfigure(); err != nil {
		return err
	}

	s.logger.Info("the server started on port " + s.config.BindAddr)
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *ApiServer) routerConfigure() {
	s.router.HandleFunc("/someroute", s.testRoute())
}

func (s *ApiServer) loggerConfigure() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)

	return nil
}

func (s *ApiServer) storeConfigure() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st
	return nil
}

func (s *ApiServer) testRoute() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, "<h1>test page</h1>")
	}
}



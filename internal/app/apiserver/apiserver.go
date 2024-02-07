package apiserver

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type API struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

func New(config *Config) *API {
	return &API{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *API) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.configureRouter()
	s.logger.Info("starting API server at", s.config.BindAddr)
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *API) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

func (s *API) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *API) handleHello() http.HandlerFunc {
	// here we can declare some types or variables
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello, it's Onnywrite!")
	}
}

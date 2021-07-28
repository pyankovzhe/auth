package apiserver

import (
	"context"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/pyankovzhe/auth/internal/app/producer"
	"github.com/pyankovzhe/auth/internal/app/store"
	"github.com/sirupsen/logrus"
)

const (
	ctxKeyAccount ctxKey = iota
)

type ctxKey int8

type server struct {
	*http.Server
	logger   *logrus.Logger
	store    store.Store
	producer producer.Producer
}

func newServer(store store.Store, logger *logrus.Logger, serverAddr string, producer producer.Producer) *server {
	s := &server{
		Server: &http.Server{
			Addr: serverAddr,
		},
		logger:   logger,
		store:    store,
		producer: producer,
	}

	r := chi.NewRouter()
	s.configureRouter(r)

	s.Handler = r

	s.logger.Info("Server initialized")

	return s
}

func (s *server) configureRouter(router *chi.Mux) {
	router.Use(middleware.RequestID)
	router.Use(middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: s.logger, NoColor: false}))
	router.Use(middleware.Recoverer)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("JWT Auth Server"))
	})

	router.Post("/signup", s.CreateAccount)
	router.Post("/signin", s.SignIn)
	router.Mount("/api", s.apiRouter())
}

func (s *server) apiRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(s.AuthorizedOnly)
	r.Get("/profile", s.GetProfile)
	return r
}

func (s *server) AuthorizedOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientToken := r.Header.Get("Authorization")
		if clientToken == "" {
			render.Render(w, r, &ErrResponse{Code: http.StatusUnauthorized, Message: "No Authorization header provided"})
			return
		}

		extractedToken := strings.Split(clientToken, "Bearer ")

		if len(extractedToken) == 2 {
			clientToken = strings.TrimSpace(extractedToken[1])
		} else {
			render.Render(w, r, &ErrResponse{Code: http.StatusBadRequest, Message: "Incorrect Format of Authorization Token"})
			return
		}

		claims, err := parseToken(clientToken)
		if err != nil {
			render.Render(w, r, &ErrResponse{Code: http.StatusUnauthorized, Message: err.Error()})
			return
		}

		acc, err := s.store.Account().FindByLogin(claims.Login)
		if err != nil {
			render.Render(w, r, &ErrResponse{Code: http.StatusUnauthorized, Message: errNotAuthenticated.Error()})
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyAccount, acc)))
	})
}

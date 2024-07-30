package api

import (
	"context"
	"net/http"
	"tags/internal/storage"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type APIServer struct {
	listenAddr string
	store      storage.Storage
}

func NewAPIServer(listenAddr string, store storage.Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	router.HandleFunc("/api/tags", makeHTTPHandleFunc(s.handleTags))
	router.HandleFunc("/api/tags/{id}", makeHTTPHandleFunc(s.handleById))

	logrus.Println("JSON API Server running on port:", s.listenAddr)
	return http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) Shutdown(ctx context.Context) error {
	return s.Shutdown(ctx)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *APIServer) handleTags(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return s.handleGetAll(w, r)
	case "POST":
		return s.handleCreate(w, r)
	default:
		return fmt.Errorf("method not allowed %s", r.Method)
	}
}

func (s *APIServer) handleById(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return s.handleGetOne(w, r)
	case "DELETE":
		return s.handleDelete(w, r)
	case "PUT":
		return s.handleUpdate(w, r)
	default:
		return fmt.Errorf("method not allowed %s", r.Method)
	}
}

func WriteJSON(w http.ResponseWriter, status int, value any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(value)
}

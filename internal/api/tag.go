package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tags/internal/entity"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// api/tags

func (s *APIServer) handleCreate(w http.ResponseWriter, r *http.Request) error {
	tagReq := new(entity.Tag)
	if err := json.NewDecoder(r.Body).Decode(tagReq); err != nil {
		return err
	}
	id := uuid.New().String()

	val := entity.NewTag(id, tagReq.Name)
	err := s.store.Create(val)

	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, val)
}

func (s *APIServer) handleGetAll(w http.ResponseWriter, r *http.Request) error {
	tags, err := s.store.GetAll()
	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, tags)
}

// api/tags/{id}

func (s *APIServer) handleGetOne(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	tag, err := s.store.GetOne(id)
	if err != nil {
		return fmt.Errorf("tag id %s not found", id)
	}
	return WriteJSON(w, http.StatusOK, tag)
}

func (s *APIServer) handleDelete(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	if err := s.store.Delete(id); err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, map[string]string{"deleted": id})
}

func (s *APIServer) handleUpdate(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	tagReq := new(entity.Tag)
	if err := json.NewDecoder(r.Body).Decode(tagReq); err != nil {
		return err
	}
	tag := entity.NewTag(id, tagReq.Name)
	err := s.store.Update(tag)

	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, tag)
}

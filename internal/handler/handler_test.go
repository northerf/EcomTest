package handler

import (
	"Ecom/internal/repository"
	"Ecom/internal/schema"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateTodo(t *testing.T) {
	repo := repository.NewRepos()
	h := NewHandler(repo)
	todo := schema.Todo{
		Title:       "test",
		Description: "test123",
	}
	body, _ := json.Marshal(todo)
	req := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewReader(body))
	w := httptest.NewRecorder()
	h.createTodo(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}

func TestCreateTodo_Title(t *testing.T) {
	repo := repository.NewRepos()
	h := NewHandler(repo)
	todo := schema.Todo{Title: ""}
	body, _ := json.Marshal(todo)
	req := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewReader(body))
	w := httptest.NewRecorder()
	h.createTodo(w, req)
	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected 400, got %d", w.Code)
	}
}

func TestGetTodo_NotFound(t *testing.T) {
	repo := repository.NewRepos()
	h := NewHandler(repo)
	req := httptest.NewRequest(http.MethodGet, "/todos/666", nil)
	w := httptest.NewRecorder()
	h.getTodo(w, req)
	if w.Code != http.StatusNotFound {
		t.Errorf("Expected 404, got %d", w.Code)
	}
}

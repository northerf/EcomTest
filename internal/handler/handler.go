package handler

import (
	"Ecom/internal/repository"
	"Ecom/internal/schema"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	repo *repository.MemoryRepos
}

func NewHandler(repo *repository.MemoryRepos) *Handler {
	return &Handler{repo: repo}
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func NewRouter(h *Handler) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/todos", h.HandleTodos)
	mux.HandleFunc("/todos/", h.HandleTodoByID)
	return LogMiddleware(mux)
}

func (h *Handler) HandleTodos(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getTodos(w)
	case http.MethodPost:
		h.createTodo(w, r)
	default:
		respondError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func (h *Handler) HandleTodoByID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getTodo(w, r)
	case http.MethodPut:
		h.updateTodo(w, r)
	case http.MethodDelete:
		h.deleteTodo(w, r)
	default:
		respondError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func (h *Handler) createTodo(w http.ResponseWriter, r *http.Request) {
	var todo schema.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	if todo.ID != 0 {
		respondError(w, http.StatusBadRequest, "ID is not required")
		return
	}
	if err := todo.Validate(); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	created, _ := h.repo.Create(todo)
	respondJSON(w, http.StatusCreated, created)
}

func (h *Handler) getTodos(w http.ResponseWriter) {
	todos := h.repo.GetAll()
	respondJSON(w, http.StatusOK, todos)
}

func (h *Handler) getTodo(w http.ResponseWriter, r *http.Request) {
	id, err := extractID(r.URL.Path)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid ID")
		return
	}
	todo, err := h.repo.GetByID(id)
	if errors.Is(err, repository.ErrNotFound) {
		respondError(w, http.StatusNotFound, "Todo not found")
		return
	}
	respondJSON(w, http.StatusOK, todo)
}

func (h *Handler) updateTodo(w http.ResponseWriter, r *http.Request) {
	id, err := extractID(r.URL.Path)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid ID")
		return
	}
	var todo schema.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	if err := todo.Validate(); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.repo.Update(id, todo); err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			respondError(w, http.StatusNotFound, "Todo not found")
			return
		}
	}
	todo.ID = id
	respondJSON(w, http.StatusOK, todo)
}

func (h *Handler) deleteTodo(w http.ResponseWriter, r *http.Request) {
	id, err := extractID(r.URL.Path)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid ID")
		return
	}
	if err := h.repo.Delete(id); err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			respondError(w, http.StatusNotFound, "Todo not found")
			return
		}
	}
	w.WriteHeader(http.StatusNoContent)
}

func extractID(path string) (int, error) {
	parts := strings.Split(strings.Trim(path, "/"), "/")
	if len(parts) < 2 {
		return 0, errors.New("invalid path")
	}
	return strconv.Atoi(parts[len(parts)-1])
}

func respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func respondError(w http.ResponseWriter, status int, message string) {
	respondJSON(w, status, ErrorResponse{Error: message})
}

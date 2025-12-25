package repository

import (
	"Ecom/internal/schema"
	"testing"
)

func TestMemoryRepos_Create(t *testing.T) {
	repo := NewRepos()
	todo1, err := repo.Create(schema.Todo{Title: "First", Completed: false})
	if err != nil {
		t.Fatalf("Create unexpected error: %v", err)
	}
	if todo1.ID != 1 {
		t.Fatalf("First ID = %d, want 1", todo1.ID)
	}
	todo2, err := repo.Create(schema.Todo{Title: "Second", Completed: true})
	if err != nil {
		t.Fatalf("Create unexpected error: %v", err)
	}
	if todo2.ID != 2 {
		t.Fatalf("Second ID = %d, want 2", todo2.ID)
	}
	if todo2.Title != "Second" || !todo2.Completed {
		t.Fatalf("Second todo fields not set correctly: %+v", todo2)
	}
}

func TestMemoryRepos_GetByID(t *testing.T) {
	repo := NewRepos()
	created, _ := repo.Create(schema.Todo{Title: "Test"})
	tests := []struct {
		name      string
		id        int
		wantErr   error
		wantTitle string
	}{
		{"found", created.ID, nil, "Test"},
		{"not found", 999, ErrNotFound, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.GetByID(tt.id)
			if err != tt.wantErr {
				t.Fatalf("GetByID error = %v, want %v", err, tt.wantErr)
			}
			if err == nil && got.Title != tt.wantTitle {
				t.Fatalf("GetByID Title = %s, want %s", got.Title, tt.wantTitle)
			}
		})
	}
}

func TestMemoryRepos_Update(t *testing.T) {
	repo := NewRepos()
	created, _ := repo.Create(schema.Todo{Title: "Old", Description: "Old"})
	err := repo.Update(created.ID, schema.Todo{
		Title:       "New",
		Description: "New",
		Completed:   true,
	})
	if err != nil {
		t.Fatalf("Update unexpected error: %v", err)
	}
	updated, _ := repo.GetByID(created.ID)
	if updated.Title != "New" || updated.Description != "New" || !updated.Completed {
		t.Fatalf("Update did not update fields correctly: %+v", updated)
	}
	if err := repo.Update(999, schema.Todo{Title: "Nope"}); err != ErrNotFound {
		t.Fatalf("Update for non-existing should return ErrNotFound, got %v", err)
	}
}

func TestMemoryRepos_Delete(t *testing.T) {
	repo := NewRepos()
	created, _ := repo.Create(schema.Todo{Title: "To delete"})
	if err := repo.Delete(created.ID); err != nil {
		t.Fatalf("Delete unexpected error: %v", err)
	}
	if _, err := repo.GetByID(created.ID); err != ErrNotFound {
		t.Fatalf("After delete GetByID must return ErrNotFound, got %v", err)
	}
	if err := repo.Delete(999); err != ErrNotFound {
		t.Fatalf("Delete for non-existing should return ErrNotFound, got %v", err)
	}
}

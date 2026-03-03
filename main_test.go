package main

import (
	"testing"
	"time"
)

func TestTodoCreation(t *testing.T) {
	todos = []Todo{}
	nextID = 1

	title := "Test task"
	priority := High
	dueDate := time.Now().Add(24 * time.Hour)
	tags := []string{"test", "important"}

	addTodoItem(title, priority, &dueDate, tags)

	if len(todos) == 0 {
		t.Fatal("Expected 1 todo, got 0")
	}

	todo := todos[0]

	if todo.ID != 1 {
		t.Errorf("Expected ID 1, got %d", todo.ID)
	}

	if todo.Title != title {
		t.Errorf("Expected title '%s', got '%s'", title, todo.Title)
	}

	if todo.Completed != false {
		t.Errorf("Expected completed false, got %v", todo.Completed)
	}

	if todo.Priority != priority {
		t.Errorf("Expected priority %v, got %v", priority, todo.Priority)
	}

	if todo.DueDate == nil {
		t.Error("Expected due date, got nil")
	}

	if len(todo.Tags) != 2 {
		t.Errorf("Expected 2 tags, got %d", len(todo.Tags))
	}

	if nextID != 2 {
		t.Errorf("Expected nextID 2, got %d", nextID)
	}
}

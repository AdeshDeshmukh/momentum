package main

import (
	"bufio"
	"fmt"
	"html/template"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
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

func TestAddTodoItem(t *testing.T) {
	tests := []struct {
		name     string
		title    string
		priority Priority
		tags     []string
		wantLen  int
	}{
		{
			name:     "add first todo",
			title:    "First task",
			priority: High,
			tags:     []string{"work"},
			wantLen:  1,
		},
		{
			name:     "add second todo",
			title:    "Second task",
			priority: Medium,
			tags:     []string{"personal"},
			wantLen:  2,
		},
		{
			name:     "add with multiple tags",
			title:    "Tagged task",
			priority: Low,
			tags:     []string{"urgent", "important", "work"},
			wantLen:  3,
		},
		{
			name:     "add with no tags",
			title:    "No tags",
			priority: Medium,
			tags:     []string{},
			wantLen:  4,
		},
	}

	todos = []Todo{}
	nextID = 1

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			addTodoItem(tt.title, tt.priority, nil, tt.tags)

			if len(todos) != tt.wantLen {
				t.Errorf("got %d todos, want %d", len(todos), tt.wantLen)
			}

			lastTodo := todos[len(todos)-1]
			if lastTodo.Title != tt.title {
				t.Errorf("got title %s, want %s", lastTodo.Title, tt.title)
			}

			if lastTodo.Priority != tt.priority {
				t.Errorf("got priority %v, want %v", lastTodo.Priority, tt.priority)
			}

			if len(lastTodo.Tags) != len(tt.tags) {
				t.Errorf("got %d tags, want %d", len(lastTodo.Tags), len(tt.tags))
			}
		})
	}
}

func TestToggleComplete(t *testing.T) {
	tests := []struct {
		name          string
		initialState  bool
		toggleID      int
		wantCompleted bool
	}{
		{
			name:          "toggle incomplete to complete",
			initialState:  false,
			toggleID:      1,
			wantCompleted: true,
		},
		{
			name:          "toggle complete to incomplete",
			initialState:  true,
			toggleID:      1,
			wantCompleted: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			todos = []Todo{
				{ID: 1, Title: "Test", Completed: tt.initialState, Priority: Medium},
			}
			nextID = 2

			toggleComplete(tt.toggleID)

			if todos[0].Completed != tt.wantCompleted {
				t.Errorf("got completed %v, want %v", todos[0].Completed, tt.wantCompleted)
			}
		})
	}
}

func TestDeleteTodoItem(t *testing.T) {
	tests := []struct {
		name       string
		deleteID   int
		wantLen    int
		wantTitles []string
	}{
		{
			name:       "delete first item",
			deleteID:   1,
			wantLen:    2,
			wantTitles: []string{"Second", "Third"},
		},
		{
			name:       "delete middle item",
			deleteID:   2,
			wantLen:    2,
			wantTitles: []string{"First", "Third"},
		},
		{
			name:       "delete last item",
			deleteID:   3,
			wantLen:    2,
			wantTitles: []string{"First", "Second"},
		},
		{
			name:       "delete non-existent item",
			deleteID:   999,
			wantLen:    3,
			wantTitles: []string{"First", "Second", "Third"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			todos = []Todo{
				{ID: 1, Title: "First", Priority: Low},
				{ID: 2, Title: "Second", Priority: Medium},
				{ID: 3, Title: "Third", Priority: High},
			}
			nextID = 4

			deleteTodoItem(tt.deleteID)

			if len(todos) != tt.wantLen {
				t.Errorf("got %d todos, want %d", len(todos), tt.wantLen)
			}

			for i, title := range tt.wantTitles {
				if i >= len(todos) {
					t.Errorf("missing todo at index %d", i)
					continue
				}
				if todos[i].Title != title {
					t.Errorf("todo[%d] got title %s, want %s", i, todos[i].Title, title)
				}
			}
		})
	}
}

func TestCalculateStatistics(t *testing.T) {
	todos = []Todo{
		{ID: 1, Title: "Task 1", Completed: true, Priority: High, Tags: []string{"work"}},
		{ID: 2, Title: "Task 2", Completed: false, Priority: Medium, Tags: []string{"personal"}},
		{ID: 3, Title: "Task 3", Completed: false, Priority: Low},
		{ID: 4, Title: "Task 4", Completed: true, Priority: High, Tags: []string{"urgent"}},
	}

	stats := calculateStatistics()

	if stats.Total != 4 {
		t.Errorf("got total %d, want 4", stats.Total)
	}

	if stats.Completed != 2 {
		t.Errorf("got completed %d, want 2", stats.Completed)
	}

	if stats.Pending != 2 {
		t.Errorf("got pending %d, want 2", stats.Pending)
	}

	if stats.High != 2 {
		t.Errorf("got high %d, want 2", stats.High)
	}

	if stats.Medium != 1 {
		t.Errorf("got medium %d, want 1", stats.Medium)
	}

	if stats.Low != 1 {
		t.Errorf("got low %d, want 1", stats.Low)
	}

	if stats.Tagged != 3 {
		t.Errorf("got tagged %d, want 3", stats.Tagged)
	}

	if stats.CompletionRate != 50.0 {
		t.Errorf("got completion rate %.1f%%, want 50.0%%", stats.CompletionRate)
	}
}

func TestSearchTodos(t *testing.T) {
	todos = []Todo{
		{ID: 1, Title: "Buy groceries", Priority: Low},
		{ID: 2, Title: "Fix bug in code", Priority: High},
		{ID: 3, Title: "Review code", Priority: Medium},
		{ID: 4, Title: "Buy birthday gift", Priority: Medium},
	}

	tests := []struct {
		name    string
		keyword string
		wantLen int
	}{
		{"search buy", "buy", 2},
		{"search code", "code", 2},
		{"search bug", "bug", 1},
		{"search nonexistent", "xyz", 0},
		{"case insensitive", "BUY", 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := searchTodos(tt.keyword)
			if len(result) != tt.wantLen {
				t.Errorf("got %d results, want %d", len(result), tt.wantLen)
			}
		})
	}
}

func TestSortByPriority(t *testing.T) {
	todos = []Todo{
		{ID: 1, Title: "Low task", Priority: Low, Completed: false},
		{ID: 2, Title: "High task", Priority: High, Completed: false},
		{ID: 3, Title: "Medium task", Priority: Medium, Completed: false},
	}

	sortByPriority()

	if todos[0].Priority != High {
		t.Errorf("first todo priority: got %v, want High", todos[0].Priority)
	}

	if todos[1].Priority != Medium {
		t.Errorf("second todo priority: got %v, want Medium", todos[1].Priority)
	}

	if todos[2].Priority != Low {
		t.Errorf("third todo priority: got %v, want Low", todos[2].Priority)
	}
}

func TestSortByStatus(t *testing.T) {
	todos = []Todo{
		{ID: 1, Title: "Completed", Completed: true},
		{ID: 2, Title: "Pending", Completed: false},
		{ID: 3, Title: "Also pending", Completed: false},
	}

	sortByStatus()

	if todos[0].Completed != false {
		t.Error("first todo should be pending")
	}

	if todos[1].Completed != false {
		t.Error("second todo should be pending")
	}

	if todos[2].Completed != true {
		t.Error("third todo should be completed")
	}
}

func TestParseTags(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantLen int
	}{
		{"single tag", "work", 1},
		{"multiple tags", "work, personal, urgent", 3},
		{"with spaces", "  work  ,  personal  ", 2},
		{"empty string", "", 0},
		{"only commas", ",,", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseTags(tt.input)
			if len(result) != tt.wantLen {
				t.Errorf("got %d tags, want %d", len(result), tt.wantLen)
			}
		})
	}
}

func TestTodoMethods(t *testing.T) {
	now := time.Now()
	yesterday := now.Add(-24 * time.Hour)
	tomorrow := now.Add(24 * time.Hour)
	nextWeek := now.Add(7 * 24 * time.Hour)

	t.Run("DueDateFormatted", func(t *testing.T) {
		todo := Todo{DueDate: &tomorrow}
		result := todo.DueDateFormatted()
		if result == "" {
			t.Error("expected formatted date, got empty string")
		}
	})

	t.Run("DueDateDisplay", func(t *testing.T) {
		todo := Todo{DueDate: &tomorrow}
		result := todo.DueDateDisplay()
		if result == "" {
			t.Error("expected display date, got empty string")
		}
	})

	t.Run("IsOverdue", func(t *testing.T) {
		tests := []struct {
			name     string
			todo     Todo
			wantTrue bool
		}{
			{"overdue task", Todo{DueDate: &yesterday, Completed: false}, true},
			{"future task", Todo{DueDate: &nextWeek, Completed: false}, false},
			{"completed overdue", Todo{DueDate: &yesterday, Completed: true}, false},
			{"no due date", Todo{DueDate: nil}, false},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := tt.todo.IsOverdue()
				if result != tt.wantTrue {
					t.Errorf("got %v, want %v", result, tt.wantTrue)
				}
			})
		}
	})

	t.Run("IsDueToday", func(t *testing.T) {
		tests := []struct {
			name     string
			todo     Todo
			wantTrue bool
		}{
			{"due tomorrow", Todo{DueDate: &tomorrow, Completed: false}, true},
			{"due next week", Todo{DueDate: &nextWeek, Completed: false}, false},
			{"completed today", Todo{DueDate: &tomorrow, Completed: true}, false},
			{"no due date", Todo{DueDate: nil}, false},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := tt.todo.IsDueToday()
				if result != tt.wantTrue {
					t.Errorf("got %v, want %v", result, tt.wantTrue)
				}
			})
		}
	})

	t.Run("TagsString", func(t *testing.T) {
		todo := Todo{Tags: []string{"work", "urgent"}}
		result := todo.TagsString()
		expected := "work, urgent"
		if result != expected {
			t.Errorf("got %s, want %s", result, expected)
		}
	})
}

func TestPriorityMethods(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		tests := []struct {
			priority Priority
			want     string
		}{
			{Low, "🟢 Low"},
			{Medium, "🟡 Medium"},
			{High, "🔴 High"},
		}

		for _, tt := range tests {
			result := tt.priority.String()
			if result != tt.want {
				t.Errorf("got %s, want %s", result, tt.want)
			}
		}
	})

	t.Run("CSSClass", func(t *testing.T) {
		tests := []struct {
			priority Priority
			want     string
		}{
			{Low, "priority-low"},
			{Medium, "priority-medium"},
			{High, "priority-high"},
		}

		for _, tt := range tests {
			result := tt.priority.CSSClass()
			if result != tt.want {
				t.Errorf("got %s, want %s", result, tt.want)
			}
		}
	})
}

func TestFilterByTag(t *testing.T) {
	todos = []Todo{
		{ID: 1, Title: "Work task", Tags: []string{"work", "urgent"}},
		{ID: 2, Title: "Personal task", Tags: []string{"personal"}},
		{ID: 3, Title: "Another work", Tags: []string{"work"}},
	}

	result := filterByTag("work")

	if len(result) != 2 {
		t.Errorf("got %d results, want 2", len(result))
	}
}

func TestSaveAndLoadTodos(t *testing.T) {
	defer func() {
		os.Remove("todos.json")
	}()

	originalTodos := todos
	originalNextID := nextID

	todos = []Todo{
		{ID: 1, Title: "Save test 1", Priority: High, Completed: false},
		{ID: 2, Title: "Save test 2", Priority: Low, Completed: true},
	}
	nextID = 3

	saveTodos()

	todos = []Todo{}
	nextID = 1

	loadTodos()

	if len(todos) != 2 {
		t.Errorf("got %d todos after load, want 2", len(todos))
	}

	if nextID != 3 {
		t.Errorf("got nextID %d, want 3", nextID)
	}

	todos = originalTodos
	nextID = originalNextID
}

func TestAddHandler(t *testing.T) {
	todos = []Todo{}
	nextID = 1

	form := url.Values{}
	form.Add("title", "New task from handler")
	form.Add("priority", "2")
	form.Add("tags", "work, urgent")

	req := httptest.NewRequest(http.MethodPost, "/add", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	addHandler(w, req)

	if len(todos) != 1 {
		t.Errorf("got %d todos, want 1", len(todos))
	}

	if todos[0].Title != "New task from handler" {
		t.Errorf("got title %s, want 'New task from handler'", todos[0].Title)
	}

	if todos[0].Priority != High {
		t.Errorf("got priority %v, want High", todos[0].Priority)
	}
}

func TestToggleHandler(t *testing.T) {
	todos = []Todo{
		{ID: 1, Title: "Task 1", Completed: false},
	}
	nextID = 2

	form := url.Values{}
	form.Add("id", "1")

	req := httptest.NewRequest(http.MethodPost, "/toggle", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	toggleHandler(w, req)

	if !todos[0].Completed {
		t.Error("expected todo to be completed")
	}
}

func TestDeleteHandler(t *testing.T) {
	todos = []Todo{
		{ID: 1, Title: "Task to delete"},
		{ID: 2, Title: "Task to keep"},
	}
	nextID = 3

	form := url.Values{}
	form.Add("id", "1")

	req := httptest.NewRequest(http.MethodPost, "/delete", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	deleteHandler(w, req)

	if len(todos) != 1 {
		t.Errorf("got %d todos, want 1", len(todos))
	}

	if todos[0].ID != 2 {
		t.Errorf("got ID %d, want 2", todos[0].ID)
	}
}

func TestSortHandler(t *testing.T) {
	todos = []Todo{
		{ID: 1, Title: "Low", Priority: Low},
		{ID: 2, Title: "High", Priority: High},
	}
	nextID = 3

	form := url.Values{}
	form.Add("type", "priority")

	req := httptest.NewRequest(http.MethodPost, "/sort", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	sortHandler(w, req)

	if todos[0].Priority != High {
		t.Errorf("got first priority %v, want High", todos[0].Priority)
	}
}

func TestSearchHandler(t *testing.T) {
	var err error
	tmpl, err = template.ParseGlob("templates/*.html")
	if err != nil {
		t.Skip("Templates not available, skipping test")
	}

	todos = []Todo{
		{ID: 1, Title: "Buy groceries"},
		{ID: 2, Title: "Fix bug"},
	}
	nextID = 3

	req := httptest.NewRequest(http.MethodGet, "/search?q=buy", nil)
	w := httptest.NewRecorder()

	searchHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("got status %d, want %d", w.Code, http.StatusOK)
	}
}

func TestHomeHandler(t *testing.T) {
	var err error
	tmpl, err = template.ParseGlob("templates/*.html")
	if err != nil {
		t.Skip("Templates not available")
	}

	todos = []Todo{{ID: 1, Title: "Test"}}
	nextID = 2

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	homeHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("got status %d, want %d", w.Code, http.StatusOK)
	}
}

func TestStatsHandler(t *testing.T) {
	var err error
	tmpl, err = template.ParseGlob("templates/*.html")
	if err != nil {
		t.Skip("Templates not available")
	}

	todos = []Todo{
		{ID: 1, Title: "Test", Completed: true},
		{ID: 2, Title: "Test 2", Completed: false},
	}
	nextID = 3

	req := httptest.NewRequest(http.MethodGet, "/stats", nil)
	w := httptest.NewRecorder()

	statsHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("got status %d, want %d", w.Code, http.StatusOK)
	}
}

func TestAddTodoItemVariants(t *testing.T) {
	todos = []Todo{}
	nextID = 1

	addTodoItem("Task without tags", Low, nil, []string{})
	addTodoItem("Task with date", Medium, &time.Time{}, []string{"tag"})

	if len(todos) != 2 {
		t.Errorf("got %d todos, want 2", len(todos))
	}
}

func TestToggleCompleteTwice(t *testing.T) {
	todos = []Todo{{ID: 99, Title: "Test", Completed: false}}
	nextID = 100

	toggleComplete(99)
	toggleComplete(99)

	if todos[0].Completed != false {
		t.Error("expected todo to be incomplete after double toggle")
	}
}

func TestDeleteNonExistent(t *testing.T) {
	todos = []Todo{{ID: 1, Title: "Keep me"}}
	nextID = 2

	deleteTodoItem(999)

	if len(todos) != 1 {
		t.Errorf("got %d todos, want 1", len(todos))
	}
}

func TestEmptySearch(t *testing.T) {
	todos = []Todo{{ID: 1, Title: "Task"}}

	result := searchTodos("")

	if len(result) != 1 {
		t.Errorf("got %d results for empty search, want 1", len(result))
	}
}

func TestLoadTodosEdgeCase(t *testing.T) {
	originalTodos := todos
	defer func() {
		todos = originalTodos
	}()

	loadTodos()

	if len(todos) >= 0 {
		t.Log("loadTodos executed successfully")
	}
}

func TestAddHandlerEdgeCases(t *testing.T) {
	t.Run("empty title should not add", func(t *testing.T) {
		todos = []Todo{}
		nextID = 1

		form := url.Values{}
		form.Add("title", "")
		form.Add("priority", "1")

		req := httptest.NewRequest(http.MethodPost, "/add", strings.NewReader(form.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		w := httptest.NewRecorder()

		addHandler(w, req)

		if len(todos) != 0 {
			t.Error("should not add todo with empty title")
		}
	})

	t.Run("whitespace only title should not add", func(t *testing.T) {
		todos = []Todo{}
		nextID = 1

		form := url.Values{}
		form.Add("title", "   ")
		form.Add("priority", "1")

		req := httptest.NewRequest(http.MethodPost, "/add", strings.NewReader(form.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		w := httptest.NewRecorder()

		addHandler(w, req)

		if len(todos) != 0 {
			t.Error("should not add todo with whitespace-only title")
		}
	})

	t.Run("GET request should redirect", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/add", nil)
		w := httptest.NewRecorder()

		addHandler(w, req)

		if w.Code != http.StatusSeeOther {
			t.Errorf("expected redirect, got status %d", w.Code)
		}
	})

	t.Run("with due date and tags", func(t *testing.T) {
		todos = []Todo{}
		nextID = 1

		form := url.Values{}
		form.Add("title", "Task with everything")
		form.Add("priority", "2")
		form.Add("dueDate", "2026-12-31")
		form.Add("tags", "work, urgent, important")

		req := httptest.NewRequest(http.MethodPost, "/add", strings.NewReader(form.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		w := httptest.NewRecorder()

		addHandler(w, req)

		if len(todos) != 1 {
			t.Fatalf("expected 1 todo, got %d", len(todos))
		}

		if todos[0].Priority != High {
			t.Error("expected High priority")
		}

		if todos[0].DueDate == nil {
			t.Error("expected due date to be set")
		}

		if len(todos[0].Tags) != 3 {
			t.Errorf("expected 3 tags, got %d", len(todos[0].Tags))
		}
	})

	t.Run("invalid date format ignored", func(t *testing.T) {
		todos = []Todo{}
		nextID = 1

		form := url.Values{}
		form.Add("title", "Task")
		form.Add("dueDate", "invalid-date")

		req := httptest.NewRequest(http.MethodPost, "/add", strings.NewReader(form.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		w := httptest.NewRecorder()

		addHandler(w, req)

		if len(todos) != 1 {
			t.Fatal("should add todo even with invalid date")
		}

		if todos[0].DueDate != nil {
			t.Error("invalid date should result in nil DueDate")
		}
	})
}

func TestToggleHandlerEdgeCases(t *testing.T) {
	t.Run("GET request should redirect", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/toggle", nil)
		w := httptest.NewRecorder()

		toggleHandler(w, req)

		if w.Code != http.StatusSeeOther {
			t.Errorf("expected redirect, got status %d", w.Code)
		}
	})

	t.Run("invalid ID should not crash", func(t *testing.T) {
		todos = []Todo{{ID: 1, Title: "Test"}}

		form := url.Values{}
		form.Add("id", "invalid")

		req := httptest.NewRequest(http.MethodPost, "/toggle", strings.NewReader(form.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		w := httptest.NewRecorder()

		toggleHandler(w, req)

		if w.Code != http.StatusSeeOther {
			t.Errorf("expected redirect, got status %d", w.Code)
		}
	})

	t.Run("non-existent ID ignored", func(t *testing.T) {
		todos = []Todo{{ID: 1, Title: "Test", Completed: false}}

		form := url.Values{}
		form.Add("id", "999")

		req := httptest.NewRequest(http.MethodPost, "/toggle", strings.NewReader(form.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		w := httptest.NewRecorder()

		toggleHandler(w, req)

		if todos[0].Completed {
			t.Error("todo should remain unchanged")
		}
	})
}

func TestDeleteHandlerEdgeCases(t *testing.T) {
	t.Run("GET request should redirect", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/delete", nil)
		w := httptest.NewRecorder()

		deleteHandler(w, req)

		if w.Code != http.StatusSeeOther {
			t.Errorf("expected redirect, got status %d", w.Code)
		}
	})

	t.Run("invalid ID should not crash", func(t *testing.T) {
		todos = []Todo{{ID: 1, Title: "Test"}}

		form := url.Values{}
		form.Add("id", "abc")

		req := httptest.NewRequest(http.MethodPost, "/delete", strings.NewReader(form.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		w := httptest.NewRecorder()

		deleteHandler(w, req)

		if len(todos) != 1 {
			t.Error("todo should not be deleted with invalid ID")
		}
	})
}

func TestSortHandlerEdgeCases(t *testing.T) {
	t.Run("GET request should redirect", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/sort", nil)
		w := httptest.NewRecorder()

		sortHandler(w, req)

		if w.Code != http.StatusSeeOther {
			t.Error("expected redirect")
		}
	})

	t.Run("sort by status", func(t *testing.T) {
		todos = []Todo{
			{ID: 1, Title: "Done", Completed: true},
			{ID: 2, Title: "Pending", Completed: false},
		}

		form := url.Values{}
		form.Add("type", "status")

		req := httptest.NewRequest(http.MethodPost, "/sort", strings.NewReader(form.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		w := httptest.NewRecorder()

		sortHandler(w, req)

		if todos[0].Completed {
			t.Error("pending todos should come first")
		}
	})

	t.Run("unknown sort type ignored", func(t *testing.T) {
		original := []Todo{
			{ID: 1, Title: "First"},
			{ID: 2, Title: "Second"},
		}
		todos = append([]Todo{}, original...)

		form := url.Values{}
		form.Add("type", "unknown")

		req := httptest.NewRequest(http.MethodPost, "/sort", strings.NewReader(form.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		w := httptest.NewRecorder()

		sortHandler(w, req)

		if todos[0].Title != original[0].Title {
			t.Error("todos should remain unchanged for unknown sort type")
		}
	})
}

func TestSearchHandlerEdgeCases(t *testing.T) {
	var err error
	tmpl, err = template.ParseGlob("templates/*.html")
	if err != nil {
		t.Skip("Templates not available")
	}

	t.Run("empty query redirects", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/search?q=", nil)
		w := httptest.NewRecorder()

		searchHandler(w, req)

		if w.Code != http.StatusSeeOther {
			t.Error("expected redirect for empty query")
		}
	})

	t.Run("whitespace query redirects", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/search?q=+++", nil)
		w := httptest.NewRecorder()

		searchHandler(w, req)

		if w.Code != http.StatusSeeOther {
			t.Error("expected redirect for whitespace query")
		}
	})
}

func TestIntegrationWorkflows(t *testing.T) {
	t.Run("complete workflow: add, toggle, delete", func(t *testing.T) {
		todos = []Todo{}
		nextID = 1

		addTodoItem("Task 1", High, nil, []string{"work"})
		addTodoItem("Task 2", Medium, nil, []string{"personal"})
		addTodoItem("Task 3", Low, nil, []string{})

		if len(todos) != 3 {
			t.Fatalf("expected 3 todos, got %d", len(todos))
		}

		toggleComplete(1)
		toggleComplete(3)

		stats := calculateStatistics()
		if stats.Completed != 2 {
			t.Errorf("expected 2 completed, got %d", stats.Completed)
		}

		deleteTodoItem(2)

		if len(todos) != 2 {
			t.Errorf("expected 2 todos after delete, got %d", len(todos))
		}

		sortByPriority()

		if todos[0].Priority < todos[1].Priority {
			t.Error("todos not sorted by priority correctly")
		}
	})

	t.Run("search and filter workflow", func(t *testing.T) {
		todos = []Todo{
			{ID: 1, Title: "Buy groceries", Tags: []string{"shopping", "personal"}},
			{ID: 2, Title: "Buy gifts", Tags: []string{"shopping"}},
			{ID: 3, Title: "Work meeting", Tags: []string{"work"}},
		}

		results := searchTodos("buy")
		if len(results) != 2 {
			t.Errorf("expected 2 search results, got %d", len(results))
		}

		filtered := filterByTag("shopping")
		if len(filtered) != 2 {
			t.Errorf("expected 2 filtered results, got %d", len(filtered))
		}

		workTasks := filterByTag("work")
		if len(workTasks) != 1 {
			t.Errorf("expected 1 work task, got %d", len(workTasks))
		}
	})

	t.Run("persistence workflow", func(t *testing.T) {
		defer os.Remove("todos.json")

		todos = []Todo{
			{ID: 1, Title: "Persist me", Priority: High, Completed: false},
		}
		nextID = 2

		saveTodos()

		todos = []Todo{}
		nextID = 1

		loadTodos()

		if len(todos) != 1 {
			t.Errorf("expected 1 todo after load, got %d", len(todos))
		}

		if todos[0].Title != "Persist me" {
			t.Error("loaded todo has wrong title")
		}

		if nextID != 2 {
			t.Errorf("expected nextID 2, got %d", nextID)
		}
	})
}

func TestStatisticsEdgeCases(t *testing.T) {
	t.Run("empty todos", func(t *testing.T) {
		todos = []Todo{}

		stats := calculateStatistics()

		if stats.Total != 0 {
			t.Error("expected 0 total")
		}

		if stats.CompletionRate != 0 {
			t.Error("expected 0% completion rate")
		}
	})

	t.Run("all completed", func(t *testing.T) {
		todos = []Todo{
			{ID: 1, Completed: true, Priority: High},
			{ID: 2, Completed: true, Priority: Low},
		}

		stats := calculateStatistics()

		if stats.CompletionRate != 100.0 {
			t.Errorf("expected 100%% completion, got %.1f%%", stats.CompletionRate)
		}
	})

	t.Run("with overdue tasks", func(t *testing.T) {
		yesterday := time.Now().Add(-24 * time.Hour)
		tomorrow := time.Now().Add(24 * time.Hour)

		todos = []Todo{
			{ID: 1, DueDate: &yesterday, Completed: false},
			{ID: 2, DueDate: &tomorrow, Completed: false},
		}

		stats := calculateStatistics()

		if stats.Overdue != 1 {
			t.Errorf("expected 1 overdue, got %d", stats.Overdue)
		}

		if stats.DueToday != 1 {
			t.Errorf("expected 1 due today, got %d", stats.DueToday)
		}
	})
}

func TestPriorityEdgeCases(t *testing.T) {
	t.Run("invalid priority defaults", func(t *testing.T) {
		p := Priority(999)
		result := p.String()
		if result != "Unknown" {
			t.Errorf("expected Unknown, got %s", result)
		}

		cssClass := p.CSSClass()
		if cssClass != "" {
			t.Errorf("expected empty CSS class, got %s", cssClass)
		}
	})
}

func TestTodoMethodEdgeCases(t *testing.T) {
	t.Run("nil due date methods", func(t *testing.T) {
		todo := Todo{DueDate: nil}

		if todo.DueDateFormatted() != "" {
			t.Error("expected empty string for nil due date")
		}

		if todo.DueDateDisplay() != "" {
			t.Error("expected empty string for nil due date")
		}

		if todo.IsOverdue() {
			t.Error("nil due date should not be overdue")
		}

		if todo.IsDueToday() {
			t.Error("nil due date should not be due today")
		}
	})

	t.Run("empty tags", func(t *testing.T) {
		todo := Todo{Tags: []string{}}

		if todo.TagsString() != "" {
			t.Error("expected empty string for empty tags")
		}
	})

	t.Run("single tag", func(t *testing.T) {
		todo := Todo{Tags: []string{"work"}}

		if todo.TagsString() != "work" {
			t.Errorf("expected 'work', got '%s'", todo.TagsString())
		}
	})
}

func TestFileIOErrorHandling(t *testing.T) {
	t.Run("load corrupted JSON", func(t *testing.T) {
		defer os.Remove("todos.json")

		os.WriteFile("todos.json", []byte("{corrupted json"), 0644)

		originalTodos := todos
		todos = []Todo{{ID: 1, Title: "Should be cleared"}}

		loadTodos()

		if len(todos) != 0 {
			t.Error("corrupted JSON should result in empty todos")
		}

		todos = originalTodos
	})

	t.Run("load empty file", func(t *testing.T) {
		defer os.Remove("todos.json")

		os.WriteFile("todos.json", []byte(""), 0644)

		originalTodos := todos
		todos = []Todo{{ID: 1, Title: "Test"}}

		loadTodos()

		if len(todos) != 0 {
			t.Error("empty file should result in empty todos")
		}

		todos = originalTodos
	})

	t.Run("load valid JSON", func(t *testing.T) {
		defer os.Remove("todos.json")

		validJSON := `[{"id":1,"title":"Test","completed":false,"priority":1}]`
		os.WriteFile("todos.json", []byte(validJSON), 0644)

		todos = []Todo{}
		nextID = 1

		loadTodos()

		if len(todos) != 1 {
			t.Errorf("expected 1 todo, got %d", len(todos))
		}

		if todos[0].Title != "Test" {
			t.Errorf("expected 'Test', got '%s'", todos[0].Title)
		}
	})

	t.Run("save creates file", func(t *testing.T) {
		defer os.Remove("todos.json")

		todos = []Todo{
			{ID: 1, Title: "Save test", Priority: Medium, Completed: false},
		}

		saveTodos()

		data, err := os.ReadFile("todos.json")
		if err != nil {
			t.Fatalf("file should be created: %v", err)
		}

		if len(data) == 0 {
			t.Error("file should not be empty")
		}
	})
}

func TestSortByPriorityComplex(t *testing.T) {
	t.Run("priority first, then completion", func(t *testing.T) {
		todos = []Todo{
			{ID: 1, Title: "High done", Priority: High, Completed: true},
			{ID: 2, Title: "Low pending", Priority: Low, Completed: false},
			{ID: 3, Title: "Medium pending", Priority: Medium, Completed: false},
		}

		sortByPriority()

		// Priority sorting is primary: High > Medium > Low
		if todos[0].Priority != High {
			t.Error("High priority should be first regardless of completion")
		}

		if todos[1].Priority != Medium {
			t.Error("Medium priority should be second")
		}

		if todos[2].Priority != Low {
			t.Error("Low priority should be last")
		}
	})

	t.Run("same priority sorted by completion", func(t *testing.T) {
		todos = []Todo{
			{ID: 1, Title: "High done", Priority: High, Completed: true},
			{ID: 2, Title: "High pending", Priority: High, Completed: false},
		}

		sortByPriority()

		if todos[0].Completed {
			t.Error("Pending should come before completed")
		}

		if !todos[1].Completed {
			t.Error("Completed should come after pending")
		}
	})
}

func TestFilterByTagEdgeCases(t *testing.T) {
	t.Run("case insensitive", func(t *testing.T) {
		todos = []Todo{
			{ID: 1, Title: "Task", Tags: []string{"Work"}},
		}

		result := filterByTag("work")

		if len(result) != 1 {
			t.Error("filter should be case insensitive")
		}
	})

	t.Run("no matches", func(t *testing.T) {
		todos = []Todo{
			{ID: 1, Title: "Task", Tags: []string{"work"}},
		}

		result := filterByTag("personal")

		if len(result) != 0 {
			t.Error("should return empty for no matches")
		}
	})

	t.Run("multiple tags", func(t *testing.T) {
		todos = []Todo{
			{ID: 1, Title: "Task1", Tags: []string{"work", "urgent"}},
			{ID: 2, Title: "Task2", Tags: []string{"personal"}},
			{ID: 3, Title: "Task3", Tags: []string{"work"}},
		}

		result := filterByTag("work")

		if len(result) != 2 {
			t.Errorf("expected 2 results, got %d", len(result))
		}
	})
}

func TestSearchTodosEdgeCases(t *testing.T) {
	t.Run("partial match", func(t *testing.T) {
		todos = []Todo{
			{ID: 1, Title: "Complete project"},
			{ID: 2, Title: "Completion report"},
			{ID: 3, Title: "Start new task"},
		}

		result := searchTodos("complet")

		if len(result) != 2 {
			t.Errorf("expected 2 partial matches, got %d", len(result))
		}
	})

	t.Run("special characters", func(t *testing.T) {
		todos = []Todo{
			{ID: 1, Title: "Bug #123"},
			{ID: 2, Title: "Feature request"},
		}

		result := searchTodos("#123")

		if len(result) != 1 {
			t.Error("should find todo with special characters")
		}
	})
}

func TestCompleteUserJourney(t *testing.T) {
	defer os.Remove("todos.json")

	todos = []Todo{}
	nextID = 1

	addTodoItem("Morning routine", High, nil, []string{"personal", "daily"})
	addTodoItem("Team meeting", Medium, nil, []string{"work", "meeting"})
	tomorrow := time.Now().Add(24 * time.Hour)
	addTodoItem("Project deadline", High, &tomorrow, []string{"work", "urgent"})

	if len(todos) != 3 {
		t.Fatalf("expected 3 todos, got %d", len(todos))
	}

	toggleComplete(1)

	stats := calculateStatistics()
	if stats.Completed != 1 {
		t.Errorf("expected 1 completed, got %d", stats.Completed)
	}
	if stats.Pending != 2 {
		t.Errorf("expected 2 pending, got %d", stats.Pending)
	}

	workTasks := filterByTag("work")
	if len(workTasks) != 2 {
		t.Errorf("expected 2 work tasks, got %d", len(workTasks))
	}

	urgentTasks := filterByTag("urgent")
	if len(urgentTasks) != 1 {
		t.Errorf("expected 1 urgent task, got %d", len(urgentTasks))
	}

	saveTodos()

	originalNextID := nextID
	todos = []Todo{}
	nextID = 1

	loadTodos()

	if len(todos) != 3 {
		t.Errorf("expected 3 todos after reload, got %d", len(todos))
	}

	if nextID != originalNextID {
		t.Errorf("expected nextID %d, got %d", originalNextID, nextID)
	}

	searchResults := searchTodos("meeting")
	if len(searchResults) != 1 {
		t.Errorf("expected 1 search result, got %d", len(searchResults))
	}

	sortByPriority()

	if todos[0].Priority != High && todos[1].Priority != High {
		t.Error("high priority tasks should be at top")
	}

	deleteTodoItem(2)

	if len(todos) != 2 {
		t.Errorf("expected 2 todos after delete, got %d", len(todos))
	}

	finalStats := calculateStatistics()
	if finalStats.Total != 2 {
		t.Errorf("expected total 2, got %d", finalStats.Total)
	}
}

func TestHomeHandlerWithFilters(t *testing.T) {
	defer os.Remove("todos.json")

	todos = []Todo{
		{ID: 1, Title: "Task A", Priority: High, Tags: []string{"work"}, Completed: false},
		{ID: 2, Title: "Task B", Priority: Low, Tags: []string{"personal"}, Completed: false},
		{ID: 3, Title: "Task C", Priority: Medium, Tags: []string{"work", "urgent"}, Completed: true},
	}

	if tmpl == nil {
		var err error
		tmpl, err = template.ParseGlob("templates/*.html")
		if err != nil {
			t.Skip("Templates not available")
		}
	}

	t.Run("filter by tag", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/?tag=work", nil)
		w := httptest.NewRecorder()

		homeHandler(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("expected 200, got %d", w.Code)
		}
	})

	t.Run("search query", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/?search=Task", nil)
		w := httptest.NewRecorder()

		homeHandler(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("expected 200, got %d", w.Code)
		}
	})

	t.Run("no filters", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()

		homeHandler(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("expected 200, got %d", w.Code)
		}
	})
}

func TestStatsHandlerComprehensive(t *testing.T) {
	defer os.Remove("todos.json")

	if tmpl == nil {
		var err error
		tmpl, err = template.ParseGlob("templates/*.html")
		if err != nil {
			t.Skip("Templates not available")
		}
	}

	t.Run("with todos", func(t *testing.T) {
		todos = []Todo{
			{ID: 1, Title: "Done task", Completed: true, Priority: High},
			{ID: 2, Title: "Pending task", Completed: false, Priority: Medium},
		}

		req := httptest.NewRequest("GET", "/stats", nil)
		w := httptest.NewRecorder()

		statsHandler(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("expected 200, got %d", w.Code)
		}
	})

	t.Run("empty todos", func(t *testing.T) {
		todos = []Todo{}

		req := httptest.NewRequest("GET", "/stats", nil)
		w := httptest.NewRecorder()

		statsHandler(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("expected 200, got %d", w.Code)
		}
	})
}

func TestAddHandlerWithDateAndTags(t *testing.T) {
	todos = []Todo{}
	nextID = 1

	t.Run("valid due date format", func(t *testing.T) {
		form := url.Values{}
		form.Add("title", "Task with date")
		form.Add("priority", "1")
		form.Add("dueDate", "2026-12-31")
		form.Add("tags", "work,urgent")

		req := httptest.NewRequest("POST", "/add", strings.NewReader(form.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		w := httptest.NewRecorder()

		addHandler(w, req)

		if w.Code != http.StatusSeeOther {
			t.Errorf("expected 303, got %d", w.Code)
		}

		if len(todos) != 1 {
			t.Fatalf("expected 1 todo, got %d", len(todos))
		}

		if todos[0].DueDate == nil {
			t.Error("expected due date to be set")
		}

		if len(todos[0].Tags) != 2 {
			t.Errorf("expected 2 tags, got %d", len(todos[0].Tags))
		}
	})

	t.Run("empty due date", func(t *testing.T) {
		todos = []Todo{}

		form := url.Values{}
		form.Add("title", "Task without date")
		form.Add("priority", "0")
		form.Add("dueDate", "")
		form.Add("tags", "")

		req := httptest.NewRequest("POST", "/add", strings.NewReader(form.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		w := httptest.NewRecorder()

		addHandler(w, req)

		if w.Code != http.StatusSeeOther {
			t.Errorf("expected 303, got %d", w.Code)
		}

		if len(todos) != 1 {
			t.Fatalf("expected 1 todo, got %d", len(todos))
		}

		if todos[0].DueDate != nil {
			t.Error("expected due date to be nil")
		}

		if len(todos[0].Tags) != 0 {
			t.Error("expected no tags")
		}
	})
}

func TestNextIDPersistence(t *testing.T) {
	defer os.Remove("todos.json")

	todos = []Todo{
		{ID: 5, Title: "Task 5"},
		{ID: 10, Title: "Task 10"},
		{ID: 3, Title: "Task 3"},
	}

	saveTodos()

	todos = []Todo{}
	nextID = 1

	loadTodos()

	if nextID != 11 {
		t.Errorf("expected nextID 11 (max ID 10 + 1), got %d", nextID)
	}
}

func TestMultipleSaveLoad(t *testing.T) {
	defer os.Remove("todos.json")

	todos = []Todo{
		{ID: 1, Title: "First batch", Priority: High},
	}
	saveTodos()

	todos = []Todo{
		{ID: 1, Title: "First batch", Priority: High},
		{ID: 2, Title: "Second batch", Priority: Low},
	}
	saveTodos()

	todos = []Todo{}
	loadTodos()

	if len(todos) != 2 {
		t.Errorf("expected 2 todos, got %d", len(todos))
	}

	if todos[1].Title != "Second batch" {
		t.Error("second save did not persist correctly")
	}
}

func TestListTodosCLI(t *testing.T) {
	t.Run("empty todos", func(t *testing.T) {
		todos = []Todo{}

		// Just ensure it doesn't panic
		listTodosCLI()
	})

	t.Run("with todos", func(t *testing.T) {
		todos = []Todo{
			{ID: 1, Title: "Task 1", Priority: High, Completed: false},
			{ID: 2, Title: "Task 2", Priority: Low, Completed: true},
		}

		// Just ensure it doesn't panic and executes
		listTodosCLI()
	})
}

func TestDisplayStatsCLI(t *testing.T) {
	t.Run("normal stats", func(t *testing.T) {
		stats := Statistics{
			Total:          10,
			Completed:      5,
			Pending:        5,
			CompletionRate: 50.0,
			High:           2,
			Medium:         4,
			Low:            4,
			Overdue:        0,
		}

		// Just ensure it doesn't panic
		displayStatsCLI(stats)
	})

	t.Run("with overdue", func(t *testing.T) {
		stats := Statistics{
			Total:          5,
			Completed:      2,
			Pending:        3,
			CompletionRate: 40.0,
			High:           1,
			Medium:         2,
			Low:            2,
			Overdue:        3,
		}

		// Just ensure it doesn't panic and handles overdue path
		displayStatsCLI(stats)
	})
}

func TestDisplayMenu(t *testing.T) {
	// Just ensure it doesn't panic
	displayMenu()
}

func TestGetPriority(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected Priority
	}{
		{"low", "1\n", Low},
		{"high", "3\n", High},
		{"medium explicit", "2\n", Medium},
		{"default empty", "\n", Medium},
		{"default invalid", "5\n", Medium},
		{"default text", "abc\n", Medium},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := bufio.NewScanner(strings.NewReader(tt.input))
			result := getPriority(scanner)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestReadInput(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"simple text", "hello\n", "hello"},
		{"empty", "\n", ""},
		{"with spaces trimmed", "  test  \n", "test"},
		{"number", "123\n", "123"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := bufio.NewScanner(strings.NewReader(tt.input))
			result := readInput(scanner)
			if result != tt.expected {
				t.Errorf("expected '%s', got '%s'", tt.expected, result)
			}
		})
	}
}

func TestSaveTodosEdgeCases(t *testing.T) {
	defer os.Remove("todos.json")

	t.Run("save empty list", func(t *testing.T) {
		todos = []Todo{}
		saveTodos()

		data, err := os.ReadFile("todos.json")
		if err != nil {
			t.Fatalf("should create file: %v", err)
		}

		if len(data) < 2 {
			t.Error("should have at least []")
		}
	})

	t.Run("save large list", func(t *testing.T) {
		todos = []Todo{}
		for i := 1; i <= 100; i++ {
			todos = append(todos, Todo{
				ID:       i,
				Title:    fmt.Sprintf("Task %d", i),
				Priority: Priority(i % 3),
			})
		}

		saveTodos()

		originalTodos := make([]Todo, len(todos))
		copy(originalTodos, todos)

		todos = []Todo{}
		loadTodos()

		if len(todos) != 100 {
			t.Errorf("expected 100 todos, got %d", len(todos))
		}
	})
}

func TestPriorityBoundaries(t *testing.T) {
	tests := []struct {
		priority Priority
		str      string
		css      string
	}{
		{Low, "🟢 Low", "priority-low"},
		{Medium, "🟡 Medium", "priority-medium"},
		{High, "🔴 High", "priority-high"},
		{Priority(99), "Unknown", ""},
	}

	for _, tt := range tests {
		t.Run(tt.str, func(t *testing.T) {
			if tt.priority.String() != tt.str {
				t.Errorf("expected '%s', got '%s'", tt.str, tt.priority.String())
			}
			if tt.priority.CSSClass() != tt.css {
				t.Errorf("expected '%s', got '%s'", tt.css, tt.priority.CSSClass())
			}
		})
	}
}

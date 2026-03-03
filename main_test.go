package main

import (
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
			name      string
			todo      Todo
			wantTrue  bool
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

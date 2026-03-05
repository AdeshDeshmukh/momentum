package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Priority int

const (
	Low Priority = iota
	Medium
	High
)

func (p Priority) String() string {
	switch p {
	case Low:
		return "🟢 Low"
	case Medium:
		return "🟡 Medium"
	case High:
		return "🔴 High"
	default:
		return "Unknown"
	}
}
func (p Priority) CSSClass() string {
	switch p {
	case Low:
		return "priority-low"
	case Medium:
		return "priority-medium"
	case High:
		return "priority-high"
	default:
		return ""
	}
}

type Todo struct {
	ID        int        `json:"id"`
	Title     string     `json:"title"`
	Completed bool       `json:"completed"`
	Priority  Priority   `json:"priority"`
	DueDate   *time.Time `json:"dueDate,omitempty"`
	Tags      []string   `json:"tags"`
}

func (t *Todo) DueDateFormatted() string {
	if t.DueDate == nil {
		return ""
	}
	return t.DueDate.Format("2006-01-02")
}
func (t *Todo) DueDateDisplay() string {
	if t.DueDate == nil {
		return ""
	}
	return t.DueDate.Format("Jan 02, 2006")
}
func (t *Todo) IsOverdue() bool {
	if t.DueDate == nil || t.Completed {
		return false
	}
	return t.DueDate.Before(time.Now())
}
func (t *Todo) IsDueToday() bool {
	if t.DueDate == nil || t.Completed {
		return false
	}
	now := time.Now()
	return t.DueDate.Sub(now).Hours() < 24 && !t.DueDate.Before(now)
}
func (t *Todo) TagsString() string {
	return strings.Join(t.Tags, ", ")
}

type PageData struct {
	Todos        []Todo
	Stats        Statistics
	Message      string
	MessageType  string
	TotalTodos   int
	PendingCount int
}
type Statistics struct {
	Total          int
	Completed      int
	Pending        int
	High           int
	Medium         int
	Low            int
	Overdue        int
	DueToday       int
	Tagged         int
	CompletionRate float64
}

var todos []Todo
var nextID = 1

const filename = "todos.json"

func addTodoItem(title string, priority Priority, dueDate *time.Time, tags []string) {
	newTodo := Todo{
		ID:        nextID,
		Title:     title,
		Completed: false,
		Priority:  priority,
		DueDate:   dueDate,
		Tags:      tags,
	}
	todos = append(todos, newTodo)
	nextID++
	saveTodos()
}
func toggleComplete(id int) {
	for i := range todos {
		if todos[i].ID == id {
			todos[i].Completed = !todos[i].Completed
			saveTodos()
			return
		}
	}
}
func deleteTodoItem(id int) {
	for i := range todos {
		if todos[i].ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			saveTodos()
			return
		}
	}
}
func sortByPriority() {
	sort.Slice(todos, func(i, j int) bool {
		if todos[i].Priority != todos[j].Priority {
			return todos[i].Priority > todos[j].Priority
		}
		return !todos[i].Completed && todos[j].Completed
	})
	saveTodos()
}
func sortByStatus() {
	sort.Slice(todos, func(i, j int) bool {
		return !todos[i].Completed && todos[j].Completed
	})
	saveTodos()
}
func calculateStatistics() Statistics {
	stats := Statistics{}
	stats.Total = len(todos)
	now := time.Now()
	for _, todo := range todos {
		if todo.Completed {
			stats.Completed++
		} else {
			stats.Pending++
		}
		switch todo.Priority {
		case Low:
			stats.Low++
		case Medium:
			stats.Medium++
		case High:
			stats.High++
		}
		if todo.DueDate != nil && !todo.Completed {
			if todo.DueDate.Before(now) {
				stats.Overdue++
			} else if todo.DueDate.Sub(now).Hours() < 24 {
				stats.DueToday++
			}
		}
		if len(todo.Tags) > 0 {
			stats.Tagged++
		}
	}
	if stats.Total > 0 {
		stats.CompletionRate = (float64(stats.Completed) / float64(stats.Total)) * 100
	}
	return stats
}
func searchTodos(keyword string) []Todo {
	keyword = strings.ToLower(strings.TrimSpace(keyword))
	var results []Todo
	for _, todo := range todos {
		if strings.Contains(strings.ToLower(todo.Title), keyword) {
			results = append(results, todo)
		}
	}
	return results
}
func filterByTag(tag string) []Todo {
	tag = strings.ToLower(strings.TrimSpace(tag))
	var results []Todo
	for _, todo := range todos {
		for _, t := range todo.Tags {
			if strings.ToLower(t) == tag {
				results = append(results, todo)
				break
			}
		}
	}
	return results
}
func saveTodos() {
	jsonData, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		log.Printf("Error marshaling todos: %v\n", err)
		return
	}
	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		log.Printf("Error saving todos: %v\n", err)
	}
}
func loadTodos() {
	jsonData, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			todos = []Todo{}
			return
		}
		log.Printf("Error reading file: %v\n", err)
		return
	}
	err = json.Unmarshal(jsonData, &todos)
	if err != nil {
		log.Printf("Error unmarshaling todos: %v\n", err)
		todos = []Todo{}
		return
	}
	maxID := 0
	for _, todo := range todos {
		if todo.ID > maxID {
			maxID = todo.ID
		}
	}
	nextID = maxID + 1
}

var tmpl *template.Template

func homeHandler(w http.ResponseWriter, r *http.Request) {
	stats := calculateStatistics()
	data := PageData{
		Todos:        todos,
		Stats:        stats,
		TotalTodos:   len(todos),
		PendingCount: stats.Pending,
	}
	tmpl.ExecuteTemplate(w, "index.html", data)
}
func addHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	title := strings.TrimSpace(r.FormValue("title"))
	if title == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	priorityStr := r.FormValue("priority")
	priority := Medium
	switch priorityStr {
	case "0":
		priority = Low
	case "2":
		priority = High
	}
	var dueDate *time.Time
	dueDateStr := r.FormValue("dueDate")
	if dueDateStr != "" {
		parsed, err := time.Parse("2006-01-02", dueDateStr)
		if err == nil {
			dueDate = &parsed
		}
	}
	tagsStr := r.FormValue("tags")
	tags := parseTags(tagsStr)
	addTodoItem(title, priority, dueDate, tags)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
func toggleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	idStr := r.FormValue("id")
	id, err := strconv.Atoi(idStr)
	if err == nil {
		toggleComplete(id)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
func deleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	idStr := r.FormValue("id")
	id, err := strconv.Atoi(idStr)
	if err == nil {
		deleteTodoItem(id)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
func sortHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	sortType := r.FormValue("type")
	if sortType == "priority" {
		sortByPriority()
	} else if sortType == "status" {
		sortByStatus()
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
func searchHandler(w http.ResponseWriter, r *http.Request) {
	keyword := strings.TrimSpace(r.FormValue("q"))
	if keyword == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	results := searchTodos(keyword)
	stats := calculateStatistics()
	data := PageData{
		Todos:       results,
		Stats:       stats,
		Message:     fmt.Sprintf("Search results for '%s' (%d found)", keyword, len(results)),
		MessageType: "info",
	}
	tmpl.ExecuteTemplate(w, "index.html", data)
}
func statsHandler(w http.ResponseWriter, r *http.Request) {
	stats := calculateStatistics()
	data := PageData{
		Todos: todos,
		Stats: stats,
	}
	tmpl.ExecuteTemplate(w, "stats.html", data)
}
func parseTags(input string) []string {
	if input == "" {
		return []string{}
	}
	rawTags := strings.Split(input, ",")
	var tags []string
	for _, tag := range rawTags {
		cleaned := strings.TrimSpace(tag)
		if cleaned != "" {
			tags = append(tags, cleaned)
		}
	}
	return tags
}
func runCLI() {
	fmt.Println("╔════════════════════════════════════════╗")
	fmt.Println("║   Welcome to Momentum v3.0            ║")
	fmt.Println("║   Web + CLI Dual Mode!                ║")
	fmt.Println("╚════════════════════════════════════════╝")
	fmt.Printf("📂 Loaded %d task(s)\n", len(todos))
	scanner := bufio.NewScanner(os.Stdin)
	for {
		displayMenu()
		choice := readInput(scanner)
		switch choice {
		case "1":
			fmt.Print("\nEnter task title: ")
			title := readInput(scanner)
			if title == "" {
				fmt.Println("❌ Title cannot be empty!")
				continue
			}
			priority := getPriority(scanner)
			fmt.Print("\nEnter due date (YYYY-MM-DD) or press Enter to skip: ")
			dateStr := readInput(scanner)
			var dueDate *time.Time
			if dateStr != "" {
				parsed, err := time.Parse("2006-01-02", dateStr)
				if err == nil {
					dueDate = &parsed
				}
			}
			fmt.Print("\nEnter tags (comma-separated) or press Enter to skip: ")
			tagsStr := readInput(scanner)
			tags := parseTags(tagsStr)
			addTodoItem(title, priority, dueDate, tags)
			fmt.Println("✅ Task added successfully!")
		case "2":
			listTodosCLI()
		case "3":
			listTodosCLI()
			if len(todos) == 0 {
				continue
			}
			fmt.Print("Enter task ID to complete: ")
			idStr := readInput(scanner)
			id, err := strconv.Atoi(idStr)
			if err == nil {
				toggleComplete(id)
				fmt.Println("✅ Task toggled!")
			}
		case "4":
			listTodosCLI()
			if len(todos) == 0 {
				continue
			}
			fmt.Print("Enter task ID to delete: ")
			idStr := readInput(scanner)
			id, err := strconv.Atoi(idStr)
			if err == nil {
				deleteTodoItem(id)
				fmt.Println("🗑️  Task deleted!")
			}
		case "5":
			sortByPriority()
			fmt.Println("✅ Sorted by priority!")
			listTodosCLI()
		case "6":
			sortByStatus()
			fmt.Println("✅ Sorted by status!")
			listTodosCLI()
		case "7":
			fmt.Print("\nEnter keyword to search: ")
			keyword := readInput(scanner)
			results := searchTodos(keyword)
			fmt.Printf("\n🔍 Found %d results:\n", len(results))
			for _, todo := range results {
				fmt.Printf("%d. [%v] %s - %s\n", todo.ID, todo.Completed, todo.Title, todo.Priority)
			}
		case "8":
			stats := calculateStatistics()
			displayStatsCLI(stats)
		case "9":
			fmt.Println("👋 Goodbye!")
			return
		default:
			fmt.Println("❌ Invalid choice! Please enter 1-9.")
		}
	}
}
func displayMenu() {
	fmt.Println("\n╔════════════════════════════════╗")
	fmt.Println("║        MOMENTUM MENU          ║")
	fmt.Println("╚════════════════════════════════╝")
	fmt.Println("1. Add a task")
	fmt.Println("2. List all tasks")
	fmt.Println("3. Toggle complete")
	fmt.Println("4. Delete a task")
	fmt.Println("5. Sort by priority")
	fmt.Println("6. Sort by status")
	fmt.Println("7. Search tasks")
	fmt.Println("8. Show statistics")
	fmt.Println("9. Quit")
	fmt.Print("\nEnter your choice (1-9): ")
}
func readInput(scanner *bufio.Scanner) string {
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}
func getPriority(scanner *bufio.Scanner) Priority {
	fmt.Println("\nSelect priority:")
	fmt.Println("1. 🟢 Low")
	fmt.Println("2. 🟡 Medium")
	fmt.Println("3. 🔴 High")
	fmt.Print("Enter priority (1-3, default 2): ")
	choice := readInput(scanner)
	switch choice {
	case "1":
		return Low
	case "3":
		return High
	default:
		return Medium
	}
}
func listTodosCLI() {
	if len(todos) == 0 {
		fmt.Println("\n❌ No tasks yet!")
		return
	}
	fmt.Println("\n📝 Your Tasks:")
	fmt.Println("──────────────────────────────────────────────────────")
	for _, todo := range todos {
		status := "[ ]"
		if todo.Completed {
			status = "[✓]"
		}
		fmt.Printf("%d. %s %-30s %s\n", todo.ID, status, todo.Title, todo.Priority)
	}
	fmt.Println("──────────────────────────────────────────────────────")
}
func displayStatsCLI(stats Statistics) {
	fmt.Println("\n╔══════════════════════════════════════════╗")
	fmt.Println("║       📊 STATISTICS DASHBOARD        ║")
	fmt.Println("╚══════════════════════════════════════════╝")
	fmt.Printf("\n📄 Total: %d | ✅ Completed: %d | 📅 Pending: %d\n", stats.Total, stats.Completed, stats.Pending)
	fmt.Printf("📈 Completion Rate: %.1f%%\n", stats.CompletionRate)
	fmt.Printf("\n🔴 High: %d | 🟡 Medium: %d | 🟢 Low: %d\n", stats.High, stats.Medium, stats.Low)
	if stats.Overdue > 0 {
		fmt.Printf("\n⚠️  Overdue: %d\n", stats.Overdue)
	}
	fmt.Println()
}
func main() {
	loadTodos()
	if len(os.Args) > 1 {
		command := os.Args[1]
		if command == "web" {
			fmt.Println("🌐 Starting Web Server...")
			fmt.Println("📂 Loaded", len(todos), "tasks from file")
			var err error
			tmpl, err = template.ParseGlob("templates/*.html")
			if err != nil {
				log.Fatal("Error parsing templates:", err)
			}
			http.HandleFunc("/", homeHandler)
			http.HandleFunc("/add", addHandler)
			http.HandleFunc("/toggle", toggleHandler)
			http.HandleFunc("/delete", deleteHandler)
			http.HandleFunc("/sort", sortHandler)
			http.HandleFunc("/search", searchHandler)
			http.HandleFunc("/stats", statsHandler)
			http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
			fmt.Println("✅ Server running at http://localhost:8080")
			fmt.Println("Press Ctrl+C to stop")
			log.Fatal(http.ListenAndServe(":8080", nil))
			return
		}
		if command == "cli" {
			runCLI()
			return
		}
	}
	fmt.Println("Usage:")
	fmt.Println("  go run main.go web    - Start web server")
	fmt.Println("  go run main.go cli    - Run CLI mode")
}

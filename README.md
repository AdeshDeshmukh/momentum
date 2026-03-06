# ✨ Momentum by Adesh

> **A full-stack productivity web application built with pure Go**  
> Features a stunning animated web UI with dual-mode support (Web + CLI) — all without any external frameworks or dependencies.

![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8?style=flat&logo=go)
![CI](https://github.com/AdeshDeshmukh/momentum/workflows/CI/badge.svg)
![Deployment](https://img.shields.io/badge/deployment-live-brightgreen?logo=render)
![License](https://img.shields.io/badge/license-MIT-green)
![Status](https://img.shields.io/badge/status-active-success)
![Made with Love](https://img.shields.io/badge/Made%20with-❤️-red)

---

## 🌐 Live Demo

**👉 [Try the Live App Here!](https://momentum-8ucf.onrender.com) 👈**

*The app is deployed on Render.com. First load may take 30 seconds if the app was sleeping.*

---

## 🎯 About

**Momentum** is a modern task management application that demonstrates full-stack development using only Go's standard library. Built from first principles to showcase:

- **Server-side rendering** with Go's `html/template`
- **RESTful web server** with `net/http`
- **Responsive UI** with modern CSS3 animations
- **Data persistence** with JSON
- **Dual-mode architecture** (Web + CLI)

Perfect for learning Go web development or as a portfolio project!

## ✨ Features

### Web Mode 🌐
- 🎨 **Beautiful Animated UI** - Gradient backgrounds, smooth transitions, glassmorphism effects
- 📊 **Analytics Dashboard** - Track productivity with completion rates and charts
- 🎯 **Priority Management** - Color-coded priority levels (Low, Medium, High)
- 📅 **Due Date Tracking** - Smart warnings for overdue and due-today tasks
- 🏷️ **Tag System** - Organize tasks with custom tags
- 🔍 **Search & Filter** - Find tasks instantly
- ↕️ **Smart Sorting** - Sort by priority or completion status
- 💾 **Auto-save** - All changes persist automatically
- 📱 **Responsive Design** - Works on mobile, tablet, and desktop

### CLI Mode ⌨️
- ⚡ **Fast Commands** - Quick task management from terminal
- 🎨 **Colored Output** - ANSI colors for better readability
- 📋 **Full Feature Parity** - All web features available in CLI
- 🔄 **Shared Data** - Seamlessly switch between web and CLI modes

## 🎬 Try It Yourself

### ☁️ Cloud Version (No Installation Required)

**Visit the live app:** [https://momentum-8ucf.onrender.com](https://momentum-8ucf.onrender.com)

- ✅ No setup needed
- ✅ Works on any device
- ✅ Try all features instantly
- ⚠️ First load takes ~30 seconds (free tier)

### 💻 Local Version (Full Features)

**Web Mode:**
```bash
go run main.go web
```
Then open your browser to **http://localhost:8080**

**CLI Mode:**
```bash
go run main.go cli
```

### Screenshots

| Web Interface | Analytics Dashboard |
|--------------|---------------------|
| ![Main Interface](screenshots/main.png) | ![Analytics](screenshots/analytics.png) |

*Note: After cloning, the app will start with sample data so you can explore features immediately.*

## 🚀 Quick Start

### Prerequisites
- Go 1.18 or higher ([Download Go](https://go.dev/dl/))
- A modern web browser (Chrome, Firefox, Safari, Arc)

### Installation

```bash
# Clone the repository
git clone https://github.com/AdeshDeshmukh/momentum.git
cd momentum

# Run in web mode (recommended for first-time users)
go run main.go web

# Open http://localhost:8080 in your browser
```

### Usage

#### Web Mode 🌐
```bash
# Start the web server
go run main.go web

# Open in browser
open http://localhost:8080  # macOS
start http://localhost:8080 # Windows
xdg-open http://localhost:8080 # Linux
```

**Features:**
- ➕ Add tasks with title, priority, due date, and tags
- ✅ Mark tasks complete with a single click
- 🗑️ Delete tasks easily
- 📊 View detailed analytics
- 🔍 Search and filter tasks
- ↕️ Sort by priority or status

#### CLI Mode ⌨️
```bash
# Interactive menu
go run main.go cli
✅ Completed: Review pull requests

$ go run main.go stats
╔══════════════════════════════════════════╗
║       📊 STATISTICS DASHBOARD        ║
╚══════════════════════════════════════════╝
📄 Overview:
   Total todos:        1
   ✅ Completed:        1
   📈 Completion rate:  100.0%
🎉 Great job! You're very productive!
```

## 🚀 Quick Start

### Prerequisites
- Go 1.25 or higher

### Installation

```bash
# Clone the repository
git clone https://github.com/AdeshDeshmukh/momentum.git
cd momentum

# Run directly
go run main.go

# Or build executable
go build -o todo main.go
./todo
```

## 📖 Usage

### Interactive Mode
Simply run without arguments for the full menu:
```bash
go run main.go
```

### CLI Mode (Power Users)
```bash
# Add a todo
go run main.go add "Task description"

# List all todos
go run main.go list

# Complete a todo by number
go run main.go complete 1

# Delete a todo
go run main.go delete 2

# Search todos
go run main.go search "keyword"

# Filter by tag
go run main.go filter "Work"

# Show statistics
go run main.go stats

# Help
go run main.go help
```

## 🎨 Features in Detail

### Priority Levels
- 🔴 **High** - Urgent tasks shown in red
- 🟡 **Medium** - Normal tasks shown in yellow  
- 🟢 **Low** - Low priority shown in green

### Due Date Alerts
- ⚠️ **OVERDUE** - Past deadlines highlighted in red
- ⏰ **Due today** - Today's deadlines in yellow
- 📅 **Upcoming** - Future dates in cyan

### Tags & Organization
Add multiple tags to organize tasks:
```
Work, Urgent, Meeting, Personal, Shopping
```

### Statistics Dashboard
Track your productivity with:
- Completion percentage
- Priority breakdown
- Overdue count
- Tagged items count

## � Technical Stack

### Backend
- **Language**: Go 1.25+
- **Web Server**: `net/http` (standard library HTTP server)
- **Templating**: `html/template` (server-side rendering)
- **Data Storage**: `encoding/json` (file-based JSON persistence)
- **I/O Operations**: `bufio`, `os` (file operations and CLI)
- **Time Management**: `time` (due date handling)
- **Routing**: HTTP handlers with form processing

### Frontend
- **HTML5**: Semantic markup, accessible forms
- **CSS3**: Modern features
  - Animations (@keyframes, transitions, transforms)
  - Gradients (linear, radial)
  - Flexbox and Grid layouts
  - Glassmorphism effects (backdrop-filter)
  - Responsive design (media queries)
- **No JavaScript**: 100% server-side rendering

### Architecture Highlights
- **Dual-Mode Design**: Shared business logic between Web and CLI
- **RESTful Routes**: Clean HTTP handlers for all operations
- **Template Methods**: Custom formatting functions in templates
- **Type Safety**: Strong typing with Go structs
- **Error Handling**: Comprehensive error management
- **Data Persistence**: Automatic JSON serialization

### Code Structure
```
main.go (541 lines of clean code)
├── Core Types
│   ├── Priority enum (Low=0, Medium=1, High=2)
│   └── Todo struct (ID, Title, Completed, Priority, DueDate, Tags)
├── HTTP Handlers
│   ├── homeHandler() - Main page
│   ├── addHandler() - Create tasks
│   ├── toggleHandler() - Toggle completion
│   ├── deleteHandler() - Remove tasks
│   ├── sortHandler() - Sort operations
│   ├── searchHandler() - Search functionality
│   └── statsHandler() - Analytics dashboard
├── CRUD Operations
│   ├── addTodoItem() - Create
│   ├── toggleComplete() - Update
│   ├── deleteTodoItem() - Delete
│   └── searchTodos() - Query
├── Analytics
│   ├── calculateStatistics()
│   ├── sortByPriority()
│   └── sortByStatus()
├── Persistence Layer
│   ├── saveTodos() - JSON marshaling
│   └── loadTodos() - JSON unmarshaling
└── CLI Interface
    ├── runCLI() - Interactive menu
    ├── displayMenu()
    └── listTodosCLI()
```

## 🏗️ Project Structure
```
momentum/
├── main.go                    # Main application (541 lines)
├── templates/
│   ├── index.html            # Main web interface
│   └── stats.html            # Analytics dashboard
├── static/
│   └── style.css             # Complete styling (1000+ lines)
├── go.mod                    # Go module definition
├── .gitignore               # Git ignore rules
├── README.md                # Project documentation
├── LICENSE                  # MIT License
├── PUSH-TO-GITHUB.md        # GitHub deployment guide
├── GITHUB-PUSH-CHECKLIST.md # Pre-deployment checklist
├── DEPLOYMENT.md            # Cloud deployment guide
└── todos.json               # Data file (auto-created)
```

## 💡 Technical Decisions

### Why Go Standard Library Only?
- **Simplicity**: No dependency hell
- **Performance**: Lightweight and fast
- **Learning**: Understanding fundamentals
- **Portability**: Works everywhere Go runs
- **Maintenance**: No breaking changes from dependencies

### Why Server-Side Rendering?
- **Simplicity**: No complex frontend framework
- **SEO Friendly**: HTML rendered on server
- **Fast Load**: No large JavaScript bundles
- **Progressive Enhancement**: Works without JS
- **Security**: Less client-side attack surface

### Why JSON File Storage?
- **Simplicity**: No database setup required
- **Portable**: Single file, easy backup
- **Human Readable**: Easy to inspect/edit
- **Version Control**: Can track changes in git
- **Zero Config**: Works out of the box

## 🎓 Learning Resources & Concepts

### Go Concepts Demonstrated
- ✅ **Structs and Methods** - Custom types with behavior
- ✅ **Enums with iota** - Type-safe constants
- ✅ **Slices** - Dynamic arrays and manipulation
- ✅ **Pointers** - Optional fields with *time.Time
- ✅ **HTTP Handlers** - Web server fundamentals
- ✅ **Template Rendering** - Server-side HTML generation
- ✅ **Form Handling** - POST/GET request processing
- ✅ **File I/O** - Reading and writing JSON
- ✅ **Error Handling** - Proper error management patterns
- ✅ **Sorting** - Custom sort functions
- ✅ **String Manipulation** - Parsing and formatting

### Web Development Concepts
- ✅ **RESTful Design** - HTTP methods (GET/POST)
- ✅ **Server-Side Rendering** - Template-based HTML generation
- ✅ **Static File Serving** - CSS/JS delivery
- ✅ **Form Processing** - User input handling
- ✅ **HTTP Routing** - URL pattern matching
- ✅ **Session State** - File-based persistence
- ✅ **Responsive Design** - Mobile-first CSS

### CSS Techniques Used
- ✅ **Animations** - @keyframes, transitions, transforms
- ✅ **Gradients** - Linear and radial backgrounds
- ✅ **Flexbox** - Modern layout system
- ✅ **Grid** - Two-dimensional layouts
- ✅ **Glassmorphism** - Frosted glass effects
- ✅ **Media Queries** - Responsive breakpoints
- ✅ **Pseudo-elements** - ::before, ::after
- ✅ **Custom Properties** - CSS variables (optional)

### Architecture Patterns
- ✅ **MVC-like Structure** - Separation of concerns
- ✅ **Data Persistence** - CRUD operations
- ✅ **Template Composition** - Reusable HTML components
- ✅ **Dual-Mode Design** - Shared business logic
- ✅ **Type Safety** - Strong typing throughout

## 🚀 Deployment

### Live Production Instance

**Momentum is deployed on Render.com:** [https://momentum-8ucf.onrender.com](https://momentum-8ucf.onrender.com)

### Deployment Pipeline
- ✅ **CI/CD**: GitHub Actions runs tests on every push
- ✅ **Auto-Deploy**: Render automatically deploys when tests pass
- ✅ **Multi-Version Testing**: Tests run on Go 1.21, 1.22, 1.23
- ✅ **Zero Downtime**: Rolling deployments with health checks
- ✅ **Environment Config**: Dynamic port configuration for cloud hosting

### Infrastructure
- **Platform**: Render.com (Free Tier)
- **Runtime**: Go 1.26.0
- **Build**: `go build -tags netgo -ldflags '-s -w' -o app`
- **Start**: `./app web`
- **Health Check**: `GET /`
- **Region**: US West (Oregon)

### Deployment Features
- 🔄 **Automatic deploys** from `main` branch
- 📊 **Real-time build logs** and monitoring
- 🌍 **Global CDN** for static assets
- 🔒 **HTTPS** enabled by default
- 📈 **Uptime monitoring** via Render dashboard

### Deploy Your Own
```bash
# Fork this repo, then:
1. Sign up at render.com (free)
2. Connect your GitHub account
3. Select your forked repository
4. Render auto-detects configuration from render.yaml
5. Click "Create Web Service"
6. Your app is live in ~5 minutes!
```

### Environment Variables
The app reads `PORT` from environment (provided by Render) and falls back to `8080` for local development.

## 🤝 Contributing

Contributions are welcome! Feel free to:
- Report bugs
- Suggest features
- Submit pull requests

## 📝 License

MIT License - feel free to use this project for learning or personal use.

## 👤 Author

**Adesh Kishor Deshmukh**

- 💼 LinkedIn: [linkedin.com/in/adesh-deshmukh](https://www.linkedin.com/in/adesh-deshmukh/)
- 🐙 GitHub: [@AdeshDeshmukh](https://github.com/AdeshDeshmukh)
- 📧 Email: adeshkd123@gmail.com

---

## 💼 Why This Project Stands Out for Resumes

### Key Selling Points:
1. **Full-Stack Development** - Built both frontend and backend from scratch
2. **Zero Dependencies** - Deep understanding of fundamentals, no framework magic
3. **Production-Ready** - Clean code, error handling, persistence, deployed live
4. **Modern UI/UX** - Advanced CSS animations and responsive design
5. **Architecture Skills** - Dual-mode design with shared logic
6. **DevOps & CI/CD** - GitHub Actions, automated testing, continuous deployment
7. **Testing** - Comprehensive test suite with 67.3% coverage
8. **Cloud Deployment** - Live production app on Render.com
9. **Best Practices** - Clean code, documentation, version control

### Perfect for Interviews:
- Can explain every line of code (no magic frameworks)
- Demonstrates HTTP, templates, and web fundamentals
- Shows CSS mastery without JavaScript dependencies
- Proves ability to build complete applications solo
- **Live demo available** - recruiters can try it immediately
- CI/CD pipeline shows DevOps understanding
- Easy to run and demo during technical interviews
- Professional README and documentation

## 🌟 Acknowledgments

Built from first principles to demonstrate clean, idiomatic Go code and modern web development practices without relying on external frameworks.

**Technologies Mastered:**
- Go standard library (`net/http`, `html/template`, `encoding/json`)
- HTML5 semantic markup
- CSS3 animations and responsive design
- Server-side rendering architecture
- RESTful API design
- File-based persistence

---

**⭐ If you find this project useful, please give it a star on GitHub!**

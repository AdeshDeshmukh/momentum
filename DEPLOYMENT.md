# 🚀 How to Make Your CLI App Live

## Step 1: Push to GitHub

```bash
# Initialize git (if not already done)
cd /Users/adeshkishordeshmukh/Documents/Projects/cli-todo-app
git init

# Add all files
git add .

# Commit
git commit -m "Initial commit: Momentum — full-stack Go productivity app"

# Create repo on GitHub (go to github.com/new)
# Then connect and push:
git branch -M main
git remote add origin https://github.com/YOUR_USERNAME/cli-todo-app.git
git push -u origin main
```

## Step 2: Enable GitHub Pages for Demo Website

1. Push `demo.html` to your repo
2. Go to Settings → Pages
3. Source: Deploy from branch → `main` → `/` (root)
4. Save
5. Your demo will be live at: `https://YOUR_USERNAME.github.io/cli-todo-app/demo.html`

## Step 3: Make it Interactive on Replit

### Option A: Import from GitHub
1. Go to https://replit.com/
2. Click "Create Repl"
3. Choose "Import from GitHub"
4. Paste: `https://github.com/YOUR_USERNAME/cli-todo-app`
5. Click "Import"
6. Hit "Run" → Your app runs in browser!
7. Get shareable link from URL bar

### Option B: Create New Repl
1. Go to https://replit.com/
2. Create new Repl → Language: Go
3. Upload your `main.go` and `go.mod`
4. Click "Run"
5. Share the Repl link

## Step 4: Record Demo GIF

### Using asciinema (Terminal Recording)
```bash
# Install
brew install asciinema

# Record
asciinema rec demo.cast

# Run your app and demonstrate features
# Press Ctrl+D when done

# Upload to asciinema.org
asciinema upload demo.cast

# You'll get a shareable link like:
# https://asciinema.org/a/XXXXX
```

Add to README:
```markdown
[![asciicast](https://asciinema.org/a/XXXXX.svg)](https://asciinema.org/a/XXXXX)
```

### Using VHS (Creates GIF)
```bash
# Install
brew install vhs

# Create script
cat > demo.tape << 'EOF'
Output demo.gif
Set FontSize 16
Set Width 1200
Set Height 700
Set Theme "Dracula"

Type "clear"
Enter
Sleep 500ms

Type "# Momentum Demo"
Enter
Sleep 1s

Type "go run main.go add 'Fix critical bug'"
Enter
Sleep 1s

Type "go run main.go list"
Enter
Sleep 2s

Type "go run main.go stats"
Enter
Sleep 2s
EOF

# Generate
vhs demo.tape

# This creates demo.gif
# Upload to repo and add to README
```

## Step 5: Create Releases with Binaries

```bash
# Build for all platforms
./build-all.sh  # or manually:

GOOS=darwin GOARCH=amd64 go build -o builds/todo-mac-intel main.go
GOOS=darwin GOARCH=arm64 go build -o builds/todo-mac-m1 main.go
GOOS=linux GOARCH=amd64 go build -o builds/todo-linux main.go
GOOS=windows GOARCH=amd64 go build -o builds/todo-windows.exe main.go

# Create GitHub Release
# Go to GitHub → Releases → New Release
# Tag: v1.0.0
# Upload all binaries from builds/
```

## Step 6: Update README with Live Links

```markdown
## 🌐 Live Demo

**Try it in your browser (no installation needed):**
- 🌍 [Interactive Demo on Replit](https://replit.com/@yourusername/cli-todo-app)
- 🎬 [Watch Demo Video](https://asciinema.org/a/XXXXX)
- 🖥️ [Web Demo Page](https://yourusername.github.io/cli-todo-app/demo.html)

**Download:**
- 📦 [Get Pre-built Binaries](https://github.com/yourusername/cli-todo-app/releases)
```

## For Your Resume

Add this to your resume projects section:

```
Momentum | Go, net/http, HTML/CSS, JSON | [GitHub] [Live Demo]
• Built a full-stack productivity web application using pure Go (net/http, html/template)
• Implemented priority system, due dates, tags, search, and statistics dashboard
• Designed dual-mode interface: interactive menu and fast CLI commands
• Used Go standard library for JSON persistence, sorting, and ANSI colors
• 100% test coverage with comprehensive error handling
• Live demo available at: [replit link]
```

## Best Practices

✅ Write detailed README with screenshots/GIFs
✅ Include installation and usage instructions
✅ Add badges (Go version, license, build status)
✅ Create releases with pre-built binaries
✅ Add LICENSE file (MIT recommended)
✅ Include demo.html for visual presentation
✅ Host on Replit for interactive demo
✅ Record terminal demo with asciinema or VHS
✅ Write good commit messages
✅ Tag releases semantically (v1.0.0, v1.1.0, etc.)

## Example GitHub URLs

Replace with your actual username:
- Repo: `https://github.com/adeshkishordeshmukh/cli-todo-app`
- Demo: `https://adeshkishordeshmukh.github.io/cli-todo-app/demo.html`
- Replit: `https://replit.com/@adeshkishordeshmukh/cli-todo-app`
- Releases: `https://github.com/adeshkishordeshmukh/cli-todo-app/releases`

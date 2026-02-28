# 🚀 GitHub Push Checklist

Complete this checklist before pushing to GitHub to ensure everything works perfectly!

---

## ✅ Pre-Push Checklist

### 1. Update Personal Information

- [ ] Update `go.mod` with your actual GitHub username
  ```bash
  # Change: module github.com/yourusername/momentum
  # To:     module github.com/YOUR-ACTUAL-USERNAME/momentum
  ```

- [ ] Update README.md author section with your real contact info:
  - [ ] Email address
  - [ ] LinkedIn profile
  - [ ] GitHub username
  - [ ] Portfolio website

- [ ] Update README.md clone URL:
  ```bash
  # Change: git clone https://github.com/yourusername/momentum.git
  # To:     git clone https://github.com/YOUR-USERNAME/momentum.git
  ```

### 2. Test Application Functionality

**Web Mode:**
- [ ] Start web server: `go run main.go web`
- [ ] Open http://localhost:8080 in browser
- [ ] Test adding a task
- [ ] Test completing a task
- [ ] Test deleting a task
- [ ] Test adding task with priority (Low/Medium/High)
- [ ] Test adding task with due date
- [ ] Test adding task with tags
- [ ] Test search functionality
- [ ] Test sort by priority
- [ ] Test sort by status
- [ ] Click "Full Statistics" - verify analytics dashboard loads
- [ ] Test responsive design (resize browser window)
- [ ] Verify animations are working (gradients, bounce, hover effects)

**CLI Mode:**
- [ ] Start CLI: `go run main.go cli`
- [ ] Test adding a task from menu
- [ ] Test listing tasks
- [ ] Test completing a task
- [ ] Verify colored output works
- [ ] Test statistics display

**Data Persistence:**
- [ ] Add some tasks in web mode
- [ ] Close browser
- [ ] Reopen http://localhost:8080
- [ ] Verify tasks are still there
- [ ] Switch to CLI mode: `go run main.go cli`
- [ ] List tasks - verify same tasks appear
- [ ] Add task in CLI
- [ ] Switch to web - verify new task appears

### 3. Clean Up Repository

- [ ] Remove unnecessary files:
  ```bash
  rm -f *.bak
  rm -f main_cli_backup.go.bak main_cli_only.go.bak
  rm -f momentum  # Remove compiled binary
  ```

- [ ] Optional: Clear test data (start fresh for GitHub):
  ```bash
  rm -f todos.json  # This will be recreated when app runs
  ```

- [ ] Check what files will be committed:
  ```bash
  git status
  ```

### 4. Verify File Structure

Ensure you have these essential files:
- [ ] `main.go` (main application)
- [ ] `templates/index.html` (web interface)
- [ ] `templates/stats.html` (analytics dashboard)
- [ ] `static/style.css` (styling)
- [ ] `go.mod` (Go module) 
- [ ] `README.md` (project documentation)
- [ ] `.gitignore` (ignore rules)
- [ ] `LICENSE` (MIT license)

Optional but recommended:
- [ ] `DEPLOYMENT.md` (deployment guide)
- [ ] `demo.html` (static demo page)

### 5. Test Build

**Build for your platform:**
```bash
go build -o momentum main.go
./momentum web
```

- [ ] Binary builds successfully
- [ ] Binary runs correctly
- [ ] Clean up: `rm momentum`

**Test cross-platform builds (optional):**
```bash
GOOS=darwin GOARCH=amd64 go build -o momentum-mac main.go
GOOS=linux GOARCH=amd64 go build -o momentum-linux main.go
GOOS=windows GOARCH=amd64 go build -o momentum.exe main.go
```

- [ ] All builds complete without errors
- [ ] Clean up: `rm momentum-* *.exe`

### 6. Initialize Git Repository

If not already initialized:
```bash
git init
```

- [ ] Git repository initialized

### 7. Create GitHub Repository

1. Go to [github.com/new](https://github.com/new)
2. Repository name: `momentum`
3. Description: "A full-stack productivity web app built with pure Go"
4. Choose: **Public** (for portfolio/resume visibility)
5. **Do NOT** initialize with README (you already have one)
6. Click "Create repository"

- [ ] GitHub repository created
- [ ] Repository is set to Public
- [ ] Note your repository URL: `https://github.com/YOUR-USERNAME/momentum.git`

### 8. Add and Commit Files

```bash
# Add all files
git add .

# Check what's being committed
git status

# Commit with descriptive message
git commit -m "Initial commit: Full-stack productivity app with Go"
```

- [ ] All necessary files added
- [ ] Files committed successfully
- [ ] `.gitignore` is working (todos.json, binaries not included)

### 9. Push to GitHub

```bash
# Add GitHub as remote
git remote add origin https://github.com/YOUR-USERNAME/momentum.git

# Push to main branch
git branch -M main
git push -u origin main
```

- [ ] Remote added successfully
- [ ] Pushed to GitHub successfully

### 10. Verify on GitHub

Visit your repository: `https://github.com/YOUR-USERNAME/momentum`

- [ ] README.md displays correctly with formatting
- [ ] All code files are present
- [ ] Folder structure is correct (templates/, static/)
- [ ] badges display correctly in README
- [ ] LICENSE file is visible
- [ ] Repository description shows up

### 11. Test Clone (Verification)

In a different directory:
```bash
cd /tmp
git clone https://github.com/YOUR-USERNAME/momentum.git
cd momentum
go run main.go web
```

- [ ] Repository clones successfully
- [ ] App runs without errors
- [ ] Browser opens and shows correct interface
- [ ] Clean up: `cd ~ && rm -rf /tmp/momentum`

### 12. Optional: Add Topics/Tags on GitHub

On your GitHub repository page:
1. Click "⚙️ Manage topics"
2. Add relevant tags:
   - `go`
   - `golang`
   - `web-app`
   - `todo-app`
   - `productivity`
   - `full-stack`
   - `server-side-rendering`
   - `html-templates`
   - `css3-animations`
   - `first-principles`

- [ ] Topics added to repository

### 13. Optional: Add Screenshot

Take a screenshot of your app and add it to README:

1. Take screenshot while app is running
2. Create `screenshots` folder:
   ```bash
   mkdir screenshots
   ```
3. Save screenshot as `screenshots/main.png`
4. Update README.md screenshot section ( remove the line or add actual screenshot)
5. Commit and push:
   ```bash
   git add screenshots/
   git commit -m "Add screenshots"
   git push
   ```

- [ ] Screenshots added (optional)

### 14. Enable GitHub Pages (Optional)

If you want to host your static demo.html:

1. Go to repository Settings
2. Click "Pages" in sidebar
3. Source: Deploy from a branch
4. Branch: main, folder: / (root)
5. Save
6. Wait a few minutes
7. Visit: `https://YOUR-USERNAME.github.io/momentum/demo.html`

- [ ] GitHub Pages enabled (optional)

---

## 🎉 Success!

Your project is now live on GitHub!

### Share Your Project:

1. **Add to your resume:**
   ```
   Momentum - Full-Stack Productivity Web App
   • Built modern task management app using Go's standard library (net/http, html/template)
   • Implemented server-side rendering with responsive CSS3 animations
   • Developed dual-mode architecture (Web UI + CLI) with shared JSON persistence
   • Tech: Go, HTML5, CSS3, Server-Side Rendering
   • GitHub: github.com/YOUR-USERNAME/momentum
   ```

2. **Update your LinkedIn:**
   - Add to Projects section
   - Include link to GitHub repository
   - Mention key technologies: Go, Full-Stack, Web Development

3. **Tweet about it:**
   ```
   Just built Momentum - a full-stack productivity app using pure Go! 
   🚀 No frameworks
   ✨ Beautiful animations
   🎨 Server-side rendering
   
   Check it out: github.com/YOUR-USERNAME/momentum
   
   #golang #webdev #buildinpublic
   ```

---

## 📝 Next Steps

- Star your own repository (shows you're proud of it!)
- Watch for issues/feedback
- Consider adding more features from the roadmap
- Share with Go community
- Apply to jobs mentioning this project!

---

## 🐛 Troubleshooting

**If push fails:**
```bash
# Make sure you're on main branch
git branch

# Check remote URL
git remote -v

# If wrong, update it
git remote set-url origin https://github.com/YOUR-USERNAME/momentum.git
```

**If files are missing on GitHub:**
```bash
# Check .gitignore isn't excluding them
cat .gitignore

# Force add if needed
git add -f filename
git commit -m "Add missing file"
git push
```

---

<div align="center">
  <p><strong>🎊 Congratulations on your GitHub deployment! 🎊</strong></p>
  <p>Your portfolio just got stronger!</p>
</div>

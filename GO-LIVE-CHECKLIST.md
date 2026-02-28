# ✅ Go-Live Checklist

Before pushing to GitHub and making your project public, complete these steps:

## 📝 Customize Your Files

- [ ] Update `README.md`:
  - Replace `yourusername` with your actual GitHub username
  - Replace `[Your Name]` and add your LinkedIn/email
  - Update any placeholder text

- [ ] Update `demo.html`:
  - Replace `yourusername` with your GitHub username in all links
  - Customize the hero text if desired

- [ ] Update `LICENSE`:
  - Replace `[Your Name]` with your actual name

- [ ] Update `go.mod`:
  - Replace `github.com/yourusername/cli-todo-app` with actual username

## 🚀 GitHub Push (First Time)

```bash
# From the cli-todo-app directory:
cd /Users/adeshkishordeshmukh/Documents/Projects/cli-todo-app

# Initialize Git (skip if already done)
git init

# Add all files
git add .

# Commit
git commit -m "Initial commit: Momentum — full-stack Go productivity app"

# Create new repo on GitHub:
# Go to https://github.com/new
# Name: cli-todo-app
# Description: "A full-stack productivity web app built with Go. Features animated web UI and CLI mode."
# Public (so it can be on your resume)
# DON'T initialize with README (we already have one)

# Connect to GitHub and push
git branch -M main
git remote add origin https://github.com/YOUR_USERNAME/cli-todo-app.git
git push -u origin main
```

## 🌐 Enable GitHub Pages

- [ ] Go to your repo → Settings → Pages
- [ ] Source: Deploy from branch
- [ ] Branch: `main` → `/` (root) → Save
- [ ] Wait 2-3 minutes
- [ ] Visit: `https://YOUR_USERNAME.github.io/cli-todo-app/demo.html`
- [ ] Add this link to your resume!

## 🎬 Create Demo Recording (Optional but Recommended)

Install asciinema:
```bash
brew install asciinema
```

Record demo:
```bash
asciinema rec demo.cast
# Now use your app - add todos, complete them, show stats
# Press Ctrl+D when done
asciinema upload demo.cast
# Save the URL and add to README
```

OR use VHS for GIF:
```bash
brew install vhs
# See DEPLOYMENT.md for VHS script
```

## 📦 Create Binaries for Download

```bash
./build.sh
# This creates builds/ folder with binaries for Mac, Linux, Windows
```

Then create GitHub Release:
- [ ] Go to repo → Releases → "Create a new release"
- [ ] Tag: `v1.0.0`
- [ ] Title: `v1.0.0 - Initial Release`
- [ ] Description: "First stable release with all 10 features"
- [ ] Upload all files from `builds/` folder
- [ ] Click "Publish release"

## 🎮 Deploy to Replit (Optional - Great for Resume!)

- [ ] Go to https://replit.com/
- [ ] Sign up / Log in
- [ ] Create new Repl → Import from GitHub
- [ ] Enter: `https://github.com/YOUR_USERNAME/cli-todo-app`
- [ ] Click Run
- [ ] Copy the Repl URL
- [ ] Add to README as "Try it live"

## 📄 Resume Entry

Add this to your resume's Projects section:

```
Momentum | Go, net/http, HTML/CSS, JSON
[GitHub Link] [Live Demo]

• Developed full-featured command-line task management app
• Implemented priority system, due dates, tags, search functionality
• Built dual interface: interactive menu + CLI commands
• Achieved persistence with JSON, sorting, and statistics features
• Deployed live demo on Replit with cross-platform binaries
```

## 🔍 Final Checks

- [ ] All files committed and pushed to GitHub
- [ ] README displays correctly on GitHub
- [ ] demo.html loads on GitHub Pages
- [ ] All links in README work
- [ ] License file is present
- [ ] .gitignore prevents todos.json from being pushed
- [ ] Build script works (`./build.sh`)
- [ ] Project is public on GitHub
- [ ] Added project link to resume/portfolio

## 🎉 Post-Launch

- [ ] Star your own repo (why not!)
- [ ] Share on LinkedIn with demo link
- [ ] Add project to your portfolio website
- [ ] Tweet about your first Go project (optional)
- [ ] Apply to jobs and reference this project

---

## Need Help?

Common issues:
- **Build fails**: Make sure you're in the project directory and Go is installed
- **Git push rejected**: Check if you created the repo correctly on GitHub
- **Demo not showing**: GitHub Pages can take 2-3 minutes to deploy
- **Colors not working on Windows**: Use Windows Terminal instead of CMD

Good luck with your job search! 🚀

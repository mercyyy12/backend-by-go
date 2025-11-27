# Quick Git & GitHub Guide

This is a **quick-start guide** to install Git and push your code to GitHub.
By following these steps, you can push your local project to GitHub successfully.

> A full Git/GitHub guide (branches, merges, pull requests) will be added soon.

## Git Command Reference

| Command | Purpose |
|---------|---------|
| `git init` | Start Git tracking in the folder |
| `git status` | Show what files changed or staged |
| `git add .` | Stage all files for commit |
| `git commit -m "message"` | Save a version of your code |
| `git branch -M main` | Rename current branch to main |
| `git remote add origin <URL>` | Connect folder to GitHub repo |
| `git push -u origin main` | Upload code to GitHub |

---

## Installation & Setup (Linux)

```bash
sudo apt update
sudo apt install git
git --version
git config --global user.name "Your Name"
git config --global user.email "your@email.com"
````

---

## Push a New Local Project to GitHub

```bash
# Initialize Git
git init

# Check status
git status

# Stage files
git add .

# Commit changes
git commit -m "initial commit"

# Rename branch to main
git branch -M main

# Link to GitHub repo
git remote add origin https://github.com/username/repo.git

# Push to GitHub
git push -u origin main
```

---

## Push Changes to a Cloned Repo

```bash
# Check status
git status

# Stage changes
git add .

# Commit changes
git commit -m "your message"

# Push updates
git push

**Quick Note:**
This is a quick-start guide to install Git and push your code to GitHub.

If you follow the steps you'll be able to push your local project to GitHub successfully.

A full Git/GitHub guide (with branches, merges, pull requests) will be added soon.


## What each command means
git init = Start Git tracking in this folder
git status = Show what’s changed, staged, or untracked
git add . = Stage all files to be saved
git commit -m "your commit message" = Save a version of your code
git branch -M main = Rename the current branch to 'main'
git remote add origin YOUR_LINK = Connect this folder to a GitHub repository
git push -u origin main = Upload your code to GitHub and set it as default


# Git Installation & Setup (Linux)
# 1. Install Git
sudo apt update
sudo apt install git

# 2. Verify installation
git --version
# Output: git version x.x.x

# 3. Configure Git identity (used for all commits)
git config --global user.name "Your Name"
git config --global user.email "your@email.com"

# To push your local project to GitHub
# Follow: A or B.

# A. If Starting from Scratch (New Local Project)
# Only use this if you're creating a new project.
# Don't run 'git init' in a cloned repo.

# 4. Initialize Git
git init

# 5. Check status
git status

# 6. Stage your files
git add .

# 7. Double-check what’s staged
git status

# 8. Commit your code
git commit -m "initial commit"

# 9. Rename branch to 'main'
git branch -M main

# 10. Link to your GitHub repo
git remote add origin https://github.com/your-username/your-repo.git

# 11. Push to GitHub
git push -u origin main


# B. If You Cloned from GitHub
# Skip git init and git remote add. Your project is already a Git repo.

# Check file status
git status

# Stage changes
git add .

# Commit
git commit -m "your message"

# Push to GitHub
git push

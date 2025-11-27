# VS Code Quickstart (Ubuntu)

**Quick Note:**
This is a simple guide to install Visual Studio Code (VS Code) and prepare it for Go and Git development on a Linux system.

You’ll also install useful extensions.

# 1. Install VS Code via Snap
sudo snap install code --classic

# 2. Verify installation
code --version
# Output: code version x.x.x

# 3. Launch VS Code from terminal
code

# 4. Install Extensions
Inside VS Code: Open the Extensions tab (Ctrl + Shift + X), search and install the following:

Go = Official Go extension by the Go Team
GitLens = (optional) commit history, blame, authorship
Code Runner = (Optional) Run code snippets easily
Docker = (Optional) Docker integration and management
REST Client = (Optional) Test APIs from within VS Code

# 5. Configure Auto Save & Format on Save
# Open settings (Ctrl + ,) or open settings.json and paste the following inside settings.json:

{
   "files.autoSave": "afterDelay",
   "editor.formatOnSave": true,
   "go.formatTool": "gopls"
}

Done!!

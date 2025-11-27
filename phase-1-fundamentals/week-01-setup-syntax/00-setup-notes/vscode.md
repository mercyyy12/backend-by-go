# VS Code Quickstart (Ubuntu)

**Quick Note:**
This is a simple guide to install Visual Studio Code (VS Code) and prepare it for Go and Git development on a Linux system.

You’ll also install useful extensions.

# 1. Install VS Code
sudo snap install code --classic

# 2. Verify Installation
code --version
# Example: code version 1.xx.x

# 3. Launch VS Code
code

# 4. Install Extensions (inside VS Code, Ctrl+Shift+X)
# - Go (Official Go extension by the Go Team)
# - GitLens (optional: commit history, blame, authorship)
# - Code Runner (optional: run code snippets easily)
# - Docker (optional: Docker integration)
# - REST Client (optional: test APIs from within VS Code)

# 5. Configure Auto Save & Format on Save
# Open Settings → Open JSON (or Ctrl + , then Open Settings JSON) and add:
echo '{
    "files.autoSave": "afterDelay",
    "editor.formatOnSave": true,
    "go.formatTool": "gopls"
}' >> ~/.config/Code/User/settings.json

Done!!

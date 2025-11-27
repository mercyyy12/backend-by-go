# VS Code Quickstart (Ubuntu)

**Quick Note:**
This is a simple guide to install Visual Studio Code (VS Code) and prepare it for Go and Git development on a Linux system.

You’ll also install useful extensions.

```bash
# 1. Install VS Code via Snap
sudo snap install code --classic

# 2. Verify installation
code --version

# 3. Launch VS Code
code

# 4. Install Extensions (inside VS Code: Ctrl + Shift + X)
Go           -> Official Go extension by the Go Team
GitLens      -> Optional: commit history, blame, authorship
Code Runner  -> Optional: run code snippets easily
Docker       -> Optional: Docker integration
REST Client  -> Optional: test APIs from within VS Code

# 5. Configure Auto Save & Format on Save
# Open Settings JSON (Ctrl + , -> Open Settings JSON) and add:
cat <<EOL >> ~/.config/Code/User/settings.json
{
    "files.autoSave": "afterDelay",
    "editor.formatOnSave": true,
    "go.formatTool": "gopls"
}
EOL

Done!!

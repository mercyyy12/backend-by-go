# Go Installation on Linux

> Replace 'X' with the latest Go version, e.g., `1.25.0`.

```bash
# 1. Download Go
wget https://go.dev/dl/go1.X.X.linux-amd64.tar.gz

# 2. Remove previous Go installation (optional)
sudo rm -rf /usr/local/go

# 3. Extract Go to /usr/local
sudo tar -C /usr/local -xzf go1.X.X.linux-amd64.tar.gz

# 4. Add Go to PATH
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc

# 5. Verify installation
go version

# Optional: check environment
go env

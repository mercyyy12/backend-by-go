# This instruction is for Linux only!!

# 1. Download Go (latest)
wget https://go.dev/dl/go1.22.3.linux-amd64.tar.gz

# 2. Remove any previous Go installation (optional)
sudo rm -rf /usr/local/go

# 3. Extract to /usr/local
sudo tar -C /usr/local -xzf go1.22.3.linux-amd64.tar.gz

# 4. Add Go to PATH
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc

# 5. Verify installation
go version
# Expected output: go version go1.22.3 linux/amd64

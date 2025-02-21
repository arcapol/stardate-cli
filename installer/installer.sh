#!/bin/bash
set -e

# Ensure the script is run with root privileges.
if [ "$EUID" -ne 0 ]; then
    echo "Please run this installer with sudo or as root."
    exit 1
fi

# Set GitHub repository details.
GITHUB_USER="YOUR_GITHUB_USERNAME"
REPO="stardate-cli"

# Detect OS.
OS=$(uname -s)
case "$OS" in
    Linux)
        OS="linux"
        ;;
    Darwin)
        OS="macos"
        ;;
    *)
        echo "Unsupported OS: $OS"
        exit 1
        ;;
esac

# Detect Architecture.
ARCH=$(uname -m)
case "$ARCH" in
    x86_64)
        ARCH="amd64"
        ;;
    arm64|aarch64)
        ARCH="arm64"
        ;;
    *)
        echo "Unsupported architecture: $ARCH"
        exit 1
        ;;
esac

echo "Detected OS: $OS, Architecture: $ARCH"

# Define the asset name based on OS and architecture.
ASSET="stardate-${OS}-${ARCH}.tar.gz"
DOWNLOAD_URL="https://github.com/${GITHUB_USER}/${REPO}/releases/latest/download/${ASSET}"

echo "Downloading ${ASSET} from ${DOWNLOAD_URL}..."

# Download the asset using curl or wget.
if command -v curl >/dev/null 2>&1; then
    curl -L -o "${ASSET}" "${DOWNLOAD_URL}"
elif command -v wget >/dev/null 2>&1; then
    wget -O "${ASSET}" "${DOWNLOAD_URL}"
else
    echo "Error: curl or wget is required to download the binary."
    exit 1
fi

echo "Download complete. Extracting the archive..."
tar -xzvf "${ASSET}"

# Check that the extracted binary exists. The archive should contain a binary named 'stardate'.
if [ ! -f stardate ]; then
    echo "Error: Extracted binary 'stardate' not found."
    exit 1
fi

echo "Installing 'stardate' to /usr/local/bin..."
mv stardate /usr/local/bin/stardate
chmod +x /usr/local/bin/stardate

echo "Cleaning up..."
rm -f "${ASSET}"

echo "Installation complete. You can now run 'stardate' from the command line."

#!/bin/bash
set -e

# Ensure the script is run with root privileges.
if [ "$EUID" -ne 0 ]; then
    echo "Please run this installer with sudo or as root."
    exit 1
fi

# Set GitHub repository details.
GITHUB_USER="arcapol"
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

# **Fetch the latest release tag dynamically from GitHub API**
LATEST_TAG=$(curl -s "https://api.github.com/repos/${GITHUB_USER}/${REPO}/releases/latest" | grep -o '"tag_name": "[^"]*' | cut -d'"' -f4)

# Validate if we got a tag
if [ -z "$LATEST_TAG" ]; then
    echo "Error: Unable to fetch latest release tag from GitHub."
    exit 1
fi

echo "Latest release tag found: $LATEST_TAG"

# Define the asset name based on OS and architecture.
ASSET="stardate-${OS}-${ARCH}.tar.gz"
BINARY_NAME="stardate-${OS}-${ARCH}"
DOWNLOAD_URL="https://github.com/${GITHUB_USER}/${REPO}/releases/download/${LATEST_TAG}/${ASSET}"

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

# Verify it's a valid tar.gz before extracting
if ! tar -tzf "${ASSET}" >/dev/null 2>&1; then
    echo "Error: The downloaded file is NOT a valid tar.gz archive!"
    echo "File type detected: $(file ${ASSET})"
    exit 1
fi

echo "Extracting the archive..."
tar -xzvf "${ASSET}"

# Check that the extracted binary exists. The archive should contain a binary named 'stardate'.
if [ ! -f "${BINARY_NAME}" ]; then
    echo "Error: Extracted binary 'stardate' not found."
    exit 1
fi

echo "Renaming '${BINARY_NAME}' to 'stardate'..."
mv "${BINARY_NAME}" "stardate"

echo "Installing 'stardate' to /usr/local/bin..."
mv stardate /usr/local/bin/stardate
chmod +x /usr/local/bin/stardate

echo "Cleaning up..."
rm -f "${ASSET}"

echo "Installation complete. You can now run 'stardate' from the command line."

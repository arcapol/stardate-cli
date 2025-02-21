#!/bin/bash
set -e

# Ensure the script is run with root privileges.
if [ "$EUID" -ne 0 ]; then
    echo "Please run this uninstaller with sudo or as root."
    exit 1
fi

INSTALL_PATH="/usr/local/bin/stardate"

if [ -f "$INSTALL_PATH" ]; then
    echo "Removing 'stardate' from ${INSTALL_PATH}..."
    rm -f "$INSTALL_PATH"
    echo "Uninstallation complete."
else
    echo "'stardate' is not installed at ${INSTALL_PATH}."
fi

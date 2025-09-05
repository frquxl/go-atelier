#!/bin/bash

# Atelier CLI Installation Script
# This script installs the atelier-cli globally for testing purposes

set -e

echo "🚀 Installing Atelier CLI..."

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Error: Go is not installed. Please install Go first."
    echo "   Visit: https://golang.org/dl/"
    exit 1
fi

# Build and install
echo "📦 Building and installing atelier-cli..."
go install -ldflags "-X 'atelier-cli/cmd.version=$(git describe --tags --abbrev=0 2>/dev/null || echo dev)'" .

# Check if installation was successful
if command -v atelier-cli &> /dev/null; then
    echo "✅ Installation successful!"
    echo "   Binary location: $(which atelier-cli)"
    echo ""
    echo "🎨 Try it out:"
    echo "   atelier-cli --help"
    echo "   atelier-cli init myproject"
else
    echo "❌ Installation may have failed."
    echo "   Check your GOPATH/bin or GOBIN environment variables."
    echo "   You might need to add \$HOME/go/bin to your PATH."
    echo ""
    echo "   To fix PATH issue:"
    echo "   echo 'export PATH=\$HOME/go/bin:\$PATH' >> ~/.bashrc"
    echo "   source ~/.bashrc"
fi

echo ""
echo "📚 For more information, see README.md"
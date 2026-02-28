#!/bin/bash

# Build script for cross-platform binaries

echo "🔨 Building CLI Todo App for all platforms..."
echo ""

# Create builds directory
mkdir -p builds

# Build for macOS Intel
echo "📦 Building for macOS (Intel)..."
GOOS=darwin GOARCH=amd64 go build -o builds/todo-mac-intel main.go

# Build for macOS M1/M2
echo "📦 Building for macOS (Apple Silicon)..."
GOOS=darwin GOARCH=arm64 go build -o builds/todo-mac-m1 main.go

# Build for Linux
echo "📦 Building for Linux..."
GOOS=linux GOARCH=amd64 go build -o builds/todo-linux main.go

# Build for Windows
echo "📦 Building for Windows..."
GOOS=windows GOARCH=amd64 go build -o builds/todo-windows.exe main.go

echo ""
echo "✅ Build complete! Binaries are in ./builds/"
echo ""
ls -lh builds/
echo ""
echo "📤 Ready to upload to GitHub Releases!"

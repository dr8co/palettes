# Contributing to Terminal Color Palette Visualizer

Thank you for your interest in contributing to this project! This document provides guidelines and instructions for
contributing.

## Adding a New Color Scheme

1. Fork the repository
2. Create a new file in the `palette` directory (e.g., `your_scheme.go`)
3. Implement the `ColorScheme` interface:

   ```go
   type ColorScheme interface {
       Name() string    // Returns the name of the color scheme
       Show()           // Displays the color scheme
       Colors() []Color // Returns the color definitions
   }
   ```

4. Register your scheme in `palette/registry.go`
5. Add tests if applicable
6. Submit a pull request

## Code Style Guidelines

- Follow standard Go code formatting (`go fmt`)
- Add comments for exported functions and types
- Use meaningful variable and function names
- Include appropriate error handling
- Ensure your code passes `golint` and `go vet`

## Pull Request Process

1. Create your feature branch (`git checkout -b feature/amazing-feature`)
2. Commit your changes (`git commit -m 'Add some amazing feature'`)
3. Push to your fork (`git push origin feature/amazing-feature`)
4. Open a Pull Request with:
    - A clear title and description
    - Any relevant issue numbers
    - Screenshots if applicable (for new color schemes)

## Development Setup

1. Clone the repository
2. Install dependencies: `go mod download`
3. Make your changes
4. Test your changes: `go test ./...`

## Bug Reports and Feature Requests

- Use the GitHub Issues section
- Include as much detail as possible
- Add labels to categorize the issue
- For bugs, include:
  - Expected behavior
  - Actual behavior
  - Steps to reproduce
  - Your environment details

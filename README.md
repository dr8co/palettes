# 🎨 Terminal Color Palette Visualizer

![Made with Go](https://img.shields.io/badge/Made%20with-Go-00ADD8?logo=go)
![Platform](https://img.shields.io/badge/platform-linux%20%7C%20macOS%20%7C%20Windows-blue)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/dr8co/palettes?logo=go)
[![Go Report Card](https://goreportcard.com/badge/github.com/dr8co/palettes)](https://goreportcard.com/report/github.com/dr8co/palettes)
![License](https://img.shields.io/github/license/dr8co/palettes?color=blue)

A command-line tool written in Go that visualizes popular color palettes and themes in your terminal.

This tool helps developers and designers preview and compare how different color schemes look
directly in their terminal environment. ✨

## ⭐ Features

- 🌈 Displays multiple popular color palettes including:
  - 🍦 Catppuccin
  - 🧛 Dracula
  - 🎨 Gruvbox
  - 🎯 Monokai
  - ❄️ Nord
  - 🌹 Rose Pine
  - ☀️ Solarized
  - 🌆 Tokyo Night
  - 👾 Eldritch
  - 🦋 Everblush
- 🔄 Shows color variations and theme variants where available
- 💻 Easy-to-use command-line interface
- 📋 Supports listing all available palettes
- 🔍 Filter palettes by name or family

## 📥 Installation

Requires Go 1.27 or later.

```bash
go install github.com/dr8co/palettes@latest
```

Alternatively, you can download the latest precompiled binary from
the [Releases](https://github.com/dr8co/palettes/releases) page.

## 🚀 Usage

```bash
palettes [options]
```

### 🛠️ Options

- `-show string`: Show specific palette or palette family (e.g., 'catppuccin', 'dark', 'mocha')
- `-list`: List all available palettes
- `-help`: Show help information

### 📖 Examples

```bash
palettes                              # Show all palettes
palettes -show catppuccin             # Show all Catppuccin variants
palettes -show dark                   # Show all dark theme palettes
palettes -show mocha                  # Show Catppuccin Mocha variant
palettes -show "Catppuccin Mocha"     # Show exact palette name
palettes -list                        # List all available palettes
```

## 🎭 Supported Color Schemes

The tool includes several popular color schemes used in terminal emulators, code editors, and other development tools:

- 🍦 **Catppuccin**: A soothing pastel theme for the high-spirited
- 🧛 **Dracula**: A dark theme for many editors, shells, and more
- 🎨 **Gruvbox**: Retro groove color scheme
- 🎯 **Monokai**: A popular color scheme known from Sublime Text
- ❄️ **Nord**: An arctic, north-bluish color palette
- 🌹 **Rose Pine**: All natural pine, faux fur and a bit of soho vibes
- ☀️ **Solarized**: Precision colors for machines and people
- 🌆 **Tokyo Night**: A clean, dark theme that celebrates the lights of Downtown Tokyo
- 👾 **Eldritch**: A mysterious and otherworldly color scheme
- 🦋 **Everblush**: An aesthetically pleasing color scheme with beautiful syntax highlighting and colors.

## 🤝 Contributing

Contributions are welcome! Please check our [Contributing Guidelines](./CONTRIBUTING.md) for detailed information about:

- ✨ Adding new color schemes
- 📝 Code style guidelines
- 🔄 Pull request process
- 🛠️ Development setup
- 🐛 Bug reports and feature requests

## 📄 License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.

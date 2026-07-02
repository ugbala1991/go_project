# ✦ Go Base Converter ✦

A simple CLI tool that converts numbers between bases (hex, bin, dec).  
It runs interactively in the terminal until you type `quit`.

---

## Features
- Input bases supported: **hex**, **bin**, **dec**
- For **hex/bin input** → outputs decimal
- For **dec input** → outputs both binary and hex
- Hex output is always uppercase
- Negative decimal numbers supported
- Graceful error handling (invalid input, wrong usage, unknown commands)

---

## Project Structure

baseconverter/
├── main.go        # Entry point
├── convert.go     # Conversion logic and helpers
├── go.mod         # Go module file
└── README.md      # Documentation

## How to Run

# 1. Clone or copy files into a folder:
   ```bash
   git clone <your-repo-url>
   cd baseconverter

# 2.Initialize Go module:
go mod init baseconverter

# Run the program:
go run base_converter
or 
go run .

# Example Interaction
✦ Welcome to the Base Converter ✦
Type 'help' to see available commands.

> convert 1E hex
✦ Decimal: 30

> convert 10 bin
✦ Decimal: 2

> convert 255 dec
✦ Binary:  11111111
✦ Hex:     FF

> convert -15 dec
✦ Binary:  -1111
✦ Hex:     -F

> convert 1G hex
Error: invalid hex number: 1G

> quit
Goodbye! ✦


# Requirements

    Go 1.18+ installed

    Runs in any terminal (Linux, macOS, Windows)

# License

Open-source and free to use.

This solution is modular, uses Go’s standard library (`strconv` for parsing/formatting), validates inputs, and ensures nothing crashes.  

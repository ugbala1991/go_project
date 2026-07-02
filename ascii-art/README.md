## ASCII-ART: 
Is a classic 01Edu challenge that tests your ability to handle file I/O, string manipulation, and nested loops in Go. The core logic involves mapping ASCII characters to their specific 8-line representation found in the banner files.

## Project Structure
.
├── main.go          # Entry point and argument handling
├── processor.go     # Logic for parsing banners and printing
├── standard.txt     # Banner file (required in same directory)
└── processor_test.go # Recommended unit testing

## Key Logic Explained:

# The Offset Formula: 
In the provided banner files (standard, shadow, thinkertoy), the space character (ASCII 32) is the first character. Each character block consists of 9 lines: one leading newline followed by 8 lines of ASCII art.

To find the first line of any character: Index = (ASCII - 32) x 9 + 1.

# Vertical Printing: 
You cannot print a character fully and then move to the next because the terminal prints horizontally. You must:
- Take the first line of every character in your word.
- Print that whole line.
- Move to the second line of every character, and so on, for all 8 lines.

# Handling Newlines: 
The input string contains "\n". We split the input by these newlines. If a segment is empty (meaning two newlines were together), we just print a blank line.

# Running the Project:
Ensure standard.txt || shadow.txt || thinkertoy.txt is in your folder.

# Run the command:

- go run . "Hello" | cat -e
- go run . "Hello\nThere"
- go run . "ASCII-ART" thinkertoy
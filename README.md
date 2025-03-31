# ASCII Art Generator (Go)

Converts images (JPEG/PNG) into ASCII art with adjustable width and console/file output.
![vg1](https://github.com/user-attachments/assets/46a2bf29-da9d-42ae-bb9c-809e21289f38)
![vgASCII](https://github.com/user-attachments/assets/92cfabf0-001e-461d-9114-aec2dd26d166)

## 🚀 Features
- Supported formats: JPEG, PNG
- Adjustable ASCII art width (in characters)
- Output to file or terminal
- Custom character gradients (e.g., `@`, `#`, ` `)
- Preserves image aspect ratio

## 🛠 Usage
go run main.go -i input.jpg -w 80 -o output.txt

Flags:
    -i — input image path (required)
    -w — output width (default: 100 characters)
    -o — output file (if omitted, prints to terminal)

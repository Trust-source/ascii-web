# ASCII Art Generator

A terminal-styled ASCII art generator built with Go. Type any text, pick a font, and get instant ASCII art rendered in the browser.

![Go](https://img.shields.io/badge/Go-1.21-00ADD8?style=flat&logo=go)
![Tailwind CSS](https://img.shields.io/badge/Tailwind-CDN-38BDF8?style=flat&logo=tailwindcss)

## Preview

```
 _   _      _ _       
| | | | ___| | | ___  
| |_| |/ _ \ | |/ _ \ 
|  _  |  __/ | | (_) |
|_| |_|\___|_|_|\___/ 
```

## Features

- 10 FIGlet fonts to choose from
- Clean terminal aesthetic
- No page reloads — results appear instantly
- Press `Enter` or click Generate

## Project Structure

```
ascii-gen/
├── main.go              # Go server
├── go.mod               # Go module file
└── templates/
    └── index.html       # Frontend UI
```

## Getting Started

**1. Clone the repo**
```bash
git clone https://github.com/your-username/ascii-gen.git
cd ascii-gen
```

**2. Install dependencies**
```bash
go mod tidy
```

**3. Run the server**
```bash
go run main.go
```

**4. Open your browser**
```
http://localhost:8080
```

## How It Works

1. User types text and picks a font in the browser
2. The frontend sends a request to `/generate?text=Hello&font=graffiti`
3. The Go server uses the [go-figure](https://github.com/common-nighthawk/go-figure) library to generate the ASCII art
4. The result is sent back as plain text and displayed on the page

## Dependencies

| Package | Purpose |
|--------|---------|
| [go-figure](https://github.com/common-nighthawk/go-figure) | FIGlet ASCII art generation |
| [Tailwind CSS](https://tailwindcss.com) | Styling via CDN |
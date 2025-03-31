package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"strings"
)

var (
	//charGradient = []string{"@", "$", "8", "W", "9", "H", "4", "Z", "l", "(", "r", "/", "!", ":", ".", " "}
	charGradient = []string{" ", ".", ":", "!", "/", "r", "(", "l", "Z", "4", "H", "9", "W", "8", "$", "@"}
)

func main() {
	// Parse command-line arguments
	imgPath := flag.String("i", "", "Input image path")
	outputWidth := flag.Int("w", 100, "Output width in characters")
	outputFile := flag.String("o", "", "Output file (default: stdout)")
	flag.Parse()

	if *imgPath == "" {
		fmt.Println("Please specify an image file using -i")
		os.Exit(1)
	}

	// Open image file
	file, err := os.Open(*imgPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Decode image
	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	// Resize image
	resized := resizeImage(img, *outputWidth)

	// Convert to ASCII
	asciiArt := imageToASCII(resized)

	// Output result
	if *outputFile != "" {
		err = os.WriteFile(*outputFile, []byte(asciiArt), 0644)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Print(asciiArt)
	}
}

func resizeImage(img image.Image, newWidth int) image.Image {
	bounds := img.Bounds()
	originalWidth := bounds.Dx()
	originalHeight := bounds.Dy()

	// Maintain aspect ratio with character height compensation
	aspectRatio := float64(originalHeight) / float64(originalWidth)
	newHeight := int(float64(newWidth) * aspectRatio * 0.5)

	// Create new image
	resized := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))

	// Scale factors
	xScale := float64(originalWidth) / float64(newWidth)
	yScale := float64(originalHeight) / float64(newHeight)

	for y := 0; y < newHeight; y++ {
		for x := 0; x < newWidth; x++ {
			// Nearest-neighbor sampling
			srcX := int(float64(x) * xScale)
			srcY := int(float64(y) * yScale)
			resized.Set(x, y, img.At(srcX, srcY))
		}
	}
	return resized
}

func imageToASCII(img image.Image) string {
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	var sb strings.Builder

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Convert pixel to grayscale
			r, g, b, _ := img.At(x, y).RGBA()
			luminance := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
			pixelValue := luminance / 256 // Scale to 0-255

			// Map to ASCII character
			charIndex := int((pixelValue * float64(len(charGradient)-1)) / 255)
			if charIndex >= len(charGradient) {
				charIndex = len(charGradient) - 1
			}
			sb.WriteString(charGradient[charIndex])
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func cropImage(imagePath string, topOrBottom string, pixels int) {
	file, err := os.Open(imagePath)
	if err != nil {
		fmt.Println("Error opening image:", err)
		return
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Error decoding image:", err)
		return
	}

	bounds := img.Bounds()
	width := bounds.Max.X
	height := bounds.Max.Y

	var croppedImg image.Image

	if topOrBottom == "top" {
		croppedImg = img.(interface {
			SubImage(r image.Rectangle) image.Image
		}).SubImage(image.Rect(0, 0, width, pixels))
	} else if topOrBottom == "bottom" {
		croppedImg = img.(interface {
			SubImage(r image.Rectangle) image.Image
		}).SubImage(image.Rect(0, height-pixels, width, height))
	} else {
		fmt.Println("Invalid option for topOrBottom. Choose either 'top' or 'bottom'.")
		return
	}

	outputDir := "output"
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		os.Mkdir(outputDir, os.ModePerm)
	}

	originalFileName := strings.TrimSuffix(filepath.Base(imagePath), filepath.Ext(imagePath))
	ext := filepath.Ext(imagePath)
	outputPath := filepath.Join(outputDir, fmt.Sprintf("%s_cropped_%s%s", originalFileName, topOrBottom, ext))
	outputFile, err := os.Create(outputPath)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	switch ext {
	case ".jpg", ".jpeg":
		if err := jpeg.Encode(outputFile, croppedImg, nil); err != nil {
			fmt.Println("Error encoding cropped image:", err)
			return
		}
	case ".png":
		if err := png.Encode(outputFile, croppedImg); err != nil {
			fmt.Println("Error encoding cropped image:", err)
			return
		}
	}

	fmt.Printf("Cropped image saved as %s\n", outputPath)
}

func main() {
	args := os.Args
	if len(args) != 4 {
    fmt.Println("Usage: go run main.go <image_path> <top_or_bottom> <pixels>")
    return
}

	imagePath := args[1]
	topOrBottom := args[2]
	pixels, err := strconv.Atoi(args[3])
	if err != nil {
		fmt.Println("Error converting pixels to integer:", err)
		return
	}

	cropImage(imagePath, topOrBottom, pixels)
}

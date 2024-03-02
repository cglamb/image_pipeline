package imageprocessing

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"

	"github.com/disintegration/imaging"
	"github.com/nfnt/resize"
)

func ReadImage(path string) (image.Image, error) { //modified output to pass error
	inputFile, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("Failure reading imagee.  Cannot find %w", err) //replaced panic with error reporting
	}
	defer inputFile.Close()

	// Decode the image
	img, _, err := image.Decode(inputFile)
	if err != nil {
		return nil, fmt.Errorf("cannot open/decode file: %w", err) //replaced panic with error reporting
	}
	return img, nil
}

func WriteImage(path string, img image.Image) error { //modified output to pass error
	outputFile, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("Failure to save Image: %w", err) //replaced panic with error reporting
	}
	defer outputFile.Close()

	// Encode the image to the new file
	err = jpeg.Encode(outputFile, img, nil)
	if err != nil {
		return fmt.Errorf("Failure to save Image: %w", err) //replaced panic with error reporting
	}
	return nil
}

func Grayscale(img image.Image) image.Image {
	// Create a new grayscale image
	bounds := img.Bounds()
	grayImg := image.NewGray(bounds)

	// Convert each pixel to grayscale
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			originalPixel := img.At(x, y)
			grayPixel := color.GrayModel.Convert(originalPixel)
			grayImg.Set(x, y, grayPixel)
		}
	}
	return grayImg
}

func Resize(img image.Image) image.Image {
	newWidth := uint(500)
	newHeight := uint(500)
	resizedImg := resize.Resize(newWidth, newHeight, img, resize.Lanczos3)
	return resizedImg
}

// rotate image
func RotateImage(img image.Image, angle float64) image.Image {
	return imaging.Rotate(img, angle, color.NRGBA{0, 0, 0, 0})
}

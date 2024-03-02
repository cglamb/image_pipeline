package imageprocessing

import (
	"image"
	"image/color"
	"os"
	"path/filepath"
	"testing"
)

// generates image of given width and height
// also allows for custom color setup in upper left and lower right quadrants
// reference : https://yourbasic.org/golang/create-image/
func GenerateImage(width, height int, color1, color2 color.Color) (image.Image, error) {

	//create the image
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// set the color
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			switch {
			case x < width/2 && y < height/2: // applying color1 to upper left quadrant
				img.Set(x, y, color1)
			case x >= width/2 && y >= height/2: // applying color2 to lower right quadrant
				img.Set(x, y, color2)
			default:
			}
		}
	}

	return img, nil
}

// helper function to check if RGB color code is the same
func colorsMatch(c1, c2 color.Color) bool {
	r1, g1, b1, a1 := c1.RGBA()
	r2, g2, b2, a2 := c2.RGBA()
	// check if a match
	return r1 == r2 && g1 == g2 && b1 == b2 && a1 == a2
}

// testing the resize function
func TestResize(t *testing.T) {

	//gerenate an image of size 1000, 1000 using generateImage
	cyan := color.RGBA{100, 200, 200, 0xff}
	img, _ := GenerateImage(1000, 1000, cyan, color.White)

	//applying the resize function
	resizedImg := Resize(img)

	//check the dimmensions of the resized image
	bounds := resizedImg.Bounds()
	if bounds.Dx() != 500 || bounds.Dy() != 500 {
		t.Errorf("image doesn't match.  not of right size")
	}
}

// skipping grayscale test as i am not understanding how GrayModel.covert is doing the color transformation

func TestWriteImage(t *testing.T) {

	// save an artficical image
	cyan := color.RGBA{100, 200, 200, 0xff}
	img, _ := GenerateImage(500, 500, cyan, color.White)
	WriteImage("is_there.png", img)

	// check that file is there
	if _, err := os.Stat("is_there.png"); os.IsNotExist(err) {
		t.Errorf("File was not created")
	}

}


func TestReadImage(t *testing.T) {

	//creates an image using the generateImage function and saves to a temporary directory
	img, _ := GenerateImage(500, 500, color.RGBA{100, 200, 200, 0xff}, color.White)
	tempDir := t.TempDir()
	imagePath := filepath.Join(tempDir, "test.png")
	WriteImage(imagePath, img)

	// read in without an error
	img, err := ReadImage(imagePath)
	if err != nil {
		t.Fatalf("read error")
	}

	// make sure the read image isnt blank
	if img == nil {
		t.Errorf("read blank image")
	}
}

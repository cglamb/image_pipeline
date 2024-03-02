package main

import (
	"testing"
)

func BenchmarkPipeline(b *testing.B) {

	imagePaths := []string{"images/image1.jpg",
		"images/image2.jpg",
		"images/image3.jpg",
		"images/image4.jpg",
	}

	b.ResetTimer()

	channel1, loadErrors := loadImage(imagePaths)
	channel2 := resize(channel1)
	channel3 := convertToGrayscale(channel2)
	channel4 := rotateImage(channel3)
	writeResults, saveErrors := saveImage(channel4)

	for range writeResults {
	}
	for err := range loadErrors {
		b.Error("Load error:", err)
	}
	for err := range saveErrors {
		b.Error("Save error:", err)
	}

	b.StopTimer()

}

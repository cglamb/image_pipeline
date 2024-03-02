package main

import (
	"fmt"
	imageprocessing "goroutines_pipeline/image_processing"
	"image"
	"strconv"
	"strings"
	"sync"
)

type Job struct {
	InputPath string
	Image     image.Image
	OutPath   string
}

func loadImage(paths []string) (<-chan Job, <-chan error) { //nodified output to include error channel
	out := make(chan Job)
	errChan := make(chan error, 1) //intialize an error channel
	go func() {
		// For each input path create a job and add it to
		// the out channel
		for _, p := range paths {
			var err error
			job := Job{InputPath: p, OutPath: strings.Replace(p, "images/", "images/output/", 1)}
			job.Image, err = imageprocessing.ReadImage(p) //modified to include error
			if err != nil {
				errChan <- fmt.Errorf("Error: %w", err) //send error to error channel
				continue
			}
			out <- job
		}
		close(out)
		close(errChan) //close error channel
	}()
	return out, errChan
}

func resize(input <-chan Job) <-chan Job {
	out := make(chan Job)
	go func() {
		// For each input job, create a new job after resize and add it to
		// the out channel
		for job := range input { // Read from the channel
			job.Image = imageprocessing.Resize(job.Image)
			out <- job
		}
		close(out)
	}()
	return out
}

func convertToGrayscale(input <-chan Job) <-chan Job {
	out := make(chan Job)
	go func() {
		for job := range input { // Read from the channel
			job.Image = imageprocessing.Grayscale(job.Image)
			out <- job
		}
		close(out)
	}()
	return out
}

func saveImage(input <-chan Job) (<-chan bool, <-chan error) { //modified output to include error channel
	out := make(chan bool)
	errChan2 := make(chan error, 1) //intialize an error channel
	go func() {
		for job := range input { // Read from the channel
			err := imageprocessing.WriteImage(job.OutPath, job.Image)
			if err != nil {
				errChan2 <- fmt.Errorf("Error: %w", err) //send error to error channel
				continue
			}
			out <- true
		}
		close(out)
		close(errChan2)
	}()
	return out, errChan2
}

// adding a new function to rotate the images
// function will take an image and rotate it by 90, 180, 270, and 360 degrees
// single input image will produce 4 output images
func rotateImage(input <-chan Job) <-chan Job {
	out := make(chan Job)
	go func() {
		for job := range input {
			rotation_amount := []float64{90, 180, 270, 360} //angles to rotate the image
			for _, rotation_amount := range rotation_amount {
				rImg := imageprocessing.RotateImage(job.Image, rotation_amount) //rotate the image
				newPath := strings.Replace(job.OutPath, ".jpg", "_"+strconv.Itoa(int(rotation_amount))+".jpg", 1)
				out <- Job{InputPath: job.InputPath, Image: rImg, OutPath: newPath}
			}
		}
		close(out)
	}()
	return out
}

func main() {

	// if an error is reported on the last image save, the program is not reporting the error
	// adding a wait group to wait for all error handling goroutines to finish
	var wg sync.WaitGroup

	imagePaths := []string{"images/image1.jpg",
		"images/image2.jpg",
		"images/image3.jpg",
		"images/image4.jpg",
	}

	channel1, loadErrors := loadImage(imagePaths)
	channel2 := resize(channel1)
	channel3 := convertToGrayscale(channel2)
	channel4 := rotateImage(channel3)
	writeResults, saveErrors := saveImage(channel4)

	// pipeline error reporting based on this exmaple: https://medium.com/@TechSavvyScribe/error-handling-in-concurrent-go-programs-8c71d90a1de8
	wg.Add(2) // Add two goroutines to the wait group for error channel (loadErrors and saveErrors)
	go func() {
		for err := range loadErrors {
			fmt.Println(err)
		}
		wg.Done()
	}()

	go func() {
		for err := range saveErrors {
			fmt.Println(err)
		}
		wg.Done()
	}()

	for success := range writeResults {
		if success {
			fmt.Println("Success!")
		} else {
			fmt.Println("Failed!")
		}
	}

	wg.Wait() // Wait for all error handling goroutines to finish

}

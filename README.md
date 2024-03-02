# Image Pipeline in Golang

**Source:** This pipeline is a modified version of Amrit Singh’s original repository [here](https://github.com/code-heim/go_21_goroutines_pipeline). `/Original` is a replica of the original repository and `/Modified` includes this author’s modifications.

**Author:** Charles Lamb

**Contact Info:** charlamb@gmail.com

**Github Project URL:** [https://github.com/cglamb/image_pipeline](https://github.com/cglamb/image_pipeline)

**Github Clone Command:** `git clone https://github.com/cglamb/image_pipeline.git`

## Introduction

This project develops a pipeline for image processing. The pipeline reads in images, resizes, applies grayscale, and rotates the images.

## Modified Pipeline

- Added error checking to `loadImage` and `saveImage` and added error channels.
- Replaced original images with four new images. Images sourced with modification from [here](https://www.kaggle.com/c/leaf-classification). Original images were black and white. Images modified to include a color background.
- Unit testing added in `image_processing_test.go`. Note no unit testing was developed for grayscale functionality as I was unable to figure out the grayscale transformation used by `Graymodel.convert`.
- Pipeline extended to include rotation of images. This may be useful in CNN model training. CNNs are invariant to rotation and a common image processing step is to rotate images in the training dataset to help overcome this weakness.
- Benchmarking added in `main_test.go`.

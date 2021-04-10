package main

import (
    "image"
    "image/color"
    "image/png"
    _ "image/jpeg"
    "os"
)

func main() {
    // Read image file and create Image instance
    file, err := os.Open("test_images/test.png");
    if err != nil {
        panic(err)
    }
    defer file.Close()

    roImage, _, err := image.Decode(file)
    if err != nil {
        panic(err)
    }

    // Convert readed image into in-memory image
    rwImage := image.NewRGBA(image.Rect(0, 0, roImage.Bounds().Max.X, roImage.Bounds().Max.Y))

    for x := roImage.Bounds().Min.X; x < roImage.Bounds().Max.X; x++ {
        for y := roImage.Bounds().Min.Y; y < roImage.Bounds().Max.Y; y++ {
            rwImage.Set(x, y, rwImage.ColorModel().Convert(roImage.At(x, y)))
        }
    }

    // Algorytmic work with image
    grayscaleImage(rwImage)

    // Output image
    newFile, err := os.Create("test_images/test_new.png")
    if err != nil {
        panic(err)
    }
    defer newFile.Close()

    err = png.Encode(newFile, rwImage)
    if err != nil {
        panic(err)
    }
}

func grayscaleImage(img *image.RGBA) {
    for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
        for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
            pixel := img.RGBAAt(x, y)
            maxValue := maxSlice(pixel.R, pixel.G, pixel.B)

            newPixel := color.RGBA{maxValue, maxValue, maxValue, pixel.A}
            img.Set(x, y, newPixel)
        }
    }
}

func maxSlice(numbers... uint8) (res uint8) {
    for i, k := range numbers {
        if i == 0 || k > res {
            res = k
        }
    }

    return
}

package main

import (
    "fmt"
    "os"
    "image"
    "image/png"
    _ "image/jpeg"
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

    var newWidth, newHeight int;

    fmt.Print("Enter width and height of the output: ")
    _, err = fmt.Scanf("%d %d", &newWidth, &newHeight)
    if err != nil {
        panic(err)
    }

    if newWidth <= 0 || newHeight <= 0 {
        panic("Width or height cannot be lower or equal than 0")
    }

    // Resizing image
    resImage := resizeImage(rwImage, newWidth, newHeight)

    // Output image
    newFile, err := os.Create("test_images/test_new.png")
    if err != nil {
        panic(err)
    }
    defer newFile.Close()

    err = png.Encode(newFile, resImage)
    if err != nil {
        panic(err)
    }
}

func resizeImage(img *image.RGBA, newWidth, newHeight int) image.Image {
    // Create result image
    resImage := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))

    // Calculate scale for x and y axis
    xOldAxis, yOldAxis := img.Bounds().Max.X, img.Bounds().Max.Y
    xScale, yScale := float64(newWidth) / float64(xOldAxis), float64(newHeight) / float64(yOldAxis)

    xLoopEnd := map[bool]int{true: img.Bounds().Max.X, false: newWidth}[xScale <= 1.0]
    yLoopEnd := map[bool]int{true: img.Bounds().Max.Y, false: newHeight}[yScale <= 1.0]

    for x := 0; x < xLoopEnd; x++ {
        for y := 0; y < yLoopEnd; y++ {
            xCur := map[bool]int{true: x, false: int(float64(x) / xScale)}[xScale <= 1.0]
            yCur := map[bool]int{true: y, false: int(float64(y) / yScale)}[yScale <= 1.0]
            pixel := img.RGBAAt(xCur, yCur)

            xNew := map[bool]int{true: int(float64(x) * xScale), false: x}[xScale <= 1.0]
            yNew := map[bool]int{true: int(float64(y) * yScale), false: y}[yScale <= 1.0]
            resImage.Set(xNew, yNew, pixel)
        }
    }

    return resImage
}

package main

import (
    "image"
    "image/jpeg"
    "os"
)

func main() {
    infile, err := os.Open(os.Args[1])
    if err != nil {
        // replace this with real error handling
        panic(err.String())
    }
    defer infile.Close()

    // Decode will figure out what type of image is in the file on its own.
    // We just have to be sure all the image packages we want are imported.
    src, _, err := image.Decode(infile)
    if err != nil {
        // replace this with real error handling
        panic(err.String())
    }

    // Create a new grayscale image
    bounds := src.Bounds()
    w, h := bounds.Max.X, bounds.Max.Y
    gray := image.NewGray(w, h)
    for x := 0; x < w; x++ {
        for y := 0; y < h; y++ {
            oldColor := src.At(x, y)
            grayColor := image.GrayColorModel.Convert(oldColor)
            gray.Set(x, y, grayColor)
        }
    }

    // Encode the grayscale image to the output file
    outfile, err := os.Create(os.Args[2])
    if err != nil {
        // replace this with real error handling
        panic(err.String())
    }
    defer outfile.Close()
    jpeg.Encode(outfile, gray)
}
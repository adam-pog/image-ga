package main

import (
    "github.com/llgcode/draw2d/draw2dimg"
    "image"
    "image/color"
)

func main() {
    // Initialize the graphic context on an RGBA image
    dest := image.NewRGBA(image.Rect(0, 0, 1920, 1080))
    gc := draw2dimg.NewGraphicContext(dest)

    // Set some properties
    gc.SetFillColor(color.RGBA{0x44, 0xff, 0x44, 0x89})
    gc.SetLineWidth(1)

    // Draw a closed shape
    gc.MoveTo(500, 500) // should always be called first for a new path
    gc.LineTo(600, 600)
    gc.LineTo(600, 400)
    gc.Close()
    gc.Fill()

    // Save to file
    draw2dimg.SaveToPngFile("/home/npzd/Pictures/hello.png", dest)
}


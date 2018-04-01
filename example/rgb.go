package main

import (
	"image"
	"image/color"
	"log"

	"github.com/BakeRolls/canvas"
)

var colors = []color.RGBA{
	color.RGBA{R: 255, A: 255},
	color.RGBA{G: 255, A: 255},
	color.RGBA{B: 255, A: 255},
}

// Create a new image and fill it line by line (the first third red, the second
// green and the third blue). Close the Window afterwards.
func main() {
	im := image.NewRGBA(image.Rect(0, 0, 480, 360))
	c, err := canvas.New(im, 1, "RGB Canvas")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	b := im.Bounds()
	for y := b.Min.Y; c.Update() && y < b.Max.Y; y++ {
		color := colors[y/(b.Max.Y/len(colors))]
		for x := b.Min.X; x < b.Max.X; x++ {
			im.Set(x, y, color)
		}
		c.Draw()
	}
}

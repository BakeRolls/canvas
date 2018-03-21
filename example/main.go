package main

import (
	"image"
	"image/color"
	"log"
	"time"

	"github.com/BakeRolls/canvas"
)

var colors = []color.RGBA{
	color.RGBA{R: 255, A: 255},
	color.RGBA{G: 255, A: 255},
	color.RGBA{B: 255, A: 255},
}

func main() {
	im := image.NewRGBA(image.Rect(0, 0, 640, 480))
	c, err := canvas.New(im, "RGB Canvas")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	b := im.Bounds()
	for y := b.Min.Y; y < b.Max.Y; y++ {
		color := colors[y/(b.Max.Y/len(colors))]
		for x := b.Min.X; x < b.Max.X; x++ {
			im.Set(x, y, color)
		}
	}
	if err := c.Draw(); err != nil {
		log.Fatal(err)
	}

	time.Sleep(5 * time.Second)
}

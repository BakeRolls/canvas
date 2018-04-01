package main

import (
	"image"
	_ "image/png"
	"log"
	"os"

	"github.com/BakeRolls/canvas"
)

// Open an image and show it twice its original size.
func main() {
	f, err := os.Open("daggett.png")
	if err != nil {
		log.Fatal(err)
	}
	im, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	c, err := canvas.New(im, 2, "Daggett")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	for c.Update() {
		c.Draw()
	}
}

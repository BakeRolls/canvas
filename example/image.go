package main

import (
	"image"
	_ "image/png"
	"log"
	"os"
	"time"

	"github.com/BakeRolls/canvas"
)

// Open an image and show it twice as big.
func main() {
	f, err := os.Open("dagget.png")
	if err != nil {
		log.Fatal(err)
	}
	im, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	c, err := canvas.New(im, 2, "Dagget")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	c.Draw()

	time.Sleep(5 * time.Second)
}

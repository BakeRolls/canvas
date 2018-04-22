package main

import (
	"image"
	"image/draw"
	_ "image/png"
	"log"
	"math"
	"os"
	"sync"
	"time"

	"github.com/BakeRolls/canvas"
)

// Open an image, show it twice its original size and animate it using a sine wave.
func main() {
	f, err := os.Open("daggett.png")
	if err != nil {
		log.Fatal(err)
	}
	src, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	dst := image.NewRGBA(src.Bounds())

	c, err := canvas.New(dst, 2, "Daggett")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	sw := &sinwave{c.Mu, src, dst}

	go sw.draw()
	for c.Update() {
		c.Draw()
	}
}

type sinwave struct {
	mu  *sync.Mutex
	src image.Image
	dst draw.Image
}

func (sw *sinwave) draw() {
	b := sw.src.Bounds()
	for t := range time.Tick(time.Second / 30) {
		ms := float64(t.UnixNano()/int64(time.Millisecond)) / 25
		sw.mu.Lock()
		for y := b.Min.Y; y < b.Max.Y; y++ {
			offX := int(math.Sin((ms+float64(y))/10) * 5)
			for x := b.Min.X; x < b.Max.X; x++ {
				offY := int(math.Sin((ms+float64(x))/10) * 5)
				sw.dst.Set(x+offX, y+offY, sw.src.At(x, y))
			}
		}
		sw.mu.Unlock()
	}
}

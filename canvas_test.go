package canvas

import (
	"image"
	"log"
)

func Example() {
	img := image.NewRGBA(image.Rect(0, 0, 640, 480))

	c, err := New(img, 1, "Canvas Title")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	// Update returns true until the window should be closed.
	for c.Update() {
		// Modify the image based on its Pix slice or Set.
		c.Draw()
	}
}

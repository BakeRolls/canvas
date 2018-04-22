# canvas

Draw an image.Image to a window. You need to have SDL2 and pkg-config to compile your program.

```go
im := image.NewRGBA(image.Rect(0, 0, 640, 480))

c, err := canvas.New(im, 1, "RGB Canvas")
if err != nil {
	log.Fatal(err)
}
defer c.Close()

// Update returns true until the window should be closed.
for c.Update() {
	// Modify the image based on its Pix slice or Set.
	c.Draw()
}
```

![Image](/example/image_screenshot.png) ![Sine Wave](/example/sinewave_screenshot.gif)

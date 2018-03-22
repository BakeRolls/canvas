# canvas

Draw an image.Image to a window. You need to have SDL2 and pkg-config to compile your program.

```go
im := image.NewRGBA(image.Rect(0, 0, 640, 480))

c, err := canvas.New(im, "RGB Canvas")
if err != nil {
	log.Fatal(err)
}
defer c.Close()

// Modify the image based on its Pix slice or Set.

c.Draw()
```

![Dagget Beaver](/example/screenshot.png)

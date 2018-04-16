package canvas

import (
	"image"

	"github.com/pkg/errors"
	"github.com/veandco/go-sdl2/sdl"
)

// Canvas to draw an image on.
type Canvas struct {
	image   image.Image
	scale   int
	window  *sdl.Window
	surface *sdl.Surface
}

// New creates a new canvas.
func New(im image.Image, scale int, title string) (*Canvas, error) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return nil, errors.Wrap(err, "could not initialize sdl")
	}
	w := int32(im.Bounds().Max.X * scale)
	h := int32(im.Bounds().Max.Y * scale)
	window, err := sdl.CreateWindow(title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, w, h, sdl.WINDOW_SHOWN)
	if err != nil {
		return nil, errors.Wrap(err, "could not create window")
	}
	surface, err := window.GetSurface()
	if err != nil {
		return nil, errors.Wrap(err, "could not get window surface")
	}
	return &Canvas{im, scale, window, surface}, nil
}

// Update determines if the window shold get closed.
func (c Canvas) Update() bool {
	_, ok := sdl.PollEvent().(*sdl.QuitEvent)
	return !ok
}

// Draw clears and redraws the image.
func (c Canvas) Draw() error {
	b := c.image.Bounds()
	rect := &sdl.Rect{W: int32(c.scale), H: int32(c.scale)}
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			rect.X, rect.Y = int32(x*c.scale), int32(y*c.scale)
			r, g, b, a := c.image.At(x, y).RGBA()
			r, g, b, a = r/257, g/257, b/257, a/257
			if err := c.surface.FillRect(rect, a<<24|r<<16|g<<8|b<<0); err != nil {
				return errors.Wrapf(err, "could not draw pixel %dx%d", x, y)
			}
		}
	}
	if err := c.window.UpdateSurface(); err != nil {
		return errors.Wrap(err, "could not update window surface")
	}
	return nil
}

// Close destroys the window.
func (c Canvas) Close() {
	c.window.Destroy()
	sdl.Quit()
}

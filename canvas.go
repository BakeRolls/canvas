package canvas

import (
	"image"

	"github.com/pkg/errors"
	"github.com/veandco/go-sdl2/sdl"
)

// Canvas to draw an image on.
type Canvas struct {
	title   string
	image   image.Image
	window  *sdl.Window
	surface *sdl.Surface
}

// New creates a new canvas.
func New(im image.Image, t string) (*Canvas, error) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return nil, errors.Wrap(err, "could not initialize sdl")
	}
	b := im.Bounds()
	w, err := sdl.CreateWindow(t, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(b.Max.X), int32(b.Max.Y), sdl.WINDOW_SHOWN)
	if err != nil {
		return nil, errors.Wrap(err, "could not create window")
	}
	s, err := w.GetSurface()
	if err != nil {
		return nil, errors.Wrap(err, "could not get window surface")
	}
	return &Canvas{t, im, w, s}, nil
}

// Draw clears and redraws the image.
func (c Canvas) Draw() error {
	b := c.image.Bounds()
	rect := &sdl.Rect{W: 1, H: 1}
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			rect.X, rect.Y = int32(x), int32(y)
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

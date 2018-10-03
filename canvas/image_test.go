package canvas

import "image"
import "testing"

import "github.com/stretchr/testify/assert"

func TestImageFromImage(t *testing.T) {
	source := image.Rect(2, 2, 4, 4)
	dest := NewImageFromImage(source)

	assert.Equal(t, dest.PixelColor(0, 0, 6, 6).A, uint8(0x00))
	assert.Equal(t, dest.PixelColor(2, 2, 6, 6).A, uint8(0xff))
	assert.Equal(t, dest.PixelColor(4, 4, 6, 6).A, uint8(0x00))
}

func TestImage_AlphaDefault(t *testing.T) {
	img := &Image{}

	assert.Equal(t, 1.0, img.Alpha())
}

func TestImage_TranslucencyDefault(t *testing.T) {
	img := &Image{}

	assert.Equal(t, 0.0, img.Translucency)
}

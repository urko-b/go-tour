package images

import (
	"golang.org/x/tour/pic"
	"image"
	"testing"
)

func TestImage(t *testing.T) {
	rect := image.Rect(0, 0, 255, 255)
	myImage := image.NewRGBA(rect)
	m := Image{myImage}
	pic.ShowImage(m)
}

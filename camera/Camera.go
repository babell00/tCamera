package camera

import (
	"image"
)

type Camera struct {
	Id    string
	Name  string
	Url   string
	Path  string
	Image image.Image
}

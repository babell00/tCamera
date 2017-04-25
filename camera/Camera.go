package camera

import (
	"image"
)

type Camera struct {
	Name   string `json:"name"`
	Id     string `json:"id"`
	Url    string `json:"url"`
	Path   string `json:"path"`
	Image  image.Image `json:"-"`
	Online bool `json:"online"`
}

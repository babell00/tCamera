package camera

import (
	"image"
)

type Camera struct {
	Name   string `json:"name"`
	Id     string `json:"id"`
	MJpegUrl    string `json:"mjpeg_url"`
	PublicUrl string `json:"public_url"`
	UrlPath   string `json:"url_path"`
	Image  image.Image `json:"-"`
	Online bool `json:"online"`
}

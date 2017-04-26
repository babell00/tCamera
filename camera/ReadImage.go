package camera

import (
	"image"
	"github.com/babell00/toc_camera/mjpeg"
)

func ReadJpeg(url string) (image.Image, error) {
	data, err := mjpeg.NewDecoderFromURL(url)
	if err != nil {
		return nil, err
	}

	img, err := data.Decode()
	if err != nil {
		return nil, err
	}
	return img, nil
}

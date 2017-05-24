package camera

import (
	"image"
	"github.com/babell00/tCamera/mjpeg"
)

func ReadJpeg(url string) (image.Image, error) {
	decoder, err := mjpeg.NewDecoderFromURL(url)
	if err != nil {
		return nil, err
	}

	img, err := decoder.Decode()
	if err != nil {
		return nil, err
	}
	return img, nil
}

func Decoder(url string) (*mjpeg.Decoder, error)  {
	decoder, err := mjpeg.NewDecoderFromURL(url)
	if err != nil {
		return nil, err
	}
	return decoder, err

}

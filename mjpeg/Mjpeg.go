package mjpeg

import (
	"image"
	"image/jpeg"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"strings"
	"sync"
)

type Decoder struct {
	reader *multipart.Reader
	mux sync.Mutex
}

func NewDecoder(reader io.Reader, boundary string) *Decoder {
	decoder := new(Decoder)
	decoder.reader = multipart.NewReader(reader, boundary)
	return decoder
}

func NewDecoderFromResponse(res *http.Response) (*Decoder, error) {
	_, param, err := mime.ParseMediaType(res.Header.Get("Content-Type"))
	if err != nil {
		return nil, err
	}
	return NewDecoder(res.Body, strings.Trim(param["boundary"], "-")), nil
}

func NewDecoderFromURL(url string) (*Decoder, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return NewDecoderFromResponse(res)
}

func (d *Decoder) Decode() (image.Image, error) {
	part, err := d.reader.NextPart()
	if err != nil {
		return nil, err
	}
	return jpeg.Decode(part)
}
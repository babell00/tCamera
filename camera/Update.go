package camera

import (
	"image"
	"github.com/babell00/toc_camera/mjpeg"
	"time"
	"log"
)

func readJpeg(url string) (image.Image, error) {
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

func (service *CameraService) SetUpdateFunction(updateInterval int) {
	ticker := time.NewTicker(time.Second * time.Duration(updateInterval))
	go func() {
		for range ticker.C {
			for _, camera := range service.GetAll() {
				updateCamera(camera, service)
			}
		}
	}()
}

func updateCamera(camera Camera, service *CameraService) {
	img, err := readJpeg(camera.Url)
	if err != nil {
		log.Printf("Cannot read camera stream : %#v", camera)
		return
	}

	camera.Image = img
	service.repository.SaveById(camera.Id, camera)
	log.Printf("Update camera: %v at %v", camera, time.Now())
}

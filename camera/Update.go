package camera

import (
	"image"
	"github.com/babell00/toc_camera/mjpeg"
	"time"
	"log"
	"net/http"
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

func SetUpdateCameraFunction(updateInterval int, service *CameraService) {
	ticker := time.NewTicker(time.Second * time.Duration(updateInterval))
	go func() {
		for range ticker.C {
			cameras := service.GetAll()
			Update(cameras, service)
		}
	}()
}

func Update(cameras []Camera, service *CameraService) {
	for _, camera := range cameras {
		go updateCamera(camera, service)
	}
}

func updateCamera(camera Camera, service *CameraService) {
	img, err := readJpeg(camera.Url)
	if err != nil {
		log.Printf("Cannot read camera stream : %#v", camera)
		return
	}

	camera.Image = img
	service.repository.Save(camera)
	log.Printf("Update camera: %v at %v", camera, time.Now())
}

func SetUpdateStatusFunction(updateInterval int, service *CameraService) {
	ticker := time.NewTicker(time.Second * time.Duration(updateInterval))
	go func() {
		for range ticker.C {
			cameras := service.GetAll()
			UpdateStatus(cameras, service)
		}
	}()
}

func UpdateStatus(cameras []Camera, service *CameraService) {
	for _, camera := range cameras {
		go updateCameraStatus(camera, service)
	}
}

func updateCameraStatus(camera Camera, service *CameraService) {
	timeout := time.Duration(10 * time.Second)
	client := &http.Client{Timeout: timeout, }
	req, _ := http.NewRequest("GET", camera.Url, nil)

	resp, err := client.Do(req)
	if err != nil {
		camera.Online = false
		service.Save(camera)
		log.Printf("Camera is not responding: %v", camera)
		return
	}
	if resp.Status == "200 OK" {
		camera.Online = true
		service.Save(camera)
	}
}

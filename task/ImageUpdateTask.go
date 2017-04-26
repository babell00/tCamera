package task

import (
	"log"
	"time"
	"github.com/babell00/toc_camera/camera"
)

func UpdateImage(service *camera.CameraService) {
	log.Println("Updating camers images")
	cameras := service.GetAll()
	for _, cam := range cameras {
		if cam.UrlPath == "test" {
			go updateCamera(cam, service)
		}
	}
}

func updateCamera(cam camera.Camera, service *camera.CameraService) {
	img, err := camera.ReadJpeg(cam.MJpegUrl)
	if err != nil {
		log.Printf("Cannot read camera stream : %#v", cam)
		return
	}

	cam.Image = img
	service.Save(cam)
	log.Printf("Update camera: %v at %v", cam, time.Now())
}

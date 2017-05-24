package task

import (
	"log"
	"github.com/babell00/tCamera/camera"
	"time"
)

func UpdateImage(service *camera.CameraService) {
	log.Println("Updating camers images")
	cameras := service.GetAll()
	for _, cam := range cameras {
		go updateCamera(cam, service)
	}
}

func updateCamera(cam camera.Camera, service *camera.CameraService) {
	img, err := camera.ReadJpeg(cam.MJpegUrl)
	cam.LastUpdate = time.Now()
	if err != nil {
		log.Printf("Cannot read camera stream : %#v", cam)
		service.Save(cam)
		return
	}

	cam.Image = img
	service.Save(cam)
	log.Printf("Update camera: %v", cam)
}

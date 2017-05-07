package task

import (
	"time"
	"net/http"
	"log"
	"github.com/babell00/toc_camera/camera"
)

func UpdateStatus(service *camera.CameraService) {
	cameras := service.GetAll()
	log.Println("Updating camers status")
	for _, cam := range cameras {
		go updateCameraStatus(cam, service)
	}
}

func updateCameraStatus(cam camera.Camera, service *camera.CameraService) {
	resp, err := request(cam.MJpegUrl)
	if err != nil {
		cam.Online = false
		service.Save(cam)
		log.Printf("Camera is not responding: %v", cam)
		return
	}

	if resp.Status == "200 OK" {
		cam.Online = true
		cam.LastOnline = time.Now()
	} else {
		cam.Online = false
	}
	service.Save(cam)
}

func request(url string) (*http.Response, error) {
	timeout := time.Duration(10 * time.Second)
	client := &http.Client{Timeout: timeout, }
	req, _ := http.NewRequest("GET", url, nil)

	return client.Do(req)
}

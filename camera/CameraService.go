package camera

import (
	"image"
	"github.com/babell00/toc_camera/mjpeg"
	"time"
	"log"
	"fmt"
)

type CameraService struct {
	repository *cameraRepository
}

func InitService(cameras []Camera) *CameraService {
	cameraModle := InitRepository(cameras)
	cameraService := CameraService{cameraModle}
	return &cameraService
}

func (service *CameraService) GetById(id string) Camera {
	return service.repository.FindById(id)
}

func (service *CameraService) GetByPath(path string) Camera {
	return service.repository.FindCameraByPath(path)
}

func (service *CameraService) GetAll() []Camera {
	return service.repository.FindAll()
}

func (service *CameraService) Save(camera Camera) {
	service.repository.Save(camera)
}

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
		for t := range ticker.C {
			for _, camera := range service.GetAll() {
				img, err := readJpeg(camera.Url)
				if err != nil {
					log.Printf("Cannot read camera stream : %#v", camera)
					continue
				}
				c := service.repository.FindById(camera.Id)
				c.Image = img
				service.repository.Save(c)
			}
			fmt.Println("Tick at", t)
		}
	}()

}

package camera

import (
	"log"
)

type cameraRepository struct {
	data map[string]Camera
}

func InitRepository(cameras []Camera) *cameraRepository {
	log.Println("Init Camera Model")
	cameraMap := make(map[string]Camera)
	cameraModle := cameraRepository{cameraMap}
	for _, camera := range cameras {
		cameraModle.data[camera.Id] = camera
	}
	return &cameraModle

}

func (repository *cameraRepository) FindAll() []Camera {
	cameraList := make([]Camera, 0, len(repository.data))
	for _, value := range repository.data {
		cameraList = append(cameraList, value)
	}
	return cameraList
}

func (repository *cameraRepository) FindById(id string) Camera {
	return repository.data[id]
}

func (repository *cameraRepository) FindCameraByPath(path string) Camera {
	var camera Camera
	for _, c := range repository.data {
		if c.Path == path {
			camera = c
		}
	}

	return camera
}


func (repository *cameraRepository) Save(camera Camera){
	repository.data[camera.Id] = camera
}

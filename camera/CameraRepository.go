package camera

import (
	"log"
	"github.com/twinj/uuid"
	"sync"
)

type cameraRepository struct {
	data map[string]Camera
	mux  *sync.Mutex
}

func InitRepository(cameras []Camera) *cameraRepository {
	log.Println("Init Camera Repository")
	mutex := &sync.Mutex{}
	data := make(map[string]Camera)
	cameraRepository := cameraRepository{data, mutex}
	cameraRepository.mux.Lock()
	for _, camera := range cameras {
		camera.Id = uuid.NewV4().String()
		cameraRepository.data[camera.Id] = camera
	}
	cameraRepository.mux.Unlock()
	return &cameraRepository

}

func (repository *cameraRepository) FindAll() []Camera {
	cameraList := make([]Camera, 0, len(repository.data))
	repository.mux.Lock()
	for _, value := range repository.data {

		cameraList = append(cameraList, value)
	}
	repository.mux.Unlock()
	return cameraList
}

func (repository *cameraRepository) FindById(id string) Camera {
	return repository.data[id]
}

func (repository *cameraRepository) FindCameraByPath(path string) Camera {
	var camera Camera
	for _, c := range repository.data {
		if c.UrlPath == path {
			camera = c
		}
	}

	return camera
}

func (repository *cameraRepository) Save(camera Camera) {
	repository.mux.Lock()
	if camera.Id == "" {
		camera.Id = uuid.NewV4().String()
	}
	repository.data[camera.Id] = camera
	repository.mux.Unlock()
}

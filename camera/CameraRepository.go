package camera

import (
	"github.com/twinj/uuid"
	"sync"
	"log"
)

var (
	repository *cameraRepository
	repositoryOnce       sync.Once
)

type cameraRepository struct {
	items map[string]Camera
	mux   sync.RWMutex
}

func Repository() *cameraRepository {
	log.Println("Init Camera Repository")
	repositoryOnce.Do(func() {
		repository = &cameraRepository{items: make(map[string]Camera)}
	})
	return repository
}

func (repository *cameraRepository) AddItems(cameras []Camera) {
	repository.mux.Lock()
	defer  repository.mux.Unlock()

	for _, camera := range cameras {
		camera.Id = uuid.NewV4().String()
		repository.items[camera.Id] = camera
	}

}


func (repository *cameraRepository) FindAll() []Camera {
	repository.mux.RLock()
	defer repository.mux.RUnlock()

	cameraList := make([]Camera, 0, len(repository.items))
	for _, value := range repository.items {

		cameraList = append(cameraList, value)
	}

	return cameraList
}

func (repository *cameraRepository) FindById(id string) Camera {
	repository.mux.RLock()
	defer repository.mux.RUnlock()

	return repository.items[id]
}

func (repository *cameraRepository) FindCameraByPath(path string) Camera {
	repository.mux.RLock()
	defer repository.mux.RUnlock()

	for _, c := range repository.items {
		if c.UrlPath == path {
			return c
		}
	}

	return Camera{}
}

func (repository *cameraRepository) Save(camera Camera) {
	repository.mux.Lock()
	defer repository.mux.Unlock()

	if camera.Id == "" {
		camera.Id = uuid.NewV4().String()
	}
	repository.items[camera.Id] = camera
}

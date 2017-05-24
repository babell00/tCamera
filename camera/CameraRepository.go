package camera

import (
	"log"
	"sync"

	"github.com/twinj/uuid"
)

var (
	repository     *cameraRepository
	repositoryOnce sync.Once
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

func (repository *cameraRepository) FindByPath(path string) Camera {
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

func (repository *cameraRepository) SaveItems(cameras []Camera) {
	for _, camera := range cameras {
		repository.Save(camera)
	}

}

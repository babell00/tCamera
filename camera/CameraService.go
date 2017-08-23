package camera

import "sync"

var (
	service     *CameraService
	serviceOnce sync.Once
)

type CameraService struct {
	repository *cameraRepository
}

func Service() *CameraService {
	serviceOnce.Do(func() {
		cameraRepository := Repository()
		service = &CameraService{cameraRepository}
	})
	return service
}

func (service *CameraService) AddCameras(cameras []Camera) {
	service.repository.SaveItems(cameras)
}

func (service *CameraService) GetById(id string) Camera {
	return service.repository.FindById(id)
}

func (service *CameraService) GetByPath(path string) Camera {
	return service.repository.FindByPath(path)
}

func (service *CameraService) GetAll() []Camera {
	return service.repository.FindAll()
}

func (service *CameraService) Save(camera Camera) {
	service.repository.Save(camera)
}

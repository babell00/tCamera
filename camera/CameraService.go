package camera

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

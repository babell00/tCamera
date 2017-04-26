package server

import (
	"net/http"
	"github.com/babell00/toc_camera/camera"
	"sort"
	"encoding/json"
	"image/jpeg"
	"github.com/gorilla/mux"
)

var _service *camera.CameraService

func InitHandlers(service *camera.CameraService) {
	_service = service

}

func CameraIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	cameras := getService().GetAll()

	sort.Slice(cameras, func(a, b int) bool {
		return cameras[a].Name < cameras[b].Name
	})

	json.NewEncoder(w).Encode(&cameras)
}

func CameraImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "image/jpeg")

	vars := mux.Vars(r)
	path :=vars["camera_path"]

	camera := getService().GetByPath(path)
	if camera.Image == nil {
		return
	}

	err := jpeg.Encode(w, camera.Image, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getService() *camera.CameraService {
	if _service == nil {
		panic("Handlers need to be initialize. InitHandlers(ser *camera.CameraService)")
	}
	return _service
}

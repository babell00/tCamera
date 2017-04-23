package network

import (
	"log"
	"net/http"
	"fmt"
	"github.com/babell00/toc_camera/camera"
	"image/jpeg"
)


func StartServer(portNumber int, service *camera.CameraService) {
	log.Println("Setting up Server")

	serverSetup(portNumber, service)
}

func formatPortNumber(portNumber int) string {
	return fmt.Sprintf(":%v", portNumber)
}

func createEndpointsForCameras(service *camera.CameraService) {
	cameras := service.GetAll()
	for _, camera := range cameras {
		log.Printf("Registering camera: %#v", camera)
		if validateCamera(camera) {
			log.Print("Registration - FAILED")
			continue
		}

		registerHandler(camera.Path, service)
		log.Print("Registration - SUCCESS")
	}

}

func validateCamera(camera camera.Camera) bool {
	return camera.Path == "" || camera.Url == ""
}

func registerHandler(path string, service *camera.CameraService) {
	convertedPath := fmt.Sprintf("/%v", path)
	http.HandleFunc(convertedPath, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/jpeg")

		camera := service.GetByPath(path)
		if camera.Image == nil {
			return
		}
		err := jpeg.Encode(w, camera.Image, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}

func serverSetup(portNumber int, service *camera.CameraService) {
	createEndpointsForCameras(service)

	formattedPortNumber := formatPortNumber(portNumber)
	log.Printf("Server is listening on port: %v", portNumber)
	log.Fatal("Server error: ", http.ListenAndServe(formattedPortNumber, nil))
}

package server

import (
	"log"
	"net/http"
	"fmt"
	"github.com/babell00/toc_camera/camera"
	"github.com/gorilla/mux"
)


func NewServer(portNumber int, service *camera.CameraService) {
	log.Println("Setting up Server")

	router := initServer(service)
	formattedPortNumber := formatPortNumber(portNumber)

	log.Printf("Server is listening on port: %v", portNumber)
	log.Fatal("Server error: ", http.ListenAndServe(formattedPortNumber, router))
}

func formatPortNumber(portNumber int) string {
	return fmt.Sprintf(":%v", portNumber)
}

func initServer(service *camera.CameraService) *mux.Router{
	InitHandlers(service)
	return NewRouter()
}

package main

import (
	"github.com/babell00/toc_camera/configuration"
	"runtime"
	"log"
	"github.com/babell00/toc_camera/camera"
	"github.com/babell00/toc_camera/network"
	"os"
)

func main() {
	setup()
}

func printInfo(config configuration.Config) {
	log.Printf("Cameras refresh interval: %vmins", config.RefreshInterval)
	log.Printf("Added %v cameras", len(config.Cameras))
}

func setup() {

	//setLogger()

	log.Println("Setting up application")

	config := configuration.ReadConfigurationFromYaml()

	log.Printf("Max number of CPUs: %v", config.MaxCpu)
	runtime.GOMAXPROCS(config.MaxCpu)

	cameras := camera.ConvertConfigCameraToCamera(config)

	cameraService := camera.InitService(cameras)

	cameras = cameraService.GetAll()
	camera.Update(cameras, cameraService)
	camera.SetUpdateCameraFunction(config.RefreshInterval, cameraService)

	camera.SetUpdateStatusFunction(10, cameraService)

	printInfo(config)

	network.StartServer(config.Server.Port, cameraService)
}

func setLogger() {
	f, err := os.OpenFile("camera.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Printf("Error opening file: %v", err)
	}
	log.SetOutput(f)
}

package main

import (
	"github.com/babell00/toc_camera/configuration"
	"runtime"
	"log"
	"github.com/babell00/toc_camera/camera"
	"github.com/babell00/toc_camera/network"
)

func main() {
	setup()
}

func printInfo(config configuration.Config) {
	log.Printf("Cameras refresh interval: %vmins", config.RefreshInterval)
	log.Printf("Added %v cameras", len(config.Cameras))
}

func setup() {
	log.Println("Setting up application")

	config := configuration.ReadConfigurationFromYaml()

	log.Printf("Max number of CPUs: %v", config.MaxCpu)
	runtime.GOMAXPROCS(config.MaxCpu)

	cameras := camera.ConvertConfigCameraToCamera(config)

	cameraService := camera.InitService(cameras)

	cameraService.SetUpdateFunction(config.RefreshInterval)

	printInfo(config)

	network.StartServer(config.Server.Port, cameraService)
}

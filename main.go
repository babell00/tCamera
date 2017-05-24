package main

import (
	"log"
	"os"
	"runtime"

	"github.com/babell00/tCamera/camera"
	"github.com/babell00/tCamera/configuration"
	"github.com/babell00/tCamera/server"
	"github.com/babell00/tCamera/task"
	"github.com/jasonlvhit/gocron"
)

func main() {
	setup()
	log.Println("Hello")
}

func printInfo(config configuration.Config) {
	log.Printf("Cameras refresh interval: %vseconds", config.ImageUpdateInterval)
	log.Printf("Added %v cameras", len(config.Cameras))
}

func setup() {

	//setLogger()

	log.Println("Setting up application")

	config := configuration.ReadConfigurationFromYaml()

	log.Printf("Max number of CPUs: %v", config.MaxCpu)
	runtime.GOMAXPROCS(config.MaxCpu)

	cameras := camera.ConvertConfigCameraToCamera(config)

	cameraService := camera.Service()
	cameraService.AddCameras(cameras)

	registerTasks(config, cameraService)

	printInfo(config)

	server.NewServer(config.Server.Port, cameraService)
}

func setLogger() {
	f, err := os.OpenFile("camera.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Printf("Error opening file: %v", err)
	}
	log.SetOutput(f)
}

func registerTasks(config configuration.Config, service *camera.CameraService) {
	task.UpdateImage(service)

	imageUpdate := gocron.NewScheduler()
	imageUpdate.Every(config.ImageUpdateInterval).Seconds().Do(task.UpdateImage, service)
	imageUpdate.Start()

	statusUpdate := gocron.NewScheduler()
	statusUpdate.Every(config.StatusUpdateInterval).Seconds().Do(task.UpdateStatus, service)
	statusUpdate.Start()

}

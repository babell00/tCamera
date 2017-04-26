package main

import (
	"github.com/babell00/toc_camera/configuration"
	"runtime"
	"log"
	"github.com/babell00/toc_camera/camera"
	"github.com/babell00/toc_camera/server"
	"os"
	"github.com/jasonlvhit/gocron"
	"github.com/babell00/toc_camera/task"
)

func main() {
	setup()
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
	cameraService := camera.InitService(cameras)

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

func registerTasks(config configuration.Config, service *camera.CameraService){
	task.UpdateImage(service)

	imageUpdate := gocron.NewScheduler()
	imageUpdate.Every(config.ImageUpdateInterval).Seconds().Do(task.UpdateImage, service)
	imageUpdate.Start()

	statusUpdate := gocron.NewScheduler()
	statusUpdate.Every(config.StatusUpdateInterval).Seconds().Do(task.UpdateStatus, service)
	statusUpdate.Start()

}

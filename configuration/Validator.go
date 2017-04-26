package configuration

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"log"
	"runtime"
	"github.com/babell00/toc_camera/marvel"
)

const DEFAULT_PORT_NUMBER int = 8080
const DEFAULT_IMAGE_UPDATE_INTERVAL uint64 = 10
const DEFAULT_STATUS_UPDATE_INTERVAL uint64 = 5
const DEFAULT_ERROR_IMAGE string = "image/no-signal.jpg"
const DEFAULT_PUBLIC_ADDRESS = "localhost"

func validate(config *Config) {
	log.Println("Configuration validation")
	validateCpuField(config)
	validateImageUpdateIntervalField(config)
	validateStatusUpdateIntervalField(config)
	validateServerPublicAddress(config)
	validateServerPortField(config)
	validateDefaultErrorImage(config)
	validateCameraNames(config)
	validateCameraErrorImage(config)
}

func validateCpuField(config *Config) {
	switch {
	case config.MaxCpu == 0:
		config.MaxCpu = runtime.NumCPU()
		break
	case config.MaxCpu > runtime.NumCPU():
		config.MaxCpu = runtime.NumCPU()
		break
	}
}

func validateImageUpdateIntervalField(config *Config) {
	if config.ImageUpdateInterval == 0 {
		config.ImageUpdateInterval = DEFAULT_IMAGE_UPDATE_INTERVAL
	}
}

func validateStatusUpdateIntervalField(config *Config) {
	if config.ImageUpdateInterval == 0 {
		config.StatusUpdateInterval = DEFAULT_STATUS_UPDATE_INTERVAL
	}
}

func validateServerPublicAddress(config *Config){
	if config.Server.PublicAddress == "" {
		config.Server.PublicAddress = DEFAULT_PUBLIC_ADDRESS
	}
}

func validateServerPortField(config *Config) {
	switch {
	case config.Server.Port == 0:
		config.Server.Port = DEFAULT_PORT_NUMBER
		break
	case config.Server.Port <= 1024:
		config.Server.Port = DEFAULT_PORT_NUMBER
		break
	case config.Server.Port >= 65535:
		config.Server.Port = DEFAULT_PORT_NUMBER
		break
	}
}

func validateCameraNames(config *Config) {
	for k := range config.Cameras {
		if config.Cameras[k].Name == "" {
			config.Cameras[k].Name = marvel.SelectRandomName()
		}
	}
}

func validateDefaultErrorImage(config *Config) {
	if config.DefaultErrorImage == "" {
		config.DefaultErrorImage = DEFAULT_ERROR_IMAGE
	}
}

func validateCameraErrorImage(config *Config) {
	for k := range config.Cameras {
		if config.Cameras[k].ErrorImage == "" {
			config.Cameras[k].ErrorImage = DEFAULT_ERROR_IMAGE
		}
	}
}

func readYamlFile(fileName string) []byte {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return data
}

func unmarshalYaml(data []byte) *Config {
	config := Config{}
	err := yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		panic(err)
	}
	return &config
}


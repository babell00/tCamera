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

var _config *Config

func validate(config *Config) {
	log.Println("Configuration validation")

	_config = config

	validateCpuField()
	validateImageUpdateIntervalField()
	validateStatusUpdateIntervalField()
	validateServerPublicAddress()
	validateServerPortField()
	validateDefaultErrorImage()
	validateCameraNames()
	validateCameraErrorImage()
}

func validateCpuField() {
	switch {
	case _config.MaxCpu == 0:
		_config.MaxCpu = runtime.NumCPU()
		break
	case _config.MaxCpu > runtime.NumCPU():
		_config.MaxCpu = runtime.NumCPU()
		break
	}
}

func validateImageUpdateIntervalField() {
	if _config.ImageUpdateInterval == 0 {
		_config.ImageUpdateInterval = DEFAULT_IMAGE_UPDATE_INTERVAL
	}
}

func validateStatusUpdateIntervalField() {
	if _config.ImageUpdateInterval == 0 {
		_config.StatusUpdateInterval = DEFAULT_STATUS_UPDATE_INTERVAL
	}
}

func validateServerPublicAddress(){
	if _config.Server.PublicAddress == "" {
		_config.Server.PublicAddress = DEFAULT_PUBLIC_ADDRESS
	}
}

func validateServerPortField() {
	switch {
	case _config.Server.Port == 0:
		_config.Server.Port = DEFAULT_PORT_NUMBER
		break
	case _config.Server.Port <= 1024:
		_config.Server.Port = DEFAULT_PORT_NUMBER
		break
	case _config.Server.Port >= 65535:
		_config.Server.Port = DEFAULT_PORT_NUMBER
		break
	}
}

func validateCameraNames() {
	for k := range _config.Cameras {
		if _config.Cameras[k].Name == "" {
			_config.Cameras[k].Name = marvel.SelectRandomName()
		}
	}
}

func validateDefaultErrorImage() {
	if _config.DefaultErrorImage == "" {
		_config.DefaultErrorImage = DEFAULT_ERROR_IMAGE
	}
}

func validateCameraErrorImage() {
	for k := range _config.Cameras {
		if _config.Cameras[k].ErrorImage == "" {
			_config.Cameras[k].ErrorImage = DEFAULT_ERROR_IMAGE
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


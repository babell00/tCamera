package configuration

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"log"
	"runtime"
	"github.com/babell00/toc_camera/marvel"
)

const FILE_PATH string = "configuration.yaml"
const DEFAULT_PORT_NUMBER int = 8080
const DEFAULT_IMAGE_UPDATE_INTERVAL uint64 = 10
const DEFAULT_STATUS_UPDATE_INTERVAL uint64 = 5
const DEFAULT_ERROR_IMAGE string = "image/no-signal.jpg"
const DEFAULT_PUBLIC_ADDRESS = "localhost"

type Config struct {
	MaxCpu            int `yaml:"max_cpu"`
	ImageUpdateInterval   uint64 `yaml:"image_update_interval"`
	StatusUpdateInterval   uint64 `yaml:"status_update_interval"`
	DefaultErrorImage string `yaml:"default_error_image"`
	Server            Server `yaml:"server"`
	Cameras           []Camera `yaml:"cameras"`
}

type Camera struct {
	Name       string `yaml:"name"`
	Url        string `yaml:"url"`
	Path       string `yaml:"path"`
	ErrorImage string `yaml:"error_image"`
}

type Server struct {
	Port int `yaml:"port"`
	PublicAddress     string `yaml:"public_address"`
}

func ReadConfigurationFromYaml() Config {
	log.Printf("Reading configuration from: %v", FILE_PATH)
	data := readYamlFile(FILE_PATH)
	self := unmarshalYaml(data)
	validate(self)

	return *self
}

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

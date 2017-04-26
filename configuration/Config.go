package configuration

import (
	"log"
)

const FILE_PATH string = "configuration.yaml"

type Config struct {
	MaxCpu               int `yaml:"max_cpu"`
	ImageUpdateInterval  uint64 `yaml:"image_update_interval"`
	StatusUpdateInterval uint64 `yaml:"status_update_interval"`
	DefaultErrorImage    string `yaml:"default_error_image"`
	Server               Server `yaml:"server"`
	Cameras              []Camera `yaml:"cameras"`
}

type Camera struct {
	Name       string `yaml:"name"`
	MJpegUrl   string `yaml:"mjpeg_url"`
	UrlPath    string `yaml:"url_path"`
	ErrorImage string `yaml:"error_image"`
}

type Server struct {
	Port          int `yaml:"port"`
	PublicAddress string `yaml:"public_address"`
}

func ReadConfigurationFromYaml() Config {
	log.Printf("Reading configuration from: %v", FILE_PATH)
	data := readYamlFile(FILE_PATH)
	self := unmarshalYaml(data)
	validate(self)

	return *self
}

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
const DEFAULT_REFRESH_INTERVAL int = 10

type Config struct {
	MaxCpu          int `yaml:"max_cpu"`
	RefreshInterval int `yaml:"refresh_interval"`
	Server          Server `yaml:"server"`
	Cameras         []Camera `yaml:"cameras"`
}

type Camera struct {
	Name string `yaml:"name"`
	Url  string `yaml:"url"`
	Path string `yaml:"path"`
}

type Server struct {
	Port int `yaml:"port"`
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
	validateServerPortField(config)
	validateRefreshIntervalField(config)
	validateCameraNames(config)
}

func validateCameraNames(config *Config) {
	for k := range config.Cameras {
		if config.Cameras[k].Name == "" {
			config.Cameras[k].Name = marvel.SelectRandomName()
		}
	}
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

func validateRefreshIntervalField(config *Config) {
	if config.RefreshInterval == 0 {
		config.RefreshInterval = DEFAULT_REFRESH_INTERVAL
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

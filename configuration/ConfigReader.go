package configuration

import "log"

const FILE_PATH string = "configuration.yaml"

func ReadConfigurationFromYaml() Config {
	log.Printf("Reading configuration from: %v", FILE_PATH)
	data := readYamlFile(FILE_PATH)
	config := unmarshalYaml(data)
	validate(config)

	return *config
}

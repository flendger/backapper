package config

import (
	"encoding/json"
	"log"
	"os"
)

type Configuration struct {
	Port          string `json:"port"`
	AppConfigPath string `json:"app_config"`
}

func Load(confFilePath string, logger *log.Logger) *Configuration {
	fileData, err := os.ReadFile(confFilePath)
	if err != nil {
		logger.Println("Couldn't read configuration: use default instead", err)
		return defConf()
	}

	var config Configuration
	errJson := json.Unmarshal(fileData, &config)
	if errJson != nil {
		logger.Println("Couldn't decode configuration: use default instead", errJson)
		return defConf()
	}

	logger.Println("Configuration:", config)

	return &config
}

func defConf() *Configuration {
	return &Configuration{Port: "8080", AppConfigPath: "apps.json"}
}

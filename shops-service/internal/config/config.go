package config

import (
	"fmt"
	"os"

	yaml "gopkg.in/yaml.v2"
)

var conf *Configuration

// Init read and store service configuration
func Init() error {
	conf := &Configuration{}

	var isConfigRead bool
	for _, path := range configFilePaths {
		isConfigRead, _ = readConfig(conf, path)
		if isConfigRead {
			break
		}
	}

	if !isConfigRead {
		return errFileNotFound
	}

	return nil
}

func readConfig(cfg interface{}, path string) (bool, error) {
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "development"
	}

	filename := fmt.Sprintf("%s.%s.yaml", path, appEnv)
	f, err := os.Open(filename)
	if err != nil {
		return false, err
	}
	defer f.Close()

	if err = yaml.NewDecoder(f).Decode(&conf); err != nil {
		return false, err
	}

	return true, nil
}

// GetServer return server configuration
func GetServer() Server {
	return conf.Server
}

// GetServiceProducts return products-service configuration
func GetServiceProducts() ServiceProducts {
	return conf.ServiceProducts
}

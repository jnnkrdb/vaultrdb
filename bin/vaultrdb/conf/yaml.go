package conf

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

const CONFIG_YAML string = "/opt/vaultrdb/vault.yaml"

var YC *YamlConfig = &YamlConfig{}

func init() {

	// load the config from the config.yaml
	if yamlF, err := os.ReadFile(CONFIG_YAML); err != nil {
		log.Fatalf("couldn't read config file: %s\n", err.Error())
	} else {
		if err = yaml.Unmarshal(yamlF, YC); err != nil {
			log.Fatalf("couldn't parse config file: %s\n", err.Error())
		}
	}
}

type YamlConfig struct {

	// setup configs for deployment and internal communication
	Setup struct {
		Replicas    uint   `yaml:"replicas"`
		Namespace   string `yaml:"namespace"`
		ServiceName string `yaml:"servicename"`
		ExternalAPI struct {
			FrontendUI struct {
				Enabled bool `yaml:"enabled"`
			} `yaml:"frontendui"`
			SwaggerUI struct {
				Enabled bool `yaml:"enabled"`
			} `yaml:"swaggerui"`
		} `yaml:"externalapi"`
	} `yaml:"setup"`

	// set the logging configs
	Log struct {
		Level      string `yaml:"level"`
		FormatJSON bool   `yaml:"formatjson"`
	} `yaml:"loglevel"`

	// authz configs
	Authz struct {
		Enabled bool `yaml:"enabled"`
		Storage struct {
			Type string `yaml:"type"`
		} `yaml:"storage"`
	} `yaml:"authz"`

	// authz configs
	Configs struct {
		Enabled bool `yaml:"enabled"`
		Storage struct {
			Type string `yaml:"type"`
		} `yaml:"storage"`
	} `yaml:"configs"`

	// authz configs
	Vault struct {
		Enabled bool `yaml:"enabled"`
		Storage struct {
			Type string `yaml:"type"`
		} `yaml:"storage"`
	} `yaml:"vault"`
}

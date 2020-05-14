package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Config appboot config from appboot.yaml
type Config struct {
	Parameters Parameters `yaml:"parameters" json:"parameters"`
	Git        Git        `yaml:"git" json:"git"`
}

// Parameters parameters
type Parameters struct {
	StringParameters []StringParameter `yaml:"string" json:"string"`
	IntParameters    []IntParameter    `yaml:"int" json:"int"`
	FloatParameters  []FloatParameter  `yaml:"float" json:"float"`
}

// Git git
type Git struct {
	Prefix string `yaml:"prefix" json:"prefix"`
}

// StringParameter string parameter
type StringParameter struct {
	Key     string `yaml:"key" json:"key"`
	Default string `yaml:"default" json:"default"`
}

// IntParameter int parameter
type IntParameter struct {
	Key     string `yaml:"key" json:"key"`
	Default int    `yaml:"default" json:"default"`
	Min     int    `yaml:"min" json:"min"`
	Max     int    `yaml:"max" json:"max"`
}

// FloatParameter float parameter
type FloatParameter struct {
	Key     string  `yaml:"key" json:"key"`
	Default float64 `yaml:"default" json:"default"`
	Min     float64 `yaml:"min" json:"min"`
	Max     float64 `yaml:"max" json:"max"`
}

// GetConfig get config from yaml path
func GetConfig(yamlPath string) (config *Config, err error) {
	config = new(Config)
	var yamlFile []byte
	if yamlFile, err = ioutil.ReadFile(yamlPath); err != nil {
		return
	}
	if err = yaml.Unmarshal(yamlFile, config); err != nil {
		return
	}
	return
}

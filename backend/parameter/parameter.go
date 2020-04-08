package parameter

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Parameters parameters
type Parameters struct {
	StringParameters []StringParameter `yaml:"string"`
	IntParameters    []IntParameter    `yaml:"int"`
}

// StringParameter string parameter
type StringParameter struct {
	Key     string `yaml:"key"`
	Default string `yaml:"default"`
}

// IntParameter int parameter
type IntParameter struct {
	Key     string `yaml:"key"`
	Default string `yaml:"default"`
	Min     int    `yaml:"min"`
	Max     int    `yaml:"max"`
}

// GetParameters get parameters from yaml path
func GetParameters(yamlPath string) (parameters *Parameters, err error) {
	parameters = new(Parameters)
	var yamlFile []byte
	if yamlFile, err = ioutil.ReadFile(yamlPath); err != nil {
		return
	}
	if err = yaml.Unmarshal(yamlFile, parameters); err != nil {
		return
	}
	return
}

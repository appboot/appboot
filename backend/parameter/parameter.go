package parameter

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Parameters parameters
type Parameters struct {
	StringParameters []StringParameter `yaml:"string" json:"string"`
	IntParameters    []IntParameter    `yaml:"int" json:"int"`
}

// StringParameter string parameter
type StringParameter struct {
	Key     string `yaml:"key" json:"key"`
	Default string `yaml:"default" json:"default"`
}

// IntParameter int parameter
type IntParameter struct {
	Key     string `yaml:"key" json:"key"`
	Default string `yaml:"default" json:"default"`
	Min     int    `yaml:"min" json:"min"`
	Max     int    `yaml:"max" json:"max"`
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

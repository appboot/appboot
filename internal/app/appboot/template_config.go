package appboot

import (
	"io/ioutil"
	"math"
	"path"

	"github.com/appboot/appboot/configs"
	"gopkg.in/yaml.v2"
)

const (
	appboot    = "appboot"
	configYaml = "appboot.yaml"
)

// GetTemplateConfig get template config
func GetTemplateConfig(template string) (*TemplateConfig, error) {
	var result = &TemplateConfig{}

	root, err := configs.GetTemplateRoot()
	if err != nil {
		return result, err
	}

	yamlPath := path.Join(root, template, appboot, configYaml)
	return GetTemplateConfigFromYaml(yamlPath)
}

// GetTemplateConfigFromYaml get config from yaml path
func GetTemplateConfigFromYaml(yamlPath string) (config *TemplateConfig, err error) {
	config = new(TemplateConfig)
	var yamlFile []byte
	if yamlFile, err = ioutil.ReadFile(yamlPath); err != nil {
		return
	}
	if err = yaml.Unmarshal(yamlFile, config); err != nil {
		return
	}
	return
}

// TemplateConfig appboot config from appboot.yaml
type TemplateConfig struct {
	Parameters Parameters `yaml:"parameters" json:"parameters"`
	Git        Git        `yaml:"git" json:"git"`
}

// Parameters parameters
type Parameters struct {
	StringParameters []StringParameter `yaml:"string" json:"string"`
	IntParameters    []IntParameter    `yaml:"int" json:"int"`
	FloatParameters  []FloatParameter  `yaml:"float" json:"float"`
	SelectParameters []SelectParameter `yaml:"select" json:"select"`
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
	Default int64  `yaml:"default" json:"default"`
	Min     int64  `yaml:"min" json:"min"`
	Max     int64  `yaml:"max" json:"max"`
}

// FloatParameter float parameter
type FloatParameter struct {
	Key     string  `yaml:"key" json:"key"`
	Default float64 `yaml:"default" json:"default"`
	Min     float64 `yaml:"min" json:"min"`
	Max     float64 `yaml:"max" json:"max"`
}

// SelectParameter select parameter
type SelectParameter struct {
	Key     string   `yaml:"key" json:"key"`
	Options []string `yaml:"options" json:"options"`
}

// UnmarshalYAML unmarshalYAML
func (p *IntParameter) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type param IntParameter
	raw := param{
		Min: math.MinInt64,
		Max: math.MaxInt64,
	}
	if err := unmarshal(&raw); err != nil {
		return err
	}

	*p = IntParameter(raw)
	return nil
}

// UnmarshalYAML unmarshalYAML
func (p *FloatParameter) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type param FloatParameter
	raw := param{
		Min: math.SmallestNonzeroFloat64,
		Max: math.MaxFloat64,
	}
	if err := unmarshal(&raw); err != nil {
		return err
	}

	*p = FloatParameter(raw)
	return nil
}

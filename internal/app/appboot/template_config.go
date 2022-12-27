package appboot

import (
	"io/ioutil"
	"path"

	"github.com/appboot/appboot/configs"
	"gopkg.in/yaml.v2"
)

const (
	defaultValue = "default"
	appboot      = "appboot"
	configYaml   = "appboot.yaml"
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
	Parameters []any   `yaml:"parameters" json:"parameters"`
	Desc       string  `yaml:"desc" json:"desc"`
	Scripts    Scripts `yaml:"scripts" json:"scripts"`
}

// Scripts scripts
type Scripts struct {
	Before []string `yaml:"before" json:"before"`
	After  []string `yaml:"after" json:"after"`
}

// UnmarshalYAML unmarshalYAML
func (p *TemplateConfig) UnmarshalYAML(unmarshal func(any) error) error {
	type param TemplateConfig
	raw := param{}
	if err := unmarshal(&raw); err != nil {
		return err
	}

	for i, v := range raw.Parameters {
		m, ok := v.(map[any]any)
		if !ok {
			continue
		}

		t := m["type"]
		switch t {
		case "string":
			p := newStringParameter(m)
			raw.Parameters[i] = p
		case "int":
			p := newIntParameter(m)
			raw.Parameters[i] = p
		case "float":
			p := newFloatParameter(m)
			raw.Parameters[i] = p
		case "select":
			p := newSelectParameter(m)
			raw.Parameters[i] = p
		}
	}

	*p = TemplateConfig(raw)
	return nil
}

func newParameter(m map[any]any) Parameter {
	return Parameter{
		Key:  getString(m, "key"),
		Type: getString(m, "type"),
		Tip:  getString(m, "tip"),
	}
}

func newStringParameter(m map[any]any) StringParameter {
	return StringParameter{
		Parameter:   newParameter(m),
		Default:     getString(m, "default"),
		Placeholder: getString(m, "placeholder"),
	}
}

func newIntParameter(m map[any]any) IntParameter {
	return IntParameter{
		Parameter: newParameter(m),
		Min:       getInt(m, "min"),
		Max:       getInt(m, "max"),
		Default:   getInt(m, "default"),
	}
}

func newFloatParameter(m map[any]any) FloatParameter {
	return FloatParameter{
		Parameter: newParameter(m),
		Min:       getFloat(m, "min"),
		Max:       getFloat(m, "max"),
		Default:   getFloat(m, "default"),
	}
}

func newSelectParameter(m map[any]any) SelectParameter {
	mops, ok := m["options"].([]any)
	ops := []string{}
	if !ok {
		return SelectParameter{
			Parameter: newParameter(m),
			Options:   ops,
			Default:   getString(m, "default"),
		}
	}

	for _, v := range mops {
		value, ok := v.(string)
		if !ok {
			value = ""
		}
		ops = append(ops, value)
	}
	return SelectParameter{
		Parameter: newParameter(m),
		Options:   ops,
		Default:   getString(m, "default"),
	}
}

type Parameter struct {
	Key  string `yaml:"key" json:"key"`
	Type string `yaml:"type" json:"type"`
	Tip  string `yaml:"tip" json:"tip"`
}

type StringParameter struct {
	Parameter
	Default     string `yaml:"default" json:"default"`
	Placeholder string `yaml:"placeholder" json:"placeholder"`
}

// IntParameter int parameter
type IntParameter struct {
	Parameter
	Min     int `yaml:"min" json:"min"`
	Max     int `yaml:"max" json:"max"`
	Default int `yaml:"default" json:"default"`
}

// FloatParameter float parameter
type FloatParameter struct {
	Parameter
	Min     float64 `yaml:"min" json:"min"`
	Max     float64 `yaml:"max" json:"max"`
	Default float64 `yaml:"default" json:"default"`
}

// SelectParameter select parameter
type SelectParameter struct {
	Parameter
	Options []string `yaml:"options" json:"options"`
	Default string   `yaml:"default" json:"default"`
}

func getString(m map[any]any, key string) string {
	v, ok := m[key].(string)
	if ok {
		return v
	}
	return ""
}

func getInt(m map[any]any, key string) int {
	v, ok := m[key].(int)
	if ok {
		return v
	}
	return 0
}

func getFloat(m map[any]any, key string) float64 {
	v, ok := m[key].(float64)
	if ok {
		return v
	}
	return 0.0
}

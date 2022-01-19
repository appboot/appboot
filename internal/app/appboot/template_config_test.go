package appboot

import (
	"reflect"
	"testing"
)

func TestGetTemplateConfigFromYaml(t *testing.T) {
	type args struct {
		yamlPath string
	}
	tests := []struct {
		name       string
		args       args
		wantConfig *TemplateConfig
		wantErr    bool
	}{
		{
			name: "",
			args: args{
				yamlPath: "./template_config_test.yaml",
			},
			wantConfig: &TemplateConfig{
				Desc: "test template",
				Parameters: []interface{}{
					StringParameter{
						Parameter: Parameter{
							Key:  "s1",
							Type: "string",
							Tip:  "s1 tip",
						},
						Default: "s1",
					},
					StringParameter{
						Parameter: Parameter{
							Key:  "s2",
							Type: "string",
							Tip:  "s2 tip",
						},
						Default: "s2",
					},
					IntParameter{
						Parameter: Parameter{
							Key:  "i1",
							Type: "int",
							Tip:  "i1 tip",
						},
						Default: 1,
						Min:     0,
						Max:     65535,
					},
					FloatParameter{
						Parameter: Parameter{
							Key:  "f1",
							Type: "float",
							Tip:  "f1 tip",
						},
						Default: 1,
						Min:     0,
						Max:     0,
					},
					SelectParameter{
						Parameter: Parameter{
							Key:  "se1",
							Type: "select",
							Tip:  "se1 tip",
						},
						Options: []string{"o1", "o2"},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotConfig, err := GetTemplateConfigFromYaml(tt.args.yamlPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTemplateConfigFromYaml() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotConfig, tt.wantConfig) {
				t.Errorf("GetTemplateConfigFromYaml() = %v, want %v", gotConfig, tt.wantConfig)
			}
		})
	}
}

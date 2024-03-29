package appboot

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"reflect"
	"testing"

	"github.com/go-ecosystem/utils/v2/file"

	"bou.ke/monkey"
	"github.com/appboot/appboot/configs"
)

func TestApplication_Description(t *testing.T) {
	type fields struct {
		Name       string
		Path       string
		Template   string
		Parameters string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "will return description",
			fields: fields{
				Name:       "hello",
				Path:       "/test/hello",
				Template:   "HTTP",
				Parameters: "{\"a\":\"b\"}",
			},
			want: fmt.Sprintf("Name:%s \nPath:%s \nTemplate:%s \nParameters:%s\n", "hello", "/test/hello", "HTTP", "{\"a\":\"b\"}"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Application{
				Name:       tt.fields.Name,
				Path:       tt.fields.Path,
				Template:   tt.fields.Template,
				Parameters: tt.fields.Parameters,
			}
			if got := p.Description(); got != tt.want {
				t.Errorf("Application.Description() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApplication_GetTemplatePath(t *testing.T) {
	sut := &Application{
		Name:       "test",
		Path:       "test",
		Template:   "test",
		Parameters: "test",
	}

	t.Run("will return error", func(t *testing.T) {
		monkey.Patch(configs.GetTemplateRoot, func() (string, error) {
			return "", errors.New("error")
		})
		defer monkey.Unpatch(configs.GetTemplateRoot)

		p := sut.GetTemplatePath()
		if p != "" {
			t.Error("GetTemplatePath error")
		}
	})

	t.Run("will return path", func(t *testing.T) {
		monkey.Patch(configs.GetTemplateRoot, func() (string, error) {
			return "/root", nil
		})
		defer monkey.Unpatch(configs.GetTemplateRoot)

		p := sut.GetTemplatePath()
		if p != "/root/test" {
			t.Error("GetTemplatePath error")
		}
	})
}

func TestApplication_GetParameters(t *testing.T) {
	type fields struct {
		Name     string
		Path     string
		Template string
		Values   string
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]string
	}{
		{
			name: "empty Parameters will return nil",
			fields: fields{
				Name:     "hello",
				Path:     "/test",
				Template: "HTTP",
				Values:   "",
			},
			want: nil,
		},
		{
			name: "error Parameters will return nil",
			fields: fields{
				Name:     "hello",
				Path:     "/test",
				Template: "HTTP",
				Values:   "aaaass}",
			},
			want: nil,
		},
		{
			name: "Parameters with Name and Path will return Parameters's Name and Path",
			fields: fields{
				Name:     "hello",
				Path:     "/test",
				Template: "HTTP",
				Values:   `{"Name": "hello2","Path": "/test2","Other": "aaa"}`,
			},
			want: map[string]string{
				Name:    "hello2",
				Path:    "/test2",
				"Other": "aaa",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &Application{
				Name:       tt.fields.Name,
				Path:       tt.fields.Path,
				Template:   tt.fields.Template,
				Parameters: tt.fields.Values,
			}
			if got, _ := app.GetParameters(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Application.GetParameters() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestApplication_IsValid(t *testing.T) {
	type fields struct {
		Name     string
		Path     string
		Template string
		Values   string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "",
			fields: fields{},
			want:   false,
		},
		{
			name: "",
			fields: fields{
				Name: "123",
			},
			want: false,
		},
		{
			name: "",
			fields: fields{
				Path: "123",
			},
			want: false,
		},
		{
			name: "",
			fields: fields{
				Template: "123",
			},
			want: false,
		},
		{
			name: "",
			fields: fields{
				Name:     "1",
				Template: "2",
				Path:     "3",
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				Name:     "1",
				Template: "2",
				Path:     "3",
				Values:   "4",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &Application{
				Name:       tt.fields.Name,
				Path:       tt.fields.Path,
				Template:   tt.fields.Template,
				Parameters: tt.fields.Values,
			}
			if got := app.IsValid(); got != tt.want {
				t.Errorf("Application.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApplication_Clean(t *testing.T) {
	tempDir, err := ioutil.TempDir(os.TempDir(), "Clean")
	if err != nil {
		t.Errorf(err.Error())
	}
	defer os.RemoveAll(tempDir)

	sut := &Application{
		Name:       "test",
		Path:       tempDir,
		Template:   "test",
		Parameters: "test",
	}

	configFolder := path.Join(tempDir, ConfigFolder)
	if err = os.MkdirAll(configFolder, 0755); err != nil {
		t.Errorf(err.Error())
	}

	sut.Clean()

	if file.Exists(configFolder) {
		t.Error("Clean() error")
	}
}

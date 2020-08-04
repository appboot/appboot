package configs

import (
	"github.com/CatchZeng/gutils/file"
	"os"
	"path"

	"github.com/mitchellh/go-homedir"
)

const (
	// TemplatesSubPath templates default sub path
	TemplatesSubPath = ".appboot/templates/"
	// TemplatesSource templates default source
	TemplatesSource = "https://github.com/appboot/templates.git"
)

// GetTemplateRoot get template root
func GetTemplateRoot() (string, error) {
	if root, err := GetConfig("templateRoot"); err == nil {
		if len(root) > 0 {
			return root, nil
		}
	}

	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}

	root := path.Join(home, TemplatesSubPath)
	if !file.Exists(root) {
		if err := os.MkdirAll(root, 0755); err != nil {
			return root, nil
		}
		return root, nil
	}
	return root, nil
}

// GetTemplateSource get template source
func GetTemplateSource() string {
	if source, err := GetConfig("templateSource"); err == nil {
		return source
	}
	return TemplatesSource
}

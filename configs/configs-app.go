package configs

import (
	"os"
	"path"

	"github.com/appboot/appboot/internal/pkg/common"
	"github.com/go-ecosystem/utils/v2/file"
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

	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	root := path.Join(home, TemplatesSubPath)
	if !file.Exists(root) {
		if err := os.MkdirAll(root, common.DefaultFileMode); err != nil {
			return root, err
		}
		return root, nil
	}
	return root, nil
}

// GetTemplateSource get template source
func GetTemplateSource() string {
	if source, err := GetConfig("templateSource"); err == nil {
		if len(source) < 1 {
			return TemplatesSource
		}
		return source
	}
	return TemplatesSource
}

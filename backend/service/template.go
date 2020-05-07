package service

import (
	"github.com/appboot/appbctl/config"
	"github.com/appboot/appboot/utils"
)

// GetTemplates get templates
func GetTemplates() []string {
	var templates []string

	root, err := config.GetTemplateRoot()
	if err != nil {
		return templates
	}

	templates, _ = utils.GetDirList(root)
	return templates
}

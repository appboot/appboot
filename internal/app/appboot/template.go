package appboot

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/go-ecosystem/utils/array"
	"github.com/go-ecosystem/utils/log"
	"gopkg.in/yaml.v2"

	"github.com/appboot/appboot/configs"
	"github.com/go-ecosystem/utils/file"
	gos "github.com/go-ecosystem/utils/os"
)

// TemplatesConfig templates config
type TemplatesConfig struct {
	Groups []*TemplateGroup `yaml:"groups" json:"groups"`
}

// TemplateGroup template group
type TemplateGroup struct {
	ID        string      `yaml:"id" json:"id"`
	Desc      string      `yaml:"desc" json:"desc"`
	Templates []*Template `yaml:"templates" json:"templates"`
}

// Template template
type Template struct {
	ID   string `yaml:"id" json:"id"`
	Desc string `yaml:"desc" json:"desc"`
}

// GetTemplateGroups get templates
func GetTemplateGroups() []*TemplateGroup {
	tc, err := GetTemplatesConfig()
	if err == nil {
		return tc.Groups
	}

	names := GetTemplateNames()
	templates := []*Template{}
	for _, name := range names {
		config, err := GetTemplateConfig(name)
		if err != nil {
			config = &TemplateConfig{}
		}
		t := &Template{
			ID:   name,
			Desc: config.Desc,
		}
		templates = append(templates, t)
	}
	groups := []*TemplateGroup{{
		ID:        defaultValue,
		Desc:      defaultValue,
		Templates: templates,
	}}
	return groups
}

// GetTemplatesConfig get templates config
func GetTemplatesConfig() (*TemplatesConfig, error) {
	var result = &TemplatesConfig{}

	root, err := configs.GetTemplateRoot()
	if err != nil {
		return result, err
	}

	yamlPath := path.Join(root, configYaml)
	return GetTemplatesConfigFromYaml(yamlPath)
}

// GetTemplatesConfigFromYaml get config from yaml path
func GetTemplatesConfigFromYaml(yamlPath string) (config *TemplatesConfig, err error) {
	config = new(TemplatesConfig)
	var yamlFile []byte
	if yamlFile, err = ioutil.ReadFile(yamlPath); err != nil {
		return
	}
	if err = yaml.Unmarshal(yamlFile, config); err != nil {
		return
	}
	return
}

// GetTemplateNames get the names of all templates
func GetTemplateNames() []string {
	var templates []string

	root, err := configs.GetTemplateRoot()
	if err != nil {
		return templates
	}

	templates, _ = file.GetDirListWithFilter(root, func(info os.FileInfo) bool {
		return !strings.HasPrefix(info.Name(), ".")
	})
	return templates
}

// GetTemplatesGitHash get templates git hash value
func GetTemplatesGitHash() string {
	root, err := configs.GetTemplateRoot()
	if err != nil {
		return ""
	}

	res, err := gos.RunCommand(context.Background(), root, "git", "rev-parse", "--short", "HEAD")
	if err != nil {
		return ""
	}
	res = strings.Replace(res, "\n", "", -1)
	return res
}

// UpdateTemplate update template with Git
func UpdateTemplate(name string) error {
	return UpdateTemplateWithDownloader(name, NewDownloader())
}

// UpdateTemplateWithDownloader update template
func UpdateTemplateWithDownloader(name string, downloader Downloader) error {
	// get template path
	root, err := configs.GetTemplateRoot()
	if err != nil {
		return err
	}
	templatePath := path.Join(root, name)

	// download templates
	tempDir, err := downloadTemplates(downloader)
	defer os.RemoveAll(tempDir)
	if err != nil {
		if file.Exists(templatePath) {
			log.W("Update template error: %v \nuse old template.", err)
			return nil
		}
		return err
	}

	// check template
	src := path.Join(tempDir, name)
	if !file.Exists(src) {
		log.W("Can not get template %v from %v", name, configs.GetTemplateSource())
		return nil
	}

	// update template
	if err := updateTemplate(name, root, src); err != nil {
		return err
	}

	if !file.Exists(templatePath) {
		return fmt.Errorf("can not find template from appboot templates, you can add your custom template to %s", root)
	}

	return nil
}

func updateTemplate(name, root, src string) error {
	templatePath := path.Join(root, name)

	// recreate template directory
	mode := file.Mode(templatePath)
	_ = os.RemoveAll(templatePath)
	if err := os.MkdirAll(templatePath, mode); err != nil {
		return err
	}

	// copy template to root from src
	cp := "cp -rf " + src + " " + root
	if err := gos.RunBashCommand(cp); err != nil {
		return err
	}

	return nil
}

// UpdateAllTemplates update all templates with git
func UpdateAllTemplates() error {
	return UpdateAllTemplatesWithDownloader(NewDownloader())
}

// UpdateAllTemplatesWithDownloader update all templates
func UpdateAllTemplatesWithDownloader(downloader Downloader) error {
	// download templates
	tempDir, err := downloadTemplates(downloader)
	defer os.RemoveAll(tempDir)
	if err != nil {
		return err
	}

	// get template root
	root, err := configs.GetTemplateRoot()
	if err != nil {
		return err
	}

	// remove existed templates
	templates := GetTemplateNames()
	for _, name := range templates {
		list, _ := file.GetDirListWithFilter(tempDir, func(info os.FileInfo) bool {
			return !strings.HasPrefix(info.Name(), ".")
		})
		if array.ContainString(list, name) {
			existed := path.Join(root, name)
			os.RemoveAll(existed)
		}
	}

	// update templates
	cp := "cp -rf " + tempDir + "/*" + " " + root
	command := "shopt -s dotglob && " + cp
	if err := gos.RunBashCommand(command); err != nil {
		return err
	}

	return nil
}

// Warning: caller should clean tempDir after used
// defer os.RemoveAll(tempDir)
func downloadTemplates(downloader Downloader) (string, error) {
	source := configs.GetTemplateSource()
	tempDir, err := ioutil.TempDir(os.TempDir(), "template")
	if err != nil {
		return tempDir, err
	}
	if err := downloader.Run(source, tempDir); err != nil {
		return tempDir, err
	}
	return tempDir, nil
}

package appboot

import (
	"container/list"
	"fmt"
	"github.com/CatchZeng/gutils/file"
	gos "github.com/CatchZeng/gutils/os"
	"github.com/appboot/appboot/configs"
	"github.com/appboot/appboot/internal/pkg/logger"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type File struct {
	Path    string
	Content string
}

// GetTemplates get templates
func GetTemplates() []string {
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

var paths = list.New()

// GetFiles get files from template path
func GetFiles(templatePath string) (*list.List, error) {
	var files = list.New()
	if err := filepath.Walk(templatePath, walkFunc); err != nil {
		return files, err
	}

	for i := paths.Front(); i != nil; i = i.Next() {
		p := i.Value.(string)
		bytes, err := ioutil.ReadFile(p)
		if err != nil {
			return files, err
		}
		content := string(bytes)
		f := File{
			Path:    p,
			Content: content,
		}
		files.PushBack(f)
	}
	return files, nil
}

// UpdateTemplate update template with Git
func UpdateTemplate(name string) error {
	return UpdateTemplateWithDownloader(name, NewDownloader())
}

// UpdateTemplateWithDownloader update template
func UpdateTemplateWithDownloader(name string, downloader Downloader) error {
	root, err := configs.GetTemplateRoot()
	if err != nil {
		return err
	}

	templatePath := path.Join(root, name)
	templateSource := configs.GetTemplateSource()
	tempDir, err := downloadTemplates(templateSource, downloader)
	defer os.RemoveAll(tempDir)
	if err != nil {
		if file.Exists(templatePath) {
			logger.LogW(fmt.Sprintf("update template error: %v \nuse old template.", err))
			return nil
		}
		return err
	}

	src := path.Join(tempDir, name)

	_ = os.RemoveAll(templatePath)
	if err = os.MkdirAll(templatePath, 0755); err != nil {
		return err
	}

	cp := "cp -rf " + src + " " + root
	if err := gos.RunBashCommand(cp); err != nil {
		return err
	}
	if !file.Exists(templatePath) {
		return fmt.Errorf("can not find template from appboot templates, you can add your custom template to %s", root)
	}
	return nil
}

// UpdateAllTemplates update all templates with git
func UpdateAllTemplates() error {
	return UpdateAllTemplatesWithDownloader(NewDownloader())
}

// UpdateAllTemplatesWithDownloader update all templates
func UpdateAllTemplatesWithDownloader(downloader Downloader) error {
	root, err := configs.GetTemplateRoot()
	if err != nil {
		return err
	}

	templateSource := configs.GetTemplateSource()
	tempDir, err := downloadTemplates(templateSource, downloader)
	defer os.RemoveAll(tempDir)
	if err != nil {
		return err
	}

	cp := "cp -rf " + tempDir + "/*" + " " + root
	if err := gos.RunBashCommand(cp); err != nil {
		return err
	}

	return nil
}

// Warning: caller should clean tempDir after used
// defer os.RemoveAll(tempDir)
func downloadTemplates(source string, downloader Downloader) (string, error) {
	tempDir, err := ioutil.TempDir(os.TempDir(), "template")
	if err != nil {
		return tempDir, err
	}
	if err := downloader.Run(source, tempDir); err != nil {
		return tempDir, err
	}
	return tempDir, nil
}

func walkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if !info.IsDir() {
		paths.PushBack(path)
	}
	return nil
}

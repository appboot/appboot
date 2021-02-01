package appboot

import (
	"errors"

	"github.com/go-ecosystem/utils/os"
)

// Downloader downloader interface
type Downloader interface {
	Run(source string, destination string) error
}

type gitDownloader struct {
}

// NewDownloader new downloader
func NewDownloader() Downloader {
	return &gitDownloader{}
}

// Run run git download
func (g *gitDownloader) Run(source string, destination string) error {
	if len(source) < 1 {
		return errors.New("source is empty")
	}
	cmd := "git clone " + source + " " + destination
	return os.RunBashCommand(cmd)
}

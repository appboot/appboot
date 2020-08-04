package appboot

import (
	"errors"
	gos "github.com/CatchZeng/gutils/os"
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

var runBashCommand = gos.RunBashCommand

// Run run git download
func (g *gitDownloader) Run(source string, destination string) error {
	if len(source) < 1 {
		return errors.New("source is empty")
	}
	cmd := "git clone " + source + " " + destination
	return runBashCommand(cmd)
}

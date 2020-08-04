package appboot

import (
	"errors"
	"testing"
)

func TestGitDownloader_Run(t *testing.T) {
	downloader := NewDownloader()
	t.Run("when source is empty will return error", func(t *testing.T) {
		if err := downloader.Run("", ""); err == nil {
			t.Error(err)
		}
	})

	runBashCommand = func(cmd string) error {
		return nil
	}
	t.Run("when runBashCommand passed will return no error", func(t *testing.T) {
		if err := downloader.Run("sss", ""); err != nil {
			t.Error(err)
		}
	})

	runBashCommand = func(cmd string) error {
		return errors.New("")
	}
	t.Run("when runBashCommand failed will return error", func(t *testing.T) {
		if err := downloader.Run("sss", ""); err == nil {
			t.Error(err)
		}
	})
}

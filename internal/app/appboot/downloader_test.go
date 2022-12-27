package appboot

import (
	"errors"
	"testing"

	"bou.ke/monkey"
	"github.com/go-ecosystem/utils/v2/os"
)

func TestGitDownloader_Run(t *testing.T) {
	downloader := NewDownloader()

	t.Run("when source is empty will return error", func(t *testing.T) {
		if err := downloader.Run("", ""); err == nil {
			t.Error(err)
		}
	})

	t.Run("when runBashCommand passed will return no error", func(t *testing.T) {
		monkey.Patch(os.RunBashCommand, func(cmd string) error {
			return nil
		})
		defer monkey.Unpatch(os.RunBashCommand)

		if err := downloader.Run("sss", ""); err != nil {
			t.Error(err)
		}
	})

	t.Run("when runBashCommand failed will return error", func(t *testing.T) {
		monkey.Patch(os.RunBashCommand, func(cmd string) error {
			return errors.New("")
		})
		defer monkey.Unpatch(os.RunBashCommand)

		if err := downloader.Run("sss", ""); err == nil {
			t.Error(err)
		}
	})
}

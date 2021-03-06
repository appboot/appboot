package path

import (
	"path"
	"testing"

	"github.com/mitchellh/go-homedir"
)

func TestHandleHomedir(t *testing.T) {
	got := HandleHomedir("~/test")

	home, err := homedir.Dir()
	if err != nil {
		t.Errorf(err.Error())
	}
	want := path.Join(home, "test")

	if got != want {
		t.Errorf("HandleHomedir() = %v, want %v", got, want)
	}

	got = HandleHomedir("test")
	want = "test"

	if got != want {
		t.Errorf("HandleHomedir() = %v, want %v", got, want)
	}
}

package path

import (
	"os"
	"path"
	"testing"
)

func TestHandleHomedir(t *testing.T) {
	got := HandleHomedir("~/test")

	home, err := os.UserHomeDir()
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

	got = HandleHomedir("$home/test")
	want = path.Join(home, "test")
	if got != want {
		t.Errorf("HandleHomedir() = %v, want %v", got, want)
	}

	got = HandleHomedir("$HOME/test")
	want = path.Join(home, "test")
	if got != want {
		t.Errorf("HandleHomedir() = %v, want %v", got, want)
	}

	got = HandleHomedir("$Home/test")
	want = "$Home/test"
	if got != want {
		t.Errorf("HandleHomedir() = %v, want %v", got, want)
	}
}

func TestHandlePWD(t *testing.T) {
	got := HandlePWD("~/test")
	want := "~/test"
	if got != want {
		t.Errorf("HandlePWD() = %v, want %v", got, want)
	}

	got = HandlePWD("$PWD/test")
	wd, _ := os.Getwd()
	want = path.Join(wd, "test")
	if got != want {
		t.Errorf("HandlePWD() = %v, want %v", got, want)
	}
}

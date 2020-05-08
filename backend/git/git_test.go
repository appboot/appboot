package git

import (
	"errors"
	"testing"
)

func TestPush(t *testing.T) {
	type args struct {
		gitURL     string
		codeFloder string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "gitURL empty will return error",
			args: args{
				gitURL:     "",
				codeFloder: "aaa",
			},
			wantErr: true,
		},
		{
			name: "codeFloder empty will return error",
			args: args{
				gitURL:     "aaa",
				codeFloder: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Push(tt.args.gitURL, tt.args.codeFloder); (err != nil) != tt.wantErr {
				t.Errorf("Push() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	runBashCommand = func(cmd string) (err error) {
		return nil
	}

	if err := Push("github.com/test.git", "test"); err != nil {
		t.Errorf("Push should be successful, when runBashCommand return nil")
	}

	runBashCommand = func(cmd string) (err error) {
		return errors.New("error")
	}

	if err := Push("github.com/test.git", "test"); err == nil {
		t.Errorf("Push will fail, when runBashCommand return error")
	}
}

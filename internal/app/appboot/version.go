package appboot

import (
	"bytes"
	"runtime"
	"text/template"
)

var (
	// Version for appboot
	Version = "1.0.0"
	// BuildTime for appboot
	BuildTime = "2020/08/04"
)

// Options for appboot
type Options struct {
	GitCommit string
	Version   string
	BuildTime string
	GoVersion string
	Os        string
	Arch      string
}

var versionTemplate = `Version:      {{.Version}}
Go version:   {{.GoVersion}}
Built:        {{.BuildTime}}
OS/Arch:      {{.Os}}/{{.Arch}}`

// DefaultOps default options
var DefaultOps = Options{
	Version:   Version,
	BuildTime: BuildTime,
	GoVersion: runtime.Version(),
	Os:        runtime.GOOS,
	Arch:      runtime.GOARCH,
}

// GetVersion get version string
func GetVersion() string {
	return GetVersionWithOps(DefaultOps)
}

// GetVersionWithOps get version string with versionOptions
func GetVersionWithOps(options Options) string {
	var doc bytes.Buffer
	tml, _ := template.New("version").Parse(versionTemplate)
	_ = tml.Execute(&doc, options)
	return doc.String()
}

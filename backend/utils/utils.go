package utils

import (
	"fmt"
	"net"
	"os"
	"path"

	"github.com/mitchellh/go-homedir"
)

// GetIP get IP
func GetIP() string {
	envIP := os.Getenv("HOST_IP")
	if len(envIP) > 0 {
		return envIP
	}

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	for _, address := range addrs {
		if inet, ok := address.(*net.IPNet); ok && !inet.IP.IsLoopback() {
			if inet.IP.To4() != nil {
				return inet.IP.String()
			}
		}
	}
	return ""
}

// GetSavePath get save path
func GetSavePath(appName string) string {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	savePath := path.Join(home, ".appboot", ".workspace", appName)
	return savePath
}

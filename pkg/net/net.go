package net

import (
	"fmt"
	"net"
	"os"
)

// GetIP get IP
func GetIP() string {
	envIP := os.Getenv("HOST_ADDRESS")
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

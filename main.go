package main

import (
	"github.com/appboot/appboot/cmd/appboot/cmd"
	"log"
)

func main() {
	log.SetFlags(0)
	cmd.Execute()
}

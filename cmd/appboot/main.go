package main

import (
	"log"

	"github.com/appboot/appboot/cmd/appboot/cmd"
)

func main() {
	log.SetFlags(0)
	cmd.Execute()
}

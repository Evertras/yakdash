package main

import (
	"log"

	"github.com/evertras/yakdash/cmd/yakdash/cmds"
)

func main() {
	if err := cmds.Execute(); err != nil {
		log.Fatal(err)
	}
}

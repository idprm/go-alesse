package main

import (
	"github.com/idprm/go-alesse/src/cmd"
	"github.com/idprm/go-alesse/src/database"
)

func init() {
	database.Connect()
}

func main() {

	// setup cobra cmd
	cmd.Execute()
}

package main

import (
	"time"

	"github.com/idprm/go-alesse/src/cmd"
	"github.com/idprm/go-alesse/src/database"
	"github.com/idprm/go-alesse/src/pkg/config"
)

func init() {
	database.Connect()

}

func main() {

	loc, _ := time.LoadLocation(config.ViperEnv("TIME_ZONE"))
	// handle err
	time.Local = loc // -> this is setting the global timezone

	// setup cobra cmd
	cmd.Execute()

}

package main

import (
	"log"
	"time"

	"github.com/idprm/go-alesse/src/cmd"
	"github.com/idprm/go-alesse/src/database"
)

func init() {
	database.Connect()
	log.Println(time.Local)
}

func main() {

	// setup cobra cmd
	cmd.Execute()

	loc, _ := time.LoadLocation("Asia/Jakarta")
	// handle err
	time.Local = loc // -> this is setting the global timezone
}

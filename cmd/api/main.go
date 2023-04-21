package main

import (
	"log"
	"rpg-go/config"
	"runtime"
)

var app config.AppConfig

func main() {
	svc, err := setup()
	if err != nil {
		panic(err)
	}
	defer app.DB.Close()

	// print server info
	log.Printf("******************************************")
	log.Printf("** %sRPG-GO%s v%s built in %s", "\033[31m", "\033[0m", app.Version, runtime.Version())
	log.Printf("**----------------------------------------")
	log.Printf("** Running with %d Processors", runtime.NumCPU())
	log.Printf("** Running on %s", runtime.GOOS)
	log.Printf("******************************************")

	log.Printf("Starting HTTP server on port %s.", app.WebServerPort)

	err = svc.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

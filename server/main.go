package main

import (
	"log"

	"github.wdf.sap.corp/I554249/sensor/app"
)

func main() {
	err := app.Run()
	if err != nil {
		log.Fatalln(err)
	}
}

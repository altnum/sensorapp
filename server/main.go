package main

import (
	"log"

	"github.com/altnum/sensorapp/app"
)

func main() {
	err := app.Run()
	if err != nil {
		log.Fatalln(err)
	}
}

package main

import (
	"log"

	"multiverse/notifier/app"
)

func init() {
	log.Println("Environment variables loaded...")
}

func main() {
	app := app.NewApp()
	log.Fatalln(app.Start(""))
}

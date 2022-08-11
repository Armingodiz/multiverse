package main

import (
	"log"

	"multiverse/notifier/app"

	"multiverse/notifier/config"

	"github.com/joho/godotenv"
)

func init() {
	log.SetFlags(log.Llongfile)
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	err = config.Load()
	if err != nil {
		panic(err)
	}
	log.Println("Environment variables loaded...")
}

func main() {
	app := app.NewApp()
	log.Fatalln(app.Start(""))
}

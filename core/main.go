package main

import (
	"log"

	"multiverse/core/app"
	"multiverse/core/config"

	"github.com/joho/godotenv"
)

func init() {
	log.SetFlags(log.Llongfile)
	err := godotenv.Load()
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
	log.Fatalln(app.Start(":" + config.Configs.App.Port))
}

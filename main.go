package main

import (
	"fmt"
	"log"
	"os"

	"EventApp/database"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	cmds := os.Args[1:]
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("cannot find env file")
	}
	database.Init()
	if helper.StringInArr("--migrate", cmds) {
		database.Run()
	}
	event_app := fiber.New()
	InitRoutes(event_app)

	log.Fatal(event_app.Listen(os.Getenv("PORT")))
}

func InitRoutes(app *fiber.App) {

}

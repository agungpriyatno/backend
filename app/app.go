package app

import (
	"log"
	"os"

	"github.com/agungpriyatno/olap-backend/app/configs"
	"github.com/agungpriyatno/olap-backend/app/routers"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func RunApp() {
	godotenv.Load()
	app := fiber.New()
	configs.ConnectDB()
	routers.Router(app)
	log.Fatalln(app.Listen(os.Getenv("SERVER")))
}

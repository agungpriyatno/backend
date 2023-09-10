package app

import (
	"log"
	"os"

	"github.com/agungpriyatno/olap-backend/app/configs"
	"github.com/agungpriyatno/olap-backend/app/routers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func RunApp() {
	godotenv.Load()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
		AllowMethods: "*",
	}))
	configs.ConnectDB()
	routers.Router(app)
	log.Fatalln(app.Listen(os.Getenv("SERVER")))
}

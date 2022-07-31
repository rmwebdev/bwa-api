package main

import (
	"api/app/config"
	"api/app/lib"
	"api/app/routes"
	"api/app/services"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func init() {
	lib.LoadEnvironment(config.Environment)
	services.InitRedis()
	services.InitDatabase()
}

// @title Master API
// @version 1.0.0
// @description API Documentation
// @termsOfService https://bbdocs.monstercode.net/en/guide/tnc.html
// @contact.name Developer
// @contact.email developer.team.tog@gmail.com
// @host bbdev.monstercode.net
// @schemes https http
// @BasePath /api/v1/master
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	app := fiber.New(fiber.Config{
		Prefork: viper.GetString("PREFORK") == "true",
	})

	routes.Handle(app)
	log.Fatal(app.Listen(":" + viper.GetString("PORT")))
}

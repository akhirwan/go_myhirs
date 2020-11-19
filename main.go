package main

import (
	"api_go/config"
	"api_go/controller"
	"api_go/exception"
	"api_go/repository"
	"api_go/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// Setup Configuration
	configuration := config.New()
	database := config.NewMongoDatabase(configuration)

	// Setup Repository
	companyRepository := repository.NewCompanyRepository(database)

	// Setup Service
	companyService := service.NewCompanyService(&companyRepository)

	// Setup Controller
	companyController := controller.NewCompanyController(&companyService)

	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	companyController.Route(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, UnderWorld!")
	})

	err := app.Listen(":3333")
	exception.PanicIfNeeded(err)
}

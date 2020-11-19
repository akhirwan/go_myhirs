package controller

import (
	"api_go/exception"
	"api_go/model"
	"api_go/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CompanyController struct {
	CompanyService service.CompanyService
}

func NewCompanyController(companyService *service.CompanyService) CompanyController {
	return CompanyController{CompanyService: *companyService}
}

func (controller *CompanyController) Route(app *fiber.App) {
	app.Post("api/company", controller.Create)
	app.Get("api/companies", controller.List)
}

func (controller *CompanyController) Create(c *fiber.Ctx) error {
	var request model.CreateCompanyRequest

	request.Id = uuid.New().String()

	err := c.BodyParser(&request)
	exception.PanicIfNeeded(err)

	response := controller.CompanyService.Create(request)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *CompanyController) List(c *fiber.Ctx) error {
	responses := controller.CompanyService.List()
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   responses,
	})
}

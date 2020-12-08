package controller

import (
	"api_go/exception"
	"api_go/model"
	"api_go/service"
	"fmt"

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
	app.Get("api/company/:id", controller.Detail)
	app.Put("api/company", controller.Edit)
	app.Put("api/company/:id", controller.Delete)
}

func (controller *CompanyController) Create(c *fiber.Ctx) error {
	var request model.CreateCompanyRequest

	request.Id = uuid.New().String()

	err := c.BodyParser(&request)
	exception.PanicIfNeeded(err)

	response := controller.CompanyService.Create(request)
	return c.JSON(model.WebResponse{
		Code:    200,
		Status:  "OK",
		Message: "Insert Company Successfull",
		Data:    response,
	})
}

func (controller *CompanyController) List(c *fiber.Ctx) error {
	responses := controller.CompanyService.List()
	return c.JSON(model.WebResponse{
		Code:    200,
		Status:  "OK",
		Message: "Get Companies Data Successfull",
		Data:    responses,
	})
}

func (controller *CompanyController) Detail(c *fiber.Ctx) error {
	compId := c.Params("id")
	responses := controller.CompanyService.Detail(compId)
	fmt.Println(responses, compId)

	return c.JSON(model.WebResponse{
		Code:    200,
		Status:  "OK",
		Message: "Get Detail Company Successfull",
		Data:    responses,
	})
}

func (controller *CompanyController) Edit(c *fiber.Ctx) error {
	return c.JSON(model.WebResponse{
		Code:    200,
		Status:  "OK",
		Message: "Update Company Successfull",
		Data:    "response",
	})
}

func (controller *CompanyController) Delete(c *fiber.Ctx) error {
	compId := c.Params("id")
	fmt.Println(compId)

	return c.JSON(model.WebResponse{
		Code:    200,
		Status:  "OK",
		Message: "Delete Company Successfull",
		Data:    "response",
	})
}

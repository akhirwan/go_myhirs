package validation

import (
	"api_go/exception"
	"api_go/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func CompanyValidate(request model.CreateCompanyRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Id, validation.Required),
		validation.Field(&request.Name, validation.Required),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}

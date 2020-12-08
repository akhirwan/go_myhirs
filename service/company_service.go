package service

import (
	"api_go/model"
)

type CompanyService interface {
	Create(request model.CreateCompanyRequest) (response model.CreateCompanyResponse)
	Detail(id string) (responses []model.ShowCompanyresponse)
	List() (responses []model.GetCompanyresponse)
}

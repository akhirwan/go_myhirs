package service

import (
	"api_go/entity"
	"api_go/model"
	"api_go/repository"
	"api_go/validation"
	"time"
)

func NewCompanyService(companyRepository *repository.CompanyRepository) CompanyService {
	return &companyServiceImpl{
		CompanyRepository: *companyRepository,
	}
}

type companyServiceImpl struct {
	CompanyRepository repository.CompanyRepository
}

func (service *companyServiceImpl) Create(request model.CreateCompanyRequest) (response model.CreateCompanyResponse) {
	validation.CompanyValidate(request)

	company := entity.Company{
		Id:         request.Id,
		Name:       request.Name,
		CreatedAt:  time.Now().UnixNano() / int64(time.Millisecond),
		ModifiedAt: time.Now().UnixNano() / int64(time.Millisecond),
		CreatedBy:  request.CreatedBy,
		ModifiedBy: request.CreatedBy,
		IsDeleted:  false,
	}

	service.CompanyRepository.Insert(company)

	response = model.CreateCompanyResponse{
		Id:         company.Id,
		Name:       company.Name,
		CreatedAt:  company.CreatedAt,
		ModifiedAt: company.ModifiedAt,
		CreatedBy:  company.CreatedBy,
		ModifiedBy: company.ModifiedBy,
		IsDeleted:  company.IsDeleted,
	}

	return response
}

func (service *companyServiceImpl) List() (responses []model.GetCompanyresponse) {
	companies := service.CompanyRepository.FindAll()
	for _, company := range companies {
		responses = append(responses, model.GetCompanyresponse{
			Id:         company.Id,
			Name:       company.Name,
			CreatedAt:  company.CreatedAt,
			ModifiedAt: company.ModifiedAt,
			CreatedBy:  company.CreatedBy,
			ModifiedBy: company.ModifiedBy,
			IsDeleted:  company.IsDeleted,
		})
	}
	return responses
}

func (service *companyServiceImpl) Detail() (responses []model.ShowCompanyresponse) {
	companies := service.CompanyRepository.Show()
	for _, company := range companies {
		responses = append(responses, model.ShowCompanyresponse{
			Id:         company.Id,
			Name:       company.Name,
			CreatedAt:  company.CreatedAt,
			ModifiedAt: company.ModifiedAt,
			CreatedBy:  company.CreatedBy,
			ModifiedBy: company.ModifiedBy,
			IsDeleted:  company.IsDeleted,
		})
	}
	return responses
}

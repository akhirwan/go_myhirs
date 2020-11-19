package repository

import "api_go/entity"

type CompanyRepository interface {
	Insert(addcompany entity.Company)
	// Update(editcompany entity.Company)
	// Remove(delcompany entity.Company)
	// Detail(findcompany entity.Company)

	FindAll() (companies []entity.Company)
}

package repository

import "api_go/entity"

type CompanyRepository interface {
	Insert(addcompany entity.Company)
	// Update(editcompany entity.Company)
	// Remove(delcompany entity.Company)
	Show() (companies []entity.Company)

	FindAll() (companies []entity.Company)
}

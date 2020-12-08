package repository

import "api_go/entity"

type CompanyRepository interface {
	Insert(addcompany entity.Company)
	// Update(id string) (editcompany entity.Company)
	// Remove(delcompany entity.Company)
	Show(id string) (companies []entity.Company)

	FindAll() (companies []entity.Company)
}

package repository

import (
	"api_go/config"
	"api_go/entity"
	"api_go/exception"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewCompanyRepository(database *mongo.Database) CompanyRepository {
	return &companyRepositoryImpl{
		Collection: database.Collection("company"),
	}
}

type companyRepositoryImpl struct {
	Collection *mongo.Collection
}

func (repository *companyRepositoryImpl) Insert(company entity.Company) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	_, err := repository.Collection.InsertOne(ctx, bson.M{
		"_id":         company.Id,
		"name":        company.Name,
		"created_at":  company.CreatedAt,
		"modified_at": company.ModifiedAt,
		"created_by":  company.CreatedBy,
		"modified_by": company.ModifiedBy,
		"is_deleted":  company.IsDeleted,
	})
	exception.PanicIfNeeded(err)
}

func (repository *companyRepositoryImpl) FindAll() (companies []entity.Company) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	cursor, err := repository.Collection.Find(ctx, bson.M{})
	exception.PanicIfNeeded(err)

	var documents []bson.M
	err = cursor.All(ctx, &documents)
	exception.PanicIfNeeded(err)

	for _, document := range documents {
		companies = append(companies, entity.Company{
			Id:         document["_id"].(string),
			Name:       document["name"].(string),
			CreatedAt:  document["created_at"].(int64),
			ModifiedAt: document["modified_at"].(int64),
			CreatedBy:  document["created_by"].(string),
			ModifiedBy: document["modified_by"].(string),
			IsDeleted:  document["is_deleted"].(bool),
		})
	}

	return companies
}

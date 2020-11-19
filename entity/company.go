package entity

type Company struct {
	Id         string
	Name       string
	CreatedAt  int64
	ModifiedAt int64
	CreatedBy  string
	ModifiedBy string
	IsDeleted  bool
}

package repositories

import "gorm.io/gorm"

type Repo struct {
	AccountCustomRepo AccountCustomRepo
}

func NewRepository(db *gorm.DB) *Repo {
	return &Repo{
		NewAccountCustom(db),
	}
}

package repositories

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type AccountCustomRepo interface {
	FindAccountById(id uint64) (Accounts, error)
	FindAccountByEmail(email string) (Accounts, error)
	ValidateAccountByEmail(email string) (bool, error)
	CreateAccount(account Accounts) (Accounts, error)
	UpdateAccount(id uint64, fullName, email string) error
}

type accountCustomRepo struct {
	db *gorm.DB
}

type Accounts struct {
	Id          uint64     `gorm:"column:id"`
	FullName    string     `gorm:"column:full_name"`
	Email       string     `gorm:"column:email"`
	Password    string     `gorm:"column:password"`
	CreatedDate time.Time  `gorm:"column:created_date;autoCreateTime"`
	UpdatedDate *time.Time `gorm:"column:updated_date"`
}

func (tb *Accounts) TableName() string {
	return "accounts"
}

func NewAccountCustom(db *gorm.DB) AccountCustomRepo {
	return &accountCustomRepo{
		db: db,
	}
}

func (repo *accountCustomRepo) FindAccountById(id uint64) (Accounts, error) {
	resp := Accounts{}
	result := repo.db.Model(&Accounts{}).Where("id = ?", id).First(&resp)
	if result.Error != nil {
		return resp, result.Error
	}

	return resp, nil
}

func (repo *accountCustomRepo) FindAccountByEmail(email string) (Accounts, error) {
	resp := Accounts{}
	result := repo.db.Model(&Accounts{}).Where("email = ?", email).First(&resp)
	if result.Error != nil {
		return resp, result.Error
	}

	return resp, nil
}

func (repo *accountCustomRepo) ValidateAccountByEmail(email string) (bool, error) {
	resp := Accounts{}
	result := repo.db.Model(&Accounts{}).Where("email = ?", email).First(&resp)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, result.Error
	}

	return true, nil
}

func (repo *accountCustomRepo) CreateAccount(account Accounts) (Accounts, error) {
	tx := repo.db.Begin()
	result := tx.Create(&account)
	if result.Error != nil {
		tx.Rollback()
		return Accounts{}, result.Error
	}
	tx.Commit()
	return account, nil
}

func (repo *accountCustomRepo) UpdateAccount(id uint64, fullName, email string) error {
	tx := repo.db.Begin()
	result := tx.Model(&Accounts{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"full_name":    fullName,
			"email":        email,
			"updated_date": time.Now(),
		})
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	if result.RowsAffected <= 0 {
		tx.Rollback()
		return errors.New("no data update")
	}

	tx.Commit()
	return nil
}

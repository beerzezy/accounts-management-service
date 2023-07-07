package account

import (
	"github.com/beerzezy/accounts-management-service/appconfig"
	"github.com/beerzezy/accounts-management-service/repositories"
)

type AccountService interface {
	Login(req RequestLoginAccount) (ResponseLoginAccount, error)
	CreateAccount(req RequestCreateAccount) (ResponseCreateAccount, error)
	UpdateAccount(req RequestUpdateAccount) error
	GetAccountById(id uint64) (ResponseGetAccount, error)
}

type accountServicer struct {
	cfg  appconfig.Config
	repo *repositories.Repo
}

func NewAccountService(cfg appconfig.Config, repo *repositories.Repo) AccountService {
	return &accountServicer{
		cfg,
		repo,
	}
}

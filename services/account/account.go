package account

import (
	"context"

	"github.com/beerzezy/accounts-management-service/repositories"
)

type AccountService interface {
	Login(ctx context.Context, req RequestLoginAccount) (ResponseLoginAccount, error)
	CreateAccount(ctx context.Context, req RequestCreateAccount) (ResponseCreateAccount, error)
	UpdateAccount(ctx context.Context, req RequestUpdateAccount) error
	GetAccountById(ctx context.Context, id uint64) (ResponseGetAccount, error)
}

type accountServicer struct {
	repo *repositories.Repo
}

func NewAccountService(repo *repositories.Repo) AccountService {
	return &accountServicer{
		repo,
	}
}

package account

import (
	"context"
	"fmt"
	"regexp"

	"github.com/beerzezy/accounts-management-service/repositories"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

func (srv *accountServicer) CreateAccount(ctx context.Context, req RequestCreateAccount) (ResponseCreateAccount, error) {

	if valid := validateEmail(req.Email); !valid {
		return ResponseCreateAccount{}, fmt.Errorf("email: %s is invalid", req.Email)
	}

	if valid := validatePassword(req.Password); !valid {
		return ResponseCreateAccount{}, fmt.Errorf("password: is invalid")
	}

	//validate email exiting
	isFound, err := srv.repo.AccountCustomRepo.ValidateAccountByEmail(req.Email)
	if err != nil {
		return ResponseCreateAccount{}, err
	}

	if isFound {
		return ResponseCreateAccount{}, fmt.Errorf("account: %s is already exists", req.Email)
	}

	result, err := srv.repo.AccountCustomRepo.CreateAccount(repositories.Accounts{
		FullName: req.FullName,
		Email:    req.Email,
		Password: hashPassword([]byte(req.Password)),
	})
	if err != nil {
		return ResponseCreateAccount{}, err
	}

	return ResponseCreateAccount{
		Id:       result.Id,
		FullName: result.FullName,
		Email:    result.Email,
	}, nil
}

func validateEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(email)
}

func validatePassword(password string) bool {
	return len(password) >= 8
}

func hashPassword(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Error(err)
	}

	return string(hash)
}

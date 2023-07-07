package account

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

type jwtCustomClaims struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
	jwt.RegisteredClaims
}

func (srv *accountServicer) Login(ctx context.Context, req RequestLoginAccount) (ResponseLoginAccount, error) {

	accountInfo, err := srv.repo.AccountCustomRepo.FindAccountByEmail(req.Email)
	if err != nil {
		return ResponseLoginAccount{}, err
	}

	matchEmail := false
	if req.Email == accountInfo.Email {
		matchEmail = true
	}

	matchPwd := comparePasswords(accountInfo.Password, []byte(req.Password))
	if !matchPwd || !matchEmail {
		return ResponseLoginAccount{}, fmt.Errorf("user or password invalid")
	}

	expiration, err := strconv.ParseUint(viper.GetString("jwt.expiration"), 10, 64)
	if err != nil {
		return ResponseLoginAccount{}, err
	}

	claims := &jwtCustomClaims{
		Id:   accountInfo.Id,
		Name: accountInfo.FullName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(msec(expiration))),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := viper.GetString("jwt.signingkey")
	signedJWT, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return ResponseLoginAccount{}, err
	}

	return ResponseLoginAccount{
		Id:       accountInfo.Id,
		FullName: accountInfo.FullName,
		Email:    accountInfo.Email,
		Token:    signedJWT,
	}, nil
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Error(err)
		return false
	}

	return true
}

func msec(t uint64) time.Duration {
	return time.Duration(t) * time.Millisecond
}

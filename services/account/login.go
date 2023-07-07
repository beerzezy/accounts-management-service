package account

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

type jwtCustomClaims struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
	jwt.RegisteredClaims
}

func (srv *accountServicer) Login(req RequestLoginAccount) (ResponseLoginAccount, error) {

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

	claims := &jwtCustomClaims{
		Id:   accountInfo.Id,
		Name: accountInfo.FullName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(msec(srv.cfg.JWT.Expiration))),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := srv.cfg.JWT.SigningKey
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

func msec(t uint32) time.Duration {
	return time.Duration(t) * time.Millisecond
}

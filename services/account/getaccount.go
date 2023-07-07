package account

import (
	"context"
)

func (srv *accountServicer) GetAccountById(ctx context.Context, id uint64) (ResponseGetAccount, error) {

	result, err := srv.repo.AccountCustomRepo.FindAccountById(id)
	if err != nil {
		return ResponseGetAccount{}, err
	}

	return ResponseGetAccount{
		Id:       result.Id,
		FullName: result.FullName,
		Email:    result.Email,
	}, nil
}

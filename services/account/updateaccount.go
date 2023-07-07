package account

import (
	"context"
)

func (srv *accountServicer) UpdateAccount(ctx context.Context, req RequestUpdateAccount) error {
	return srv.repo.AccountCustomRepo.UpdateAccount(req.Id, req.FullName, req.Email)
}

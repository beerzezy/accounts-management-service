package account

func (srv *accountServicer) UpdateAccount(req RequestUpdateAccount) error {
	return srv.repo.AccountCustomRepo.UpdateAccount(req.Id, req.FullName, req.Email)
}

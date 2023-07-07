package account

type RequestLoginAccount struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResponseLoginAccount struct {
	Id       uint64 `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

type RequestCreateAccount struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResponseCreateAccount struct {
	Id       uint64 `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

type RequestUpdateAccount struct {
	Id       uint64 `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

type ResponseGetAccount struct {
	Id       uint64 `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

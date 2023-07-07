package account

type RequestLoginAccount struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type ResponseLoginAccount struct {
	Id       uint64 `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

type RequestCreateAccount struct {
	FullName string `json:"full_name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type ResponseCreateAccount struct {
	Id       uint64 `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

type RequestUpdateAccount struct {
	Id       uint64 `json:"id" validate:"required"`
	FullName string `json:"full_name" validate:"required"`
	Email    string `json:"email" validate:"required"`
}

type ResponseGetAccount struct {
	Id       uint64 `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

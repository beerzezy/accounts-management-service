package errors

import "fmt"

type ServiceError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ResourceNotFoundError struct {
	ServiceError
}

type ValidationError struct {
	ServiceError
}

type UnauthorizedError struct {
	ServiceError
}

type CannotBindRequestStructError struct {
	ServiceError
}

type InvalidRequireFieldError struct {
	ServiceError
}

type DuplicateError struct {
	ServiceError
}

type HolderError struct {
	ServiceError
}
type AnnounceBondError struct {
	ServiceError
}
type InquiryBondTrasactiondError struct {
	ServiceError
}

func NewServiceError(code int, message string) ServiceError {
	return ServiceError{Code: code, Message: message}
}

func NewValidationError(code int, message string) ValidationError {
	return ValidationError{
		ServiceError: ServiceError{
			Code:    code,
			Message: message,
		},
	}
}

func NewUnauthorizedError(message string) UnauthorizedError {
	return UnauthorizedError{
		ServiceError: ServiceError{
			Code:    Unauthorized,
			Message: message,
		},
	}
}

func NewResourceNotFoundError() ResourceNotFoundError {
	return ResourceNotFoundError{
		ServiceError: ServiceError{
			Code:    ResourceNotFound,
			Message: "Resource not found",
		},
	}
}

func NewCannotBindRequestStructError() ServiceError {
	return ServiceError{
		Code:    CannotBindRequestStruct,
		Message: "Cannot Bind Request Struct",
	}
}

func NewInvalidRequireFieldError() ServiceError {
	return ServiceError{
		Code:    InvalidRequireField,
		Message: "Invalid Require Field",
	}
}

func NewDuplicateError() ServiceError {
	return ServiceError{
		Code:    Duplicate,
		Message: "Duplicate Data",
	}
}

func NewDuplicateWithMessageError(msg string) ServiceError {
	return ServiceError{
		Code:    Duplicate,
		Message: msg,
	}
}

func NewInvalidRequestWithMessageError(msg string) ServiceError {
	return ServiceError{
		Code:    InvalidRequest,
		Message: msg,
	}
}

func (e ServiceError) Error() string {
	return fmt.Sprintf("error {code=%d, message=%s}", e.Code, e.Message)
}

func NewExceptionError(msg string) ServiceError {
	return ServiceError{
		Code:    ExceptionError,
		Message: msg,
	}
}

func NewRecordNotFoundError() ServiceError {
	return ServiceError{
		Code:    RecordNotFoud,
		Message: "Record not found",
	}
}

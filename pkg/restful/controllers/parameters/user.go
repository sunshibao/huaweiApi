package parameters

import (
	"huaweiApi/pkg/utils/validator"
)

type UserRegisterRequest struct {
	Id       uint64 `json:"id"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

func (request *UserRegisterRequest) Validate() error {
	return validator.NewWrapper(
		validator.ValidateString(request.UserName, "userName", validator.ItemNotEmptyLimit, UserNameLengthLimit),
		validator.ValidateString(request.Email, "email", validator.ItemNotEmptyLimit, EmailLengthLimit),
		validator.ValidateString(request.Mobile, "mobile", validator.ItemNotEmptyLimit, MobileLengthLimit),
		validator.ValidateString(request.Password, "password", validator.ItemNotEmptyLimit, PasswordLengthLimit),
	).Validate()
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (request *UserLoginRequest) Validate() error {
	return validator.NewWrapper(
		validator.ValidateString(request.Email, "email", validator.ItemNotEmptyLimit, EmailLengthLimit),
		validator.ValidateString(request.Password, "password", validator.ItemNotEmptyLimit, PasswordLengthLimit),
	).Validate()
}

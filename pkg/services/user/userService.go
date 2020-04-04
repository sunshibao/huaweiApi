package user

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/jinzhu/gorm"

	"huaweiApi/pkg/config"
	userModel "huaweiApi/pkg/models/user"
	userRep "huaweiApi/pkg/repositorys/user"
	"huaweiApi/pkg/utils/middleware/auth"
)

const EmailAlreadyExist = 1
const EmailAlNoExist = 0

func UserRegister(user *userModel.Users) (emailStatus int, err error) {
	emailStatus = EmailAlNoExist
	_, err = userRep.GetUserInfo(user.Email)

	if !gorm.IsRecordNotFoundError(err) {
		emailStatus = EmailAlreadyExist
		return emailStatus, nil
	}

	h := md5.New()
	h.Write([]byte(user.Password))
	passwordMd5 := h.Sum(nil)
	user.Password = hex.EncodeToString(passwordMd5)
	err = userRep.UserRegister(user)
	if err != nil {
		return emailStatus, err
	}
	return emailStatus, nil

}
func UserLogin(email string, password string) (userResponse *userModel.Users, userToken string, err error) {
	userResponse, err = userRep.GetUserInfo(email)
	if !gorm.IsRecordNotFoundError(err) {
		h := md5.New()
		h.Write([]byte(password))
		passwordMd5 := h.Sum(nil)
		newPassword := hex.EncodeToString(passwordMd5)
		if userResponse.Password == newPassword {
			userToken, err := auth.NewUserJwtToken(userResponse.Id, nil, config.Config.RESTfulService.Auth.GetUserTokenKey())
			if err != nil {
				return nil, "", err
			}

			return userResponse, userToken, nil
		}
	}

	return nil, "", err
}

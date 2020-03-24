package user

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/jinzhu/gorm"

	userModel "huaweiApi/pkg/models/user"
	userRep "huaweiApi/pkg/repositorys/user"
)

func UserRegister(user *userModel.Users) (err error) {
	h := md5.New()
	h.Write([]byte(user.Password))
	passwordMd5 := h.Sum(nil)
	user.Password = hex.EncodeToString(passwordMd5)
	err = userRep.UserRegister(user)
	if err != nil {
		return err
	}
	return nil

}
func UserLogin(email string, password string) (userResponse *userModel.Users, err error) {
	userResponse, err = userRep.GetUserInfo(email)
	if err != nil {
		return nil, err
	}
	if !gorm.IsRecordNotFoundError(err) {
		h := md5.New()
		h.Write([]byte(password))
		passwordMd5 := h.Sum(nil)
		newPassword := hex.EncodeToString(passwordMd5)
		if userResponse.Password == newPassword {
			return userResponse, nil
		}
	}
	return nil, err

}

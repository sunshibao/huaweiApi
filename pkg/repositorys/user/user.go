package user

import (
	"github.com/sunshibao/connection"

	userModel "huaweiApi/pkg/models/user"
	"huaweiApi/pkg/utils/idcreator"
)

func UserRegister(user *userModel.Users) (err error) {
	user.Id = idcreator.NextID()
	err = connection.GetMySQL().Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUserInfo(email string) (user *userModel.Users, err error) {
	userSetting := new(userModel.Users)
	err = connection.GetMySQL().Where("email=?", email).First(userSetting).Error
	if err != nil {
		return nil, err
	}
	return userSetting, nil
}

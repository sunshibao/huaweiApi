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

func UpdateUser(user *userModel.Users) (err error) {
	userSetting := new(userModel.Users)
	err = connection.GetMySQL().Model(userSetting).Debug().Where("id=?", user.Id).Update(user).Error
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

func DeductionGold(id uint64, gold int64) (err error) {
	userSetting := new(userModel.Users)
	userSetting.Gold = gold
	err = connection.GetMySQL().Debug().Model(userSetting).Where("id = ?", id).Update(userSetting).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUserInfoById(id uint64) (user *userModel.Users, err error) {
	userSetting := new(userModel.Users)
	err = connection.GetMySQL().Where("id=?", id).First(userSetting).Error
	if err != nil {
		return nil, err
	}
	return userSetting, nil
}

func AddDeductionRecord(deductionRecord *userModel.DeductionRecord) (err error) {
	deductionRecord.Id = idcreator.NextID()
	err = connection.GetMySQL().Debug().Create(deductionRecord).Error
	if err != nil {
		return err
	}
	return nil
}


package user

import (
	"huaweiApi/pkg/databases"
)

type Users struct {
	databases.BaseModel
	UserName string `gorm:"column:userName;type:varchar(64);default:'';comment:'用户名称'"`
	Email    string `gorm:"column:email;type:varchar(32);default:'';comment:'邮箱'"`
	Mobile   string `gorm:"column:mobile;type:varchar(11);default:'';comment:'手机号'"`
	Password string `gorm:"column:password;type:varchar(40);default:'';comment:'密码'"`
	Gold     int64  `gorm:"column:gold;default:0;comment:'金币'"`
}

func (Users) TableName() string {
	return "users"
}

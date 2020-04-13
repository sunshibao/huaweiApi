package user

import (
	"huaweiApi/pkg/databases"
)

type DeductionRecord struct {
	databases.BaseModel
	UserId uint64 `gorm:"column:userId;default:0;comment:'用户Id'"`
	Gold   int64  `gorm:"column:gold;default:0;comment:'金币'"`
}

func (DeductionRecord) TableName() string {
	return "deduction_record"
}

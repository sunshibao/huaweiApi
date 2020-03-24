package databases

import (
	"time"
)

type BaseModel struct {
	Id        uint64    `gorm:"column:id;primary_key;auto_increment:false;comment:'主键ID'"`
	CreatedAt time.Time `gorm:"column:createdAt;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:'创建时间'"`
	UpdatedAt time.Time `gorm:"column:updatedAt;type:datetime;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:'更新时间'"`
}

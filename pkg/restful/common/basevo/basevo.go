package basevo

import (
	"fmt"
	"time"
)

type BaseVo struct {
	ID        string    `json:"id"`
	CreatedBy string    `json:"createdBy"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedBy string    `json:"updatedBy"`
	UpdatedAt time.Time `json:"updatedAt"`
}

//暂时不用，还是使用默认的RFC3339
type JsonTime time.Time

func (t *JsonTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(*t).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

func (t *JsonTime) UnmarshalJSON(data []byte) (err error) {
	if data == nil || len(data) == 2 {
		return
	}
	parse, err := time.ParseInLocation("\"2006-01-02 15:04:05\"", string(data), time.Local)
	*t = JsonTime(parse)
	return
}

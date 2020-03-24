package idcreator

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/sony/sonyflake"
)

const (
	DefaultServiceID = 0
)

var idCreator *sonyflake.Sonyflake

func init() {

	InitCreator(DefaultServiceID)
}

func InitCreator(serviceId uint16) {

	idCreator = sonyflake.NewSonyflake(
		sonyflake.Settings{
			MachineID: func() (u uint16, e error) {
				return serviceId, nil
			},
		})
}

func NextID() uint64 {
	id, err := idCreator.NextID()
	if err != nil {
		// Based on the readme from Sonyflake: NextID can continue to generate IDs for about 174 years from StartTime.
		// After that time, an error will return. In our case, we can ignore this error.
		// So we eat the error here.
		logrus.Error(err)
	}

	return id
}

func NextString() string {
	return fmt.Sprintf("%x", NextID())
}

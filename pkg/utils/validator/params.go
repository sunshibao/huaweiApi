package validator

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"huaweiApi/pkg/constants"
	"huaweiApi/pkg/utils/h"
)

func Params(c *gin.Context, v Validator) error {

	if err := c.ShouldBindJSON(v); err != nil {
		return h.EBindBodyError.WithError(err)
	}

	reqId, _ := c.Value(constants.RequestID).(string)
	logrus.Debugf("reqId: %s, v = %+v\n", reqId, v)
	if err := v.Validate(); err != nil {
		return h.EParamsError.WithError(err)
	}
	return nil
}

package auth

import (
	"github.com/gin-gonic/gin"

	"huaweiApi/pkg/constants"
)

func GetUserId(c *gin.Context) (u64 uint64) {
	if val, ok := c.Get(constants.UserID); ok && val != nil {
		u64, _ = val.(uint64)
	}
	return
}

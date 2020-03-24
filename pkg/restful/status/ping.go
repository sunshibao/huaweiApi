package status

import (
	"github.com/gin-gonic/gin"

	"huaweiApi/pkg/utils/h"
)

// @ID Ping
// @Summary Return pong when services is normal
// @Description Check RESTful services port status
// @Tags ping
// @Produce text/plain
// @Success 200 {string} string
// @Router /api/v1/ping [get]
func Ping(c *gin.Context) {

	h.R(c, "pong")
}

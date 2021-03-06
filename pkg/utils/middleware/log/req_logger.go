package log

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"huaweiApi/pkg/constants"
)

func genReqId() string {
	var b [12]byte
	_, err := io.ReadFull(rand.Reader, b[:])
	if err != nil {
		//基本不会出现
		logrus.WithError(err).Warn("generate request id failed")
		return "fail_gen_req_id"
	}
	return base64.URLEncoding.EncodeToString(b[:])
}

func ReqLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		logrus.SetFormatter(&logrus.TextFormatter{TimestampFormat: time.RFC3339})

		reqId := c.Request.Header.Get(constants.RequestID)
		if reqId == "" {
			reqId = genReqId()
			c.Request.Header.Set(constants.RequestID, reqId)
		}
		c.Set(constants.RequestID, reqId)
		// Set request id into response header
		c.Writer.Header().Set(constants.RequestID, reqId)

		c.Next()

		end := time.Now()
		latency := end.Sub(start)

		entry := logrus.WithFields(logrus.Fields{
			"reqId":      reqId,
			"status":     c.Writer.Status(),
			"method":     c.Request.Method,
			"path":       c.Request.URL,
			"size":       c.Writer.Size(),
			"ip":         c.ClientIP(),
			"latency":    latency,
			"user-agent": c.Request.UserAgent(),
		})

		if len(c.Errors) > 0 {
			entry.Info(c.Errors.String())
		} else {
			entry.Info()
		}
	}
}

// usage: ReqEntry(c).Debug(".....")
func ReqEntry(c context.Context) *logrus.Entry {
	reqIdVal := c.Value(constants.RequestID)
	if reqIdVal != nil {
		reqId, _ := reqIdVal.(string)
		return logrus.WithField("reqId", reqId)
	} else {
		return logrus.WithField("reqId", "unknown")
	}
}

package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func WithPath(handlerFunc gin.HandlerFunc, urlPrefixs ...string) gin.HandlerFunc {
	return doWithPath(handlerFunc, true, urlPrefixs...)
}

func WithoutPath(handlerFunc gin.HandlerFunc, urlPrefixs ...string) gin.HandlerFunc {
	return doWithPath(handlerFunc, false, urlPrefixs...)
}

//目前只需要个简单的实现
func doWithPath(handlerFunc gin.HandlerFunc, doWithMatch bool, urlPrefixs ...string) func(c *gin.Context) {
	return func(c *gin.Context) {
		path := c.Request.URL.Path

		if !strings.HasSuffix(path, "/") {
			path = path + "/"
		}

		pathLen := len(path)

		matched := false
		for _, urlPrefix := range urlPrefixs {
			if !strings.HasSuffix(urlPrefix, "/") {
				urlPrefix = urlPrefix + "/"
			}
			if prefixLen := len(urlPrefix); pathLen >= prefixLen && path[:prefixLen] == urlPrefix {
				matched = true
				break
			}
		}

		if doWithMatch {
			if matched {
				handlerFunc(c)
			} else {
				c.Next()
			}
		} else {
			if matched {
				c.Next()
			} else {
				handlerFunc(c)
			}
		}
	}
}

package auth

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/shomali11/util/xstrings"

	"huaweiApi/pkg/constants"
	"huaweiApi/pkg/utils/middleware/log"
	"huaweiApi/pkg/utils/h"
	"huaweiApi/pkg/utils/idcreator"
)

/* gin用户jwt认证中间件 */
func UserJwtAuthentication(tokenKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenVal := c.GetHeader(constants.RequestUserToken)
		if xstrings.IsBlank(tokenVal) {
			c.Abort()
			h.TokenInvalid(c)
			return
		}

		token, err := jwt.Parse(tokenVal, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				log.ReqEntry(c).WithField("token", tokenVal).WithField("sign_method", token.Method).Errorf("invalid token sign method")
				return nil, fmt.Errorf("Unexpected sign method: %v", token.Header["alg"])
			}
			return []byte(tokenKey), nil
		})

		if err != nil {
			c.Abort()
			validationErr, ok := err.(*jwt.ValidationError)
			if ok && (validationErr.Errors&jwt.ValidationErrorExpired != 0) && (validationErr.Errors&jwt.ValidationErrorSignatureInvalid == 0) {
				h.TokenExpired(c)
				return
			}
			log.ReqEntry(c).WithField("token", tokenVal).WithError(err).Errorf("token is invalid")
			h.TokenInvalid(c)
			return
		}

		if !token.Valid {
			log.ReqEntry(c).WithField("token", tokenVal).Errorf("token is invalid")
			c.Abort()
			h.TokenInvalid(c)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			log.ReqEntry(c).WithField("token", tokenVal).Errorf("token is invalid")
			c.Abort()
			h.TokenInvalid(c)
			return
		}

		subject, ok := claims["sub"]

		if !ok {
			log.ReqEntry(c).WithField("token", tokenVal).Errorf("token is invalid")
			c.Abort()
			h.TokenInvalid(c)
			return
		}

		subjectStr, ok := subject.(string)

		if !ok {
			log.ReqEntry(c).WithField("token", tokenVal).Errorf("token is invalid")
			c.Abort()
			h.TokenInvalid(c)
			return
		}

		userId, err := strconv.ParseUint(subjectStr, 10, 64)

		if err != nil {
			log.ReqEntry(c).WithField("token", tokenVal).WithError(err).Errorf("token is invalid")
			c.Abort()
			h.TokenInvalid(c)
			return
		}

		c.Set(constants.UserID, userId)
		c.Next()
	}
}

func NewUserJwtToken(userId uint64, info map[string]string, tokenKey string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["jti"] = strconv.FormatUint(idcreator.NextID(), 10)
	claims["iss"] = "wanxin"
	claims["sub"] = strconv.FormatUint(userId, 10)
	claims["exp"] = time.Now().AddDate(0, 0, 1).Unix()
	if info != nil {
		infoBytes, err := json.Marshal(info)
		if err != nil {
			return "", err
		}
		claims["info"] = string(infoBytes)
	}
	token.Claims = claims
	return token.SignedString([]byte(tokenKey))
}

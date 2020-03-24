package h

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type H = map[string]interface{}

type ResponseData struct {
	Status  int         `json:"-"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"msg,omitempty"`
}

func (data ResponseData) Error() string {

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		logrus.WithField("error", err).Info("ResponseData.Error()")
		// 这里错误类型是定义好的，正常不应该出现Marshal错误，并且对这个情况，现在用单测覆盖这部分输出，保证输出正常。
		jsonBytes, _ = json.Marshal(EUnknown)
	}
	return string(jsonBytes)
}

func (data ResponseData) String() string {

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		logrus.WithField("error", err).Info("ResponseData.String():", err)
		return ""
	}
	return string(jsonBytes)
}

func (data ResponseData) WithCode(code int) *ResponseData {

	data.Code = code
	return &data
}

func (data ResponseData) WithError(err error) *ResponseData {

	switch err := err.(type) {
	case *ResponseData:
		data.Code = err.Code
		data.Message = err.Message
	default:
		data.Message = err.Error()
	}

	return &data
}

func (data ResponseData) WithMessage(message string) *ResponseData {

	data.Message = message
	return &data
}

var (
	ENotFound      = &ResponseData{http.StatusOK, 404, nil, "NotFound"}
	EParamsError   = &ResponseData{http.StatusOK, 400, nil, "ParamsError"}
	EBindBodyError = &ResponseData{http.StatusOK, 400, nil, "BindBodyError"}
	EUnknown       = &ResponseData{http.StatusOK, 500, nil, "Unknown"}
	EExists        = &ResponseData{http.StatusOK, 409, nil, "Exists"}
)

func E(c *gin.Context, httpStatus, errCode int, message string) {
	c.JSON(httpStatus, ResponseData{
		Code:    errCode,
		Message: message,
	})
}

func InternalErr(c *gin.Context, errCode int, message string) {
	E(c, http.StatusOK, errCode, message)
}

func Data(c *gin.Context, data interface{}) {

	R(c, ResponseData{
		Data: data,
	})
}

func R(c *gin.Context, body interface{}) {
	if c.Request.Method == "POST" {
		c.JSON(http.StatusCreated, body)
		return
	}
	if c.Request.Method == "DELETE" {
		c.JSON(http.StatusNoContent, body)
		return
	}
	c.JSON(http.StatusOK, body)
}

func RJsonP(c *gin.Context, body interface{}) {
	if c.Request.Method == "POST" {
		c.JSONP(http.StatusCreated, body)
		return
	}
	if c.Request.Method == "DELETE" {
		c.JSONP(http.StatusNoContent, body)
		return
	}
	c.JSONP(http.StatusOK, body)
}

func allowAccessOtherSites(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Credentials", "true")
}

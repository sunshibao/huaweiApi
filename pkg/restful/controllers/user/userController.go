package user

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"huaweiApi/pkg/models/user"
	userModel "huaweiApi/pkg/models/user"
	"huaweiApi/pkg/restful/controllers/parameters"
	"huaweiApi/pkg/restful/errorcode"
	"huaweiApi/pkg/restful/returncode"
	userService "huaweiApi/pkg/services/user"
	"huaweiApi/pkg/utils/h"
	"huaweiApi/pkg/utils/log"
	"huaweiApi/pkg/utils/middleware/auth"
	"huaweiApi/pkg/utils/validator"
)

func UserRegister(c *gin.Context) {

	var (
		err error
	)
	requestData, hasError := userRegisterRequestData(c)
	if hasError {
		return
	}

	newUser := changeToUserRegister(requestData)

	emailStatus, err := userService.UserRegister(newUser)

	if err != nil {
		h.InternalErr(c, errorcode.CommonError, errorcode.StatusText(errorcode.CommonError))
		return
	}
	if emailStatus == 1 {
		h.InternalErr(c, errorcode.EmailAlreadyExistError, errorcode.StatusText(errorcode.EmailAlreadyExistError))
		return
	}
	h.Data(c, returncode.SuccessfulOption{Success: true})
}

func UserLogin(c *gin.Context) {

	var (
		err error
	)
	requestData, hasError := userLoginRequestData(c)
	if hasError {
		return
	}
	newUser := changeToUserLogin(requestData)

	userResponse, userToken, err := userService.UserLogin(newUser.Email, newUser.Password)

	if err != nil {
		h.InternalErr(c, errorcode.CommonError, errorcode.StatusText(errorcode.CommonError))
		return
	}

	h.Data(c, toUserLoginVo(userResponse, userToken))
}

func DeductionGold(c *gin.Context) {
	userId := auth.GetUserId(c)
	requestData, hasError := goldRequestData(c)
	if hasError {
		return
	}
	gold := requestData.Gold
	goldType := requestData.Type

	status, err := userService.DeductionGold(userId, gold, goldType)

	if err != nil {
		h.InternalErr(c, errorcode.NullDataError, errorcode.StatusText(errorcode.NullDataError))
		return
	}

	if status == 2 {
		h.InternalErr(c, errorcode.BalanceDeficiencyError, errorcode.StatusText(errorcode.BalanceDeficiencyError))
	}
	h.Data(c, returncode.SuccessfulOption{Success: true})
}

func userRegisterRequestData(c *gin.Context) (requestData *parameters.UserRegisterRequest, hasError bool) {

	var err error
	requestData = new(parameters.UserRegisterRequest)
	logger := log.ReqEntry(c)

	if err = validator.Params(c, requestData); err != nil {
		logger.WithField("action", "parameter json parse").Info(err)
		h.InternalErr(c, errorcode.JSONParseError, errorcode.StatusText(errorcode.JSONParseError))
		return nil, true
	}

	logger.WithField("data", requestData).Debug("get create user data")
	return requestData, false
}

func userLoginRequestData(c *gin.Context) (requestData *parameters.UserLoginRequest, hasError bool) {

	var err error
	requestData = new(parameters.UserLoginRequest)
	logger := log.ReqEntry(c)

	if err = validator.Params(c, requestData); err != nil {
		logger.WithField("action", "parameter json parse").Info(err)
		h.InternalErr(c, errorcode.JSONParseError, errorcode.StatusText(errorcode.JSONParseError))
		return nil, true
	}

	logger.WithField("data", requestData).Debug("get create user data")
	return requestData, false
}

func goldRequestData(c *gin.Context) (requestData *parameters.GoldRequest, hasError bool) {

	var err error
	requestData = new(parameters.GoldRequest)
	logger := log.ReqEntry(c)

	if err = validator.Params(c, requestData); err != nil {
		logger.WithField("action", "parameter json parse").Info(err)
		h.InternalErr(c, errorcode.JSONParseError, errorcode.StatusText(errorcode.JSONParseError))
		return nil, true
	}

	logger.WithField("data", requestData).Debug("gold Request Data")
	return requestData, false
}

//类型转换
func changeToUserRegister(NewUser *parameters.UserRegisterRequest) *user.Users {

	return &user.Users{
		UserName: NewUser.UserName,
		Email:    NewUser.Email,
		Mobile:   NewUser.Mobile,
		Password: NewUser.Password,
	}
}

//类型转换
func changeToUserLogin(NewUser *parameters.UserLoginRequest) *user.Users {

	return &user.Users{
		Email:    NewUser.Email,
		Password: NewUser.Password,
	}
}

type UserLoginVo struct {
	Id        string `json:"id"`
	UserName  string `json:"userName"`
	Email     string `json:"email"`
	Mobile    string `json:"mobile"`
	Gold      int64  `json:"gold"`
	UserToken string `json:"userToken"`
}

func toUserLoginVo(user *userModel.Users, userToken string) *UserLoginVo {
	return &UserLoginVo{
		Id:        strconv.FormatUint(user.Id, 10),
		UserName:  user.UserName,
		Mobile:    user.Mobile,
		Email:     user.Email,
		Gold:      user.Gold,
		UserToken: userToken,
	}
}

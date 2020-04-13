package application

import (
	"github.com/gin-gonic/gin"

	"huaweiApi/pkg/config"
	huaweiApiController "huaweiApi/pkg/restful/controllers/huaweiApi"
	userController "huaweiApi/pkg/restful/controllers/user"
	"huaweiApi/pkg/utils/middleware"
	"huaweiApi/pkg/utils/middleware/auth"

	_ "huaweiApi/pkg/restful/swaggerdocs"
)

func (a *app) setRESTfulRoutes() {
	restfulSetting := config.Config.RESTfulService

	prefixGroup := a.httpHandler.Group("/huaweiapi")
	v1 := prefixGroup.Group("/v1")

	v1.Use(middleware.WithoutPath(
		auth.UserJwtAuthentication(restfulSetting.Auth.GetUserTokenKey()),
		v1.BasePath()+"/user/register",
		v1.BasePath()+"/user/login",

	))

	aggregatorGroup := v1.Group("/aggregator")
	userGroup := v1.Group("/user")

	a.setAggregatorGouters(aggregatorGroup)
	a.setUserGouters(userGroup)

}

// 华为接口
func (a *app) setAggregatorGouters(aggregator *gin.RouterGroup) {
	aggregator.POST("/createPayment", huaweiApiController.CreatePayment)
	aggregator.POST("/createSubscription", huaweiApiController.CreateSubscription)
	aggregator.POST("/syncPayment", huaweiApiController.SyncPayment)
	aggregator.GET("/getPaymentInfo/:paymentID", huaweiApiController.GetPaymentInfo)

}

func (a *app) setUserGouters(user *gin.RouterGroup) {
	user.POST("/register", userController.UserRegister)
	user.POST("/login", userController.UserLogin)
	user.POST("/deductionGold", userController.DeductionGold)

}

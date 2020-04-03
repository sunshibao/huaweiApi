package application

import (
	"github.com/gin-gonic/gin"

	huaweiApiController "huaweiApi/pkg/restful/controllers/huaweiApi"
	userController "huaweiApi/pkg/restful/controllers/user"

	_ "huaweiApi/pkg/restful/swaggerdocs"
)

func (a *app) setRESTfulRoutes() {
	prefixGroup := a.httpHandler.Group("/huaweiapi")
	v1 := prefixGroup.Group("/v1")

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
}

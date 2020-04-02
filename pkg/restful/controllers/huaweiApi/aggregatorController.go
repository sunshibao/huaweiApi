package huaweiApi

import (
	"github.com/gin-gonic/gin"

	"huaweiApi/pkg/services/aggregator"
	"huaweiApi/pkg/utils/h"
)

type CreatePaymentRequest struct {
	Msisdn       string `json:"msisdn"`
	OperatorCode string `json:"operatorCode"`
	ProductID    string `json:"productId"`
	ExtRef       string `json:"extRef"`
	ExtInfos     []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"extInfos"`
}

func CreatePayment(c *gin.Context) {

	paymentReply, err := aggregator.CreatePayment()
	if err != nil {
		return
	}
	h.Data(c, paymentReply)
}
func CreateSubscription(c *gin.Context) {

	msisdn:=c.Param("msisdn")
	productId:=c.Param("productId")
	extRef:=c.Param("msisdn")

	paymentReply, err := aggregator.CreateSubscription(msisdn, productId, extRef)
	if err != nil {
		return
	}
	h.Data(c, paymentReply)
}

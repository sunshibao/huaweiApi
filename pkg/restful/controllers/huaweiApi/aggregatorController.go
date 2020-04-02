package huaweiApi

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"huaweiApi/pkg/models/huawei"
	"huaweiApi/pkg/restful/controllers/parameters"
	"huaweiApi/pkg/restful/errorcode"
	"huaweiApi/pkg/restful/returncode"
	"huaweiApi/pkg/services/aggregator"
	"huaweiApi/pkg/utils/h"
	"huaweiApi/pkg/utils/log"
	"huaweiApi/pkg/utils/validator"
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

	msisdn := c.Param("msisdn")
	productId := c.Param("productId")
	extRef := c.Param("msisdn")

	paymentReply, err := aggregator.CreateSubscription(msisdn, productId, extRef)
	if err != nil {
		return
	}
	h.Data(c, paymentReply)
}

func SyncPayment(c *gin.Context) {

	requestData, hasError := getSyncPaymentRequestData(c)
	if hasError {
		return
	}
	record := changeAddPaymentRecord(requestData)

	err := aggregator.AddPaymentRecord(record)

	if err != nil {
		return
	}

	h.Data(c, returncode.SuccessfulOption{Success: true})
}

func getSyncPaymentRequestData(c *gin.Context) (requestData *parameters.SyncPaymentRequest, hasError bool) {

	var err error
	requestData = new(parameters.SyncPaymentRequest)
	logger := log.ReqEntry(c)

	if err = validator.Params(c, requestData); err != nil {
		logger.WithField("action", "parameter json parse").Info(err)
		h.InternalErr(c, errorcode.JSONParseError, errorcode.StatusText(errorcode.JSONParseError))
		return nil, true
	}

	logger.WithField("data", fmt.Sprintf("%#v", requestData)).Debug("getSyncPaymentRequestData")
	return requestData, false
}

func changeAddPaymentRecord(requestData *parameters.SyncPaymentRequest) *huawei.PaymentRecord {
	return &huawei.PaymentRecord{
		PaymentID:     requestData.PaymentID,
		Msisdn:        requestData.Msisdn,
		ProductID:     requestData.ProductID,
		ExtRef:        requestData.ExtRef,
		Status:        uint64(requestData.Status),
		Amount:        uint64(requestData.Amount),
		SubTime:       requestData.SubTime,
		StartTime:     requestData.SubTime,
		EndTime:       requestData.EndTime,
		SvcName:       requestData.PaymentExt.SvcName,
		ChannelName:   requestData.PaymentExt.ChannelName,
		RenewalType:   requestData.PaymentExt.RenewalType,
		BillingRate:   uint64(requestData.PaymentExt.BillingRate),
		BillingCycle:  requestData.PaymentExt.BillingCycle,
		UpdatedAt:     requestData.PaymentExt.UpdatedAt,
		LastBilledAt:  requestData.PaymentExt.LastBilledAt,
		NextBillingAt: requestData.PaymentExt.NextBillingAt,
	}
}

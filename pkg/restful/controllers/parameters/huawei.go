package parameters

import (
	"huaweiApi/pkg/utils/validator"
)

type SyncPaymentRequest struct {
	PaymentID  string `json:"paymentId"`
	Msisdn     string `json:"msisdn"`
	ProductID  string `json:"productId"`
	ExtRef     string `json:"extRef"`
	Status     int    `json:"status"`
	Amount     int    `json:"amount"`
	SubTime    string `json:"subTime"`
	StartTime  string `json:"startTime"`
	EndTime    string `json:"endTime"`
	PaymentExt struct {
		SvcName       string `json:"svcName"`
		ChannelName   string `json:"channelName"`
		RenewalType   string `json:"renewalType"`
		BillingRate   int    `json:"billingRate"`
		BillingCycle  string `json:"billingCycle"`
		UpdatedAt     string `json:"updatedAt"`
		LastBilledAt  string `json:"lastBilledAt"`
		NextBillingAt string `json:"nextBillingAt"`
	} `json:"paymentExt"`
}

func (request *SyncPaymentRequest) Validate() error {
	return validator.NewWrapper().Validate()
}

type CreateSubscriptionRequest struct {
	Msisdn     string `json:"msisdn"`
	ProductID  string `json:"productId"`
	ExtRef     string `json:"extRef"`
}

func (request *CreateSubscriptionRequest) Validate() error {
	return validator.NewWrapper().Validate()
}
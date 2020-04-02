package huawei

import (
	"github.com/sunshibao/connection"

	"huaweiApi/pkg/models/huawei"
	"huaweiApi/pkg/utils/idcreator"
)

func AddPaymentRecord(paymentRecord *huawei.PaymentRecord) (err error) {
	paymentRecord.Id = idcreator.NextID()
	err = connection.GetMySQL().Create(paymentRecord).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdatePaymentRecord(paymentRecord *huawei.PaymentRecord) (err error) {

	err = connection.GetMySQL().Model(paymentRecord).Where("id = ?", paymentRecord.Id).Update(paymentRecord).Error
	if err != nil {
		return err
	}
	return nil
}

func GetPaymentRecordByPaymentId(paymentID string) (paymentRecord *huawei.PaymentRecord, err error) {
	paymentRecordSetting := new(huawei.PaymentRecord)
	err = connection.GetMySQL().Debug().Where("paymentID = ?", paymentID).First(paymentRecordSetting).Error
	return paymentRecordSetting, err
}

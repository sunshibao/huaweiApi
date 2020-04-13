package aggregator

import (
	"crypto/sha256"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/jinzhu/gorm"

	"huaweiApi/pkg/models/huawei"
	userModel "huaweiApi/pkg/models/user"
	huaweiRep "huaweiApi/pkg/repositorys/huawei"
	userRep "huaweiApi/pkg/repositorys/user"
)

type aggregatorService struct {
}

type (
	UrlParamView struct {
		Msisdn       string              `json:"msisdn"`
		OperatorCode string              `json:"operatorCode"`
		ProductId    string              `json:"productId"`
		ExtRef       string              `json:"extRef"`
		ExtInfos     []map[string]string `json:"extInfos"`
	}
)

type CreatePaymentReply struct {
	Code                 string `json:"code"`
	Description          string `json:"description"`
	ContractSubscription struct {
		SubscriptionID  string    `json:"subscriptionId"`
		ProductID       string    `json:"productId"`
		Status          int       `json:"status"`
		Amount          int       `json:"amount"`
		ExtRef          string    `json:"extRef"`
		Msisdn          string    `json:"msisdn"`
		SubTime         time.Time `json:"subTime"`
		StartTime       time.Time `json:"startTime"`
		EndTime         time.Time `json:"endTime"`
		SubscriptionExt struct {
			SvcName       string    `json:"svcName"`
			BillingCycle  string    `json:"billingCycle"`
			NextBillingAt time.Time `json:"nextBillingAt"`
			BillingRate   string    `json:"billingRate"`
			ChannelName   string    `json:"channelName"`
			RenewalType   string    `json:"renewalType"`
			LastBilledAt  time.Time `json:"lastBilledAt"`
			UpdatedAt     time.Time `json:"updatedAt"`
		} `json:"paymentExt"`
	} `json:"contractSubscription"`
}

type CreateSubscriptionReply struct {
	ContractSubscription struct {
		SubscriptionID  string    `json:"subscriptionId"`
		ProductID       string    `json:"productId"`
		SubTime         time.Time `json:"subTime"`
		SubscriptionExt struct {
			SvcName       string    `json:"svcName"`
			BillingCycle  string    `json:"billingCycle"`
			NextBillingAt time.Time `json:"nextBillingAt"`
			BillingRate   string    `json:"billingRate"`
			ChannelName   string    `json:"channelName"`
			RenewalType   string    `json:"renewalType"`
			LastBilledAt  time.Time `json:"lastBilledAt"`
			UpdatedAt     time.Time `json:"updatedAt"`
		} `json:"subscriptionExt"`
		Status  int       `json:"status"`
		EndTime time.Time `json:"endTime"`
		ExtRef  string    `json:"extRef"`
		Msisdn  string    `json:"msisdn"`
	} `json:"contractSubscription"`
	Description string `json:"description"`
	Code        string `json:"code"`
}

func CreatePayment(msisdn string, productId string, extRef string) (paymentReply *CreatePaymentReply, err error) {

	urlParam := &UrlParamView{
		Msisdn:       msisdn,
		OperatorCode: "70201",
		ProductId:    productId,
		ExtRef:       extRef,
		ExtInfos:     make([]map[string]string, 0),
	}
	urlParam.ExtInfos = append(urlParam.ExtInfos,
		map[string]string{
			"key":   "channel",
			"value": "SMS",
		},
		map[string]string{
			"key":   "doi_channel",
			"value": "SMS",
		},
	)

	jsonBody, err := json.Marshal(urlParam)

	if err != nil {
		return
	}
	body := strings.NewReader(string(jsonBody))

	bytes, err := HttpPostRequest(body, "https://159.138.167.235:17131/apiaccess/ita/createPayment/v1")

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(bytes, &paymentReply); err != nil {
		return nil, err

	}

	return paymentReply, nil
}

func CreateSubscription(msisdn string, productId string, extRef string) (subscriptionReply *CreateSubscriptionReply, err error) {
	urlParam := &UrlParamView{
		Msisdn:       msisdn,
		OperatorCode: "70201",
		ProductId:    productId,
		ExtRef:       extRef,
		ExtInfos:     make([]map[string]string, 0),
	}
	urlParam.ExtInfos = append(urlParam.ExtInfos,
		map[string]string{
			"key":   "channel",
			"value": "SMS",
		},
		map[string]string{
			"key":   "doi_channel",
			"value": "SMS",
		},
	)

	jsonBody, err := json.Marshal(urlParam)

	if err != nil {
		return
	}
	body := strings.NewReader(string(jsonBody))

	bytes, err := HttpPostRequest(body, "https://159.138.167.235:17131/apiaccess/ita/createSubscription/v1")

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(bytes, &subscriptionReply); err != nil {
		return nil, err
	}
	return subscriptionReply, nil
}

func HttpPostRequest(body *strings.Reader, url string) (bytes []byte, err error) {

	nonce := "66C92B11FF8A425FB8D4CCFE"
	username := "67ef6805e6004cce9cce24b7f13767c6"
	password := "1e88f55b4e034b1783e2686ec9dfbffd"
	now := time.Now()
	time8, _ := time.ParseDuration("-1h")
	created := now.Add(8 * time8).Format("2006-01-02T15:04:05Z")
	nonce = base64.StdEncoding.EncodeToString([]byte(nonce))
	signRawString := nonce + created + password
	h := sha256.New()
	h.Write([]byte(signRawString))
	passwordDigest := base64.StdEncoding.EncodeToString(h.Sum(nil))

	req, err := http.NewRequest("POST", url, body)
	authorization := `"WSSE realm="SDP", profile="UsernameToken", type="Appkey"`
	wsse := "UsernameToken Username=\"" + username + "\",PasswordDigest=\"" + passwordDigest + "\",Nonce=\"" + nonce + "\",Created=\"" + created + "\""
	contentType := "Content-Type: application/json; charset=UTF-8"
	req.Header.Set("Authorization", authorization)
	req.Header.Set("X-WSSE", wsse)
	req.Header.Set("Content-Type", contentType)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	respClient, _ := client.Do(req)

	bytes, err = ioutil.ReadAll(respClient.Body)

	if err != nil {
		return nil, err
	}
	return bytes, nil

}

func AddPaymentRecord(paymentRecord *huawei.PaymentRecord, userId uint64) (err error) {

	_, err = huaweiRep.GetPaymentRecordByPaymentId(paymentRecord.PaymentID)

	if gorm.IsRecordNotFoundError(err) {
		err = huaweiRep.AddPaymentRecord(paymentRecord)
	} else {
		err = huaweiRep.UpdatePaymentRecord(paymentRecord)
	}

	if err != nil {
		return err
	}
	if paymentRecord.Status == 2 {
		user := new(userModel.Users)

		user.Id = userId

		users, err := userRep.GetUserInfoById(userId)
		if err != nil {
			return err
		}
		user.Gold = users.Gold + 4000
		//充值
		err = userRep.AddUserGold(user)
		if err != nil {
			return err
		}
	}

	return nil
}

func GetPaymentRecordByPaymentId(paymentID string) (paymentRecord *huawei.PaymentRecord, err error) {
	paymentRecord, err = huaweiRep.GetPaymentRecordByPaymentId(paymentID)
	if err != nil {
		return nil, err
	}
	return paymentRecord, nil

}

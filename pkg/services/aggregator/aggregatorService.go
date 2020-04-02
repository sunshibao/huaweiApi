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

func CreatePayment() (paymentReply *CreatePaymentReply, err error) {

	urlParam := &UrlParamView{
		Msisdn:       "27785385741",
		OperatorCode: "70201",
		ProductId:    "948d7b50-9f3f-4a0e-9a28-d73c15ba5141",
		ExtRef:       "1e88f55b4e034b1783e2686ec9dfbffd",
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

func CreateSubscription(msisdn string, productId string, extRef string) (paymentReply *CreatePaymentReply, err error) {

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

	if err := json.Unmarshal(bytes, &paymentReply); err != nil {
		return nil, err
	}
	return paymentReply, nil
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

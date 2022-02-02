package smsutils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dongri/phonenumber"
	gohttp "github.com/ochom/go-http"
)

var token = os.Getenv("ELEZA_SMS_TOKEN")

//SendSMS send sms to "http://api.eleza.online/v1/sms/send/"
func SendSMS(phoneNumber, message string) error {
	ctx := context.Background()
	httpClient := gohttp.NewHTTPService(time.Second * 30)

	headers := map[string]string{
		"Content-Type": "application/json",
		"X-Token":      token,
	}

	data := map[string]string{
		"msisdn":      phonenumber.Parse(phoneNumber, "KE"),
		"sms":         message,
		"productID":   "KWIKBET",
		"offercode":   "001003802278",
		"callBackUrl": "https://betsms.kwikbet.io/api/v1/sms/delivery/save",
	}

	body, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("json marshal err: %v", err)
	}

	payload := gohttp.RequestPayload{
		URL:     "http://api.eleza.online/v1/sms/send/",
		Method:  http.MethodPost,
		Headers: headers,
		Body:    bytes.NewBuffer(body),
	}

	if _, err := httpClient.MakeRequest(ctx, payload); err != nil {
		return err
	}

	return nil
}

//ReplySMS replies an sms with LinkID sms to "http://api.eleza.online/v1/sms/reply/"
func ReplySMS(phoneNumber, message, linkID string) error {
	ctx := context.Background()
	httpClient := gohttp.NewHTTPService(time.Second * 30)

	headers := map[string]string{
		"Content-Type": "application/json",
		"X-Token":      token,
	}

	data := map[string]string{
		"msisdn":    phonenumber.Parse(phoneNumber, "KE"),
		"sms":       message,
		"offercode": "001003802278",
		"linkID":    linkID,
	}

	body, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("json marshal err: %v", err)
	}

	payload := gohttp.RequestPayload{
		URL:     "http://api.eleza.online/v1/sms/reply/",
		Method:  http.MethodPost,
		Headers: headers,
		Body:    bytes.NewBuffer(body),
	}

	if _, err := httpClient.MakeRequest(ctx, payload); err != nil {
		return err
	}

	return nil
}

package smsutils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dongri/phonenumber"
	gohttp "github.com/ochom/go-http"
)

//SendSMS send sms to "http://api.eleza.online/v1/sms/send/"
//Must define ELEZA_SMS_TOKEN, ELEZA_OFFER_CODE and ELEZA_PRODUCT_ID in environment file
func SendSMS(phoneNumber, message string) error {
	token, err := getEnv("ELEZA_SMS_TOKEN")
	if err != nil {
		return err
	}

	offerCode, err := getEnv("ELEZA_OFFER_CODE")
	if err != nil {
		return err
	}

	productID, err := getEnv("ELEZA_PRODUCT_ID")
	if err != nil {
		return err
	}

	ctx := context.Background()
	httpClient := gohttp.NewHTTPService(time.Second * 30)

	headers := map[string]string{
		"Content-Type": "application/json",
		"X-Token":      token,
	}

	data := map[string]string{
		"msisdn":      phonenumber.Parse(phoneNumber, "KE"),
		"sms":         message,
		"productID":   productID,
		"offercode":   offerCode,
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
//Must define ELEZA_SMS_TOKEN and ELEZA_OFFER_CODE in environment file
func ReplySMS(phoneNumber, message, linkID string) error {

	token, err := getEnv("ELEZA_SMS_TOKEN")
	if err != nil {
		return err
	}

	offerCode, err := getEnv("ELEZA_OFFER_CODE")
	if err != nil {
		return err
	}

	ctx := context.Background()
	httpClient := gohttp.NewHTTPService(time.Second * 30)

	headers := map[string]string{
		"Content-Type": "application/json",
		"X-Token":      token,
	}

	data := map[string]string{
		"msisdn":    phonenumber.Parse(phoneNumber, "KE"),
		"sms":       message,
		"offercode": offerCode,
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

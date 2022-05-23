package smsutils

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/dongri/phonenumber"
	gohttp "github.com/ochom/go-http"
)

var (
	sendURL  = "http://api.eleza.online/v1/sms/send/"
	replyURL = "http://api.eleza.online/v1/sms/reply/"
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
	httpClient := gohttp.New(time.Second * 30)

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

	if _, _, err := httpClient.Post(ctx, sendURL, headers, body); err != nil {
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
	httpClient := gohttp.New(time.Second * 30)

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

	if _, _, err := httpClient.Post(ctx, replyURL, headers, body); err != nil {
		return err
	}

	return nil
}

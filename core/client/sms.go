package client

import (
	"fmt"
	"ms-otp/model/dto"
	"ms-otp/service"
)

func NewSMSClient() service.SMSClient {
	return &SMSMock{}
}

type SMSMock struct{}

func (c *SMSMock) Send(sms dto.SMS) error {
	fmt.Printf("SMS { %s } sent to: %s\n", sms.Text, sms.Phone)
	return nil
}

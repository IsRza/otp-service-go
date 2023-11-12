package main

import (
	"fmt"
	"ms-otp/core/client"
	"ms-otp/core/storage"
	"ms-otp/model/dto"
	"ms-otp/service"
	"strings"
)

func main() {
	smsClient := client.NewSMSClient()
	sessionStorage := storage.NewSessionStorage()
	otpService := service.NewOTPService(smsClient, sessionStorage)

	phone := "+994-51-511-51-51"
	session, err := otpService.Send(dto.SendOTP{Phone: phone})
	if err != nil {
		panic(err)
	}

	for {
		var otpCode string
		fmt.Print("Enter OTP code: ")
		_, err = fmt.Scan(&otpCode)
		if err != nil {
			panic(err)
		}

		err = otpService.Check(dto.CheckOTP{
			SessionID: session.ID,
			OTPCode:   otpCode,
		})

		if err != nil {
			if strings.Contains(err.Error(), "Invalid OTP code") {
				continue
			} else {
				panic(err)
			}
		} else {
			break
		}
	}

	fmt.Println("main - success")
}

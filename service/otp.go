package service

import (
	"fmt"
	"math/rand"
	"ms-otp/mapper"
	"ms-otp/model/dto"
	"ms-otp/model/entity"
	"os"
)

var random = rand.NewSource(int64(os.Getpid()))

type (
	SMSClient interface {
		Send(sms dto.SMS) error
	}

	SessionStorage interface {
		GetByID(ID string) (*entity.Session, error)
		DeleteByID(ID string) (*entity.Session, error)
		Save(session entity.Session) error
	}

	OTP interface {
		Send(sendDTO dto.SendOTP) (*dto.Session, error)
		Check(checkDTO dto.CheckOTP) error
	}

	OTPService struct {
		smsClient      SMSClient
		sessionStorage SessionStorage
	}
)

func NewOTPService(smsClient SMSClient, sessionStorage SessionStorage) OTP {
	return &OTPService{smsClient, sessionStorage}
}

func (s *OTPService) Send(sendDTO dto.SendOTP) (*dto.Session, error) {
	fmt.Println("OTPService.Send - start")

	session := entity.Session{
		ID:      generateRandomCode(16),
		Code:    generateRandomCode(6),
		Attempt: 0,
	}

	err := s.sessionStorage.Save(session)
	if err != nil {
		return nil, fmt.Errorf("OTPService.Send - sessionStorage.Save: %w", err)
	}

	sms := dto.SMS{
		Phone: sendDTO.Phone,
		Text:  session.Code,
	}

	err = s.smsClient.Send(sms)
	if err != nil {
		return nil, fmt.Errorf("OTPService.Send - smsClient.Send: %w", err)
	}

	responseDTO := mapper.SessionEntityToDTO(session)

	fmt.Println("OTPService.Send - success")
	return &responseDTO, nil
}

func (s *OTPService) Check(checkDTO dto.CheckOTP) error {
	fmt.Println("OTPService.Send - start")

	session, err := s.sessionStorage.GetByID(checkDTO.SessionID)
	if err != nil {
		return fmt.Errorf("OTPService.Check - sessionStorage.GetByID: %w", err)
	}

	if session.Attempt >= 3 {
		_, err = s.sessionStorage.DeleteByID(session.ID)
		if err != nil {
			return fmt.Errorf("OTPService.Check - sessionStorage.DeleteByID: %w", err)
		}
		return fmt.Errorf("OTPService.Check - Attempt overflow")
	}

	if session.Code == checkDTO.OTPCode {
		fmt.Println("OTPService.Send - success")
		return nil
	}

	session.Attempt += 1
	err = s.sessionStorage.Save(*session)
	if err != nil {
		return fmt.Errorf("OTPService.Check - sessionStorage.Save: %w", err)
	}

	return fmt.Errorf("OTPService.Check - Invalid OTP code")
}

func generateRandomCode(len int) string {
	code := ""
	for i := 0; i < len; i++ {
		digit := random.Int63() % 10
		code = fmt.Sprintf("%s%d", code, digit)
	}
	return code
}

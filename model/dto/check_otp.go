package dto

type CheckOTP struct {
	SessionID string `json:"sessionId"`
	OTPCode   string `json:"OTPCode"`
}

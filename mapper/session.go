package mapper

import (
	"ms-otp/model/dto"
	"ms-otp/model/entity"
)

func SessionEntityToDTO(session entity.Session) dto.Session {
	return dto.Session{ID: session.ID}
}

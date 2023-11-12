package storage

import (
	"fmt"
	"ms-otp/model/entity"
	"ms-otp/service"
)

func NewSessionStorage() service.SessionStorage {
	return &SessionMock{
		data: map[string]entity.Session{},
	}
}

type SessionMock struct {
	data map[string]entity.Session
}

func (s *SessionMock) GetByID(ID string) (*entity.Session, error) {
	session, exists := s.data[ID]
	if exists {
		return &session, nil
	} else {
		return nil, fmt.Errorf("session with ID(%s) not found", ID)
	}
}

func (s *SessionMock) DeleteByID(ID string) (*entity.Session, error) {
	session, exists := s.data[ID]
	if exists {
		delete(s.data, ID)
		return &session, nil
	} else {
		return nil, fmt.Errorf("session with ID(%s) not found", ID)
	}
}

func (s *SessionMock) Save(session entity.Session) error {
	s.data[session.ID] = session
	return nil
}

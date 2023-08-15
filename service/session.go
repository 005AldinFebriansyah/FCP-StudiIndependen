package service

import (
	"a21hc3NpZ25tZW50/model"
	repo "a21hc3NpZ25tZW50/repository"
	"errors"
)

type SessionService interface {
	GetSessionByEmail(email string) (model.Session, error)
}

type sessionService struct {
	sessionRepo repo.SessionRepository
}

func NewSessionService(sessionRepo repo.SessionRepository) *sessionService {
	return &sessionService{sessionRepo}
}

func (c *sessionService) GetSessionByEmail(email string) (model.Session, error) {
	session, err := c.sessionRepo.SessionAvailEmail(email)
	if err != nil {
		return model.Session{}, err
	}
	if session.Email == "" {
		return model.Session{}, errors.New("session not found")
	}
	return session, nil // TODO: replace this
}

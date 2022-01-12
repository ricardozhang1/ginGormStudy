package token

import (
	"errors"
	"github.com/google/uuid"

	"time"
)

type Maker interface {
	CreateToken(username string, userID string, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}

var (
	ErrInvalidToken = errors.New("token is valid")
	ErrExpiredToken = errors.New("token has expired")
)

type Payload struct {
	ID uuid.UUID `json:"id"`
	UserId string `json:"user_id"`
	Username string `json:"username"`
	IssuedAt time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(username string, userID string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID: tokenID,
		UserId: userID,
		Username: username,
		IssuedAt: time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}






package token

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

var (
	ErrExpiredToken = errors.New("token has expired")
	ErrInvalidToken = errors.New("token is invalid")
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

type PayloadResponse struct {
	ID                 uuid.UUID `json:"id"`
	Email              string    `json:"email"`
	IssuedAtTimeStamp  int64     `json:"issued_at"`
	ExpiredAtTimeStamp int64     `json:"expired_at"`
}

func (p *PayloadResponse) IssuedAt(IssuedAt time.Time) {
	p.IssuedAtTimeStamp = IssuedAt.Unix()
}
func (p *PayloadResponse) ExpiredAt(ExpiredAt time.Time) {
	p.ExpiredAtTimeStamp = ExpiredAt.Unix()
}
func (p *PayloadResponse) Valid() error {
	expiredAt := time.Unix(p.ExpiredAtTimeStamp, 0)
	if time.Now().After(expiredAt) {
		return ErrExpiredToken
	}
	return nil
}
func (p Payload) Valid() error {
	if time.Now().After(p.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
func NewPayload(email string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	payload := Payload{
		ID:        tokenID,
		Email:     email,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return &payload, nil
}

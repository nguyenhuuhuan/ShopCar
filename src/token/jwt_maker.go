package token

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jinzhu/copier"
	"time"
)

const minSecretKeySize = 32

type JWTMaker struct {
	secretKey string
}

func (j *JWTMaker) CreateToken(email string, duration time.Duration) (string, error) {
	var payloadResponse PayloadResponse
	payload, err := NewPayload(email, duration)
	if err != nil {
		return "", err
	}
	_ = copier.Copy(&payloadResponse, payload)
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &payloadResponse)
	return jwtToken.SignedString([]byte(j.secretKey))
}

func (j *JWTMaker) VerifyToken(token string) (*PayloadResponse, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(j.secretKey), nil
	}
	jwtToken, err := jwt.ParseWithClaims(token, &PayloadResponse{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*PayloadResponse)
	if !ok {
		return nil, ErrInvalidToken
	}
	return payload, nil
}

func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be least %d characters", minSecretKeySize)
	}
	return &JWTMaker{secretKey}, nil
}

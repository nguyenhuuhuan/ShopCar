package token

import (
	"Improve/src/utils"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestJWTMaker(t *testing.T) {
	maker, err := NewJWTMaker(utils.RandomString(32))
	require.NoError(t, err)

	email := utils.RandomString(6)
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := time.Now().Add(duration)

	token, err := maker.CreateToken(email, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payloadResponse, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payloadResponse)

	require.NotZero(t, payloadResponse.ID)
	require.Equal(t, email, payloadResponse.Email)
	require.WithinDuration(t, issuedAt, time.Unix(payloadResponse.IssuedAtTimeStamp, 0), time.Second)
	require.WithinDuration(t, expiredAt, time.Unix(payloadResponse.ExpiredAtTimeStamp, 0), time.Second)

}

func TestExpiredToken(t *testing.T) {
	maker, err := NewJWTMaker(utils.RandomString(32))
	require.NoError(t, err)

	token, err := maker.CreateToken(utils.RandomString(6), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)

}

func TestInvalidJWTTokenAlgNone(t *testing.T) {

	payload, err := NewPayload(utils.RandomString(6), time.Minute)
	require.NoError(t, err)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)

	maker, err := NewJWTMaker(utils.RandomString(32))
	require.NoError(t, err)

	payloadResponse, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrInvalidToken.Error())
	require.Nil(t, payloadResponse)
}

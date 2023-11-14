package securityJWT

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var ErrInvalidToken = errors.New("token is invalid")

const minSecretKeySize = 32

type JWTMaker struct {
	SecretKey string
	OptionJWT OptionJwt
}

// Option Jwt
type OptionJwt struct {
	TimeDuration time.Duration
}

type Maker interface {
	BuildToken(string) (string, error)
	VerifyToken(string) (*Payload, error)
}

func NewJWTMaker(secretKey string, OptionJWT OptionJwt) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("Invalid key size: must be at least %d characters", minSecretKeySize)
	}
	return &JWTMaker{
		SecretKey: secretKey,
		OptionJWT: OptionJWT}, nil
}

func (maker *JWTMaker) BuildToken(username string) (string, error) {
	payload, err := NewPayLoad(username, maker.OptionJWT.TimeDuration)
	if err != nil {
		return "", err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	return jwtToken.SignedString([]byte(maker.SecretKey))
}

func (marker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	keyfunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(marker.SecretKey), nil
	}
	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyfunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrInvalidToken) {
			return nil, ErrInvalidToken
		}
		return nil, ErrExprireDate
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}

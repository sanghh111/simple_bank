package securityJWT

import (
	"errors"
	"fmt"
	"time"
)

var ErrExprireDate = errors.New("token has expired")

type Payload struct {
	Username    string    `json:"username"`
	ExpiredDate time.Time `json:"expiredDate"`
}

// NewPayLoad creates a new token payload with a specific username and duration
func NewPayLoad(Username string, d time.Duration) (*Payload, error) {
	now := time.Now()
	return &Payload{
		Username:    Username,
		ExpiredDate: now.Add(d),
	}, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredDate) {
		fmt.Println(1)
		return ErrExprireDate
	}
	return nil
}

package jwt

import (
	"echo_shop/app/models"
	"echo_shop/pkg/constants"
	"echo_shop/pkg/errno"
	"time"
)

// Info token info
type Info struct {
	Token     string `json:"token"`
	Type      string `json:"type"`
	ExpiresIn string `json:"expires_in"`
}

// Sign 签发 token
func Sign(u *models.User) (*Info, *errno.Errno) {
	t, claims, err := create(u)
	if err != nil || t == "" {
		return nil, errno.JWTTokenErr.SetMessage(err.Error())
	}

	return &Info{
		Token:     t,
		Type:      tokenInHeaderIdentification,
		ExpiresIn: time.Unix(claims.ExpiresAt, 0).Format(constants.DateTimeLayout),
	}, nil
}

// Refresh 刷新 token
func Refresh(tokenString string) (*Info, *errno.Errno) {
	t, claims, err := refresh(tokenString)
	if err != nil || t == "" {
		return nil, err
	}

	return &Info{
		Token:     t,
		Type:      tokenInHeaderIdentification,
		ExpiresIn: time.Unix(claims.ExpiresAt, 0).Format(constants.DateTimeLayout),
	}, nil
}

// Forget 使 token 失效
func Forget(tokenString string, remainTime time.Duration) {
	forget(tokenString, remainTime)
}

package entity

import "errors"

var (
	ErrLoginRequired       = errors.New("login required")
	ErrCodeAlreadyRedeemed = errors.New("code already redeemed")
	ErrCodeExpired         = errors.New("code expired")
	ErrCoolDownRequired    = errors.New("try again later")
)

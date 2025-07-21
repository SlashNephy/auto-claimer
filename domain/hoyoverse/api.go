package hoyoverse

import (
	"fmt"

	"github.com/SlashNephy/auto-claimer/domain/entity"
)

type APIResponse[T any] struct {
	Data    *T     `json:"data"`
	Code    int    `json:"retcode"`
	Message string `json:"message"`
}

func MapAPIError(code int, message string) error {
	switch code {
	case 0:
		return nil
	case -100, -1071:
		return entity.ErrLoginRequired
	case -2001:
		return entity.ErrCodeExpired
	case -2016:
		return entity.ErrCoolDownRequired
	case -2017, -2018:
		return entity.ErrCodeAlreadyRedeemed
	default:
		return fmt.Errorf("HoYoverse API error: %s (%d)", message, code)
	}
}

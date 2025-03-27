package workflow

import (
	"context"

	"github.com/SlashNephy/auto-claimer/domain/entity"
	"github.com/SlashNephy/auto-claimer/domain/hoyoverse"
)

type LoginHoYoverseAccountStore interface {
	Login(ctx context.Context, email, password string) error
	ListGameAccounts(ctx context.Context) ([]*hoyoverse.GameAccount, error)
}

type RedeemHonkaiStarrailCodeStore interface {
	RedeemCode(ctx context.Context, account *hoyoverse.GameAccount, code *hoyoverse.Code) error
}

type MarkCodeAsRedeemedStore interface {
	MarkCodeAsRedeemed(ctx context.Context, account entity.Account, code entity.Code) error
}

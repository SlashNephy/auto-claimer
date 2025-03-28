package query

import (
	"context"

	"github.com/SlashNephy/auto-claimer/domain/entity"
	"github.com/SlashNephy/auto-claimer/domain/hoyoverse"
)

type RedeemedCodeQuery interface {
	ListRedeemedCodes(ctx context.Context, account entity.Account) ([]string, error)
}

type HonkaiStarrailCodeQuery interface {
	ListAvailableCodes(ctx context.Context) ([]*hoyoverse.Code, error)
}

type GenshinImpactCodeQuery interface {
	ListAvailableCodes(ctx context.Context) ([]*hoyoverse.Code, error)
}

type ZenlessZoneZeroCodeQuery interface {
	ListAvailableCodes(ctx context.Context) ([]*hoyoverse.Code, error)
}

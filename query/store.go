package query

import (
	"context"

	"github.com/SlashNephy/auto-claimer/domain/entity"
	"github.com/SlashNephy/auto-claimer/domain/hoyoverse"
)

type RedeemedCodeStore interface {
	ListRedeemedCodes(ctx context.Context, account entity.Account) ([]string, error)
}

type HonkaiStarrailCodeStore interface {
	ListAvailableHonkaiStarrailCodes(ctx context.Context) ([]*hoyoverse.Code, error)
}

type GenshinImpactCodeStore interface {
	ListAvailableGenshinImpactCodes(ctx context.Context) ([]*hoyoverse.Code, error)
}

type ZenlessZoneZeroCodeStore interface {
	ListAvailableZenlessZoneZeroCodes(ctx context.Context) ([]*hoyoverse.Code, error)
}

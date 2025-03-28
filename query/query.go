package query

import (
	"context"

	"github.com/SlashNephy/auto-claimer/domain/entity"
	"github.com/SlashNephy/auto-claimer/domain/hoyoverse"
)

type RedeemedCodeQuery interface {
	ListRedeemedCodes(ctx context.Context, account entity.Account) ([]string, error)
}

type HonkaiStarrailQuery interface {
	ListHonkaiStarrailGameAccounts(ctx context.Context) ([]*hoyoverse.GameAccount, error)
	ListAvailableHonkaiStarrailCodes(ctx context.Context) ([]*hoyoverse.Code, error)
}

type GenshinImpactQuery interface {
	ListGenshinImpactGameAccounts(ctx context.Context) ([]*hoyoverse.GameAccount, error)
	ListAvailableGenshinImpactCodes(ctx context.Context) ([]*hoyoverse.Code, error)
}

type ZenlessZoneZeroQuery interface {
	ListZenlessZoneZeroGameAccounts(ctx context.Context) ([]*hoyoverse.GameAccount, error)
	ListAvailableZenlessZoneZeroCodes(ctx context.Context) ([]*hoyoverse.Code, error)
}

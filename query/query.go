package query

import (
	"context"

	"github.com/SlashNephy/auto-claimer/domain/hoyoverse"
)

type HonkaiStarrailCodeQuery interface {
	ListAvailableCodes(ctx context.Context) ([]*hoyoverse.Code, error)
	ListRedeemedCodes(ctx context.Context, account *hoyoverse.GameAccount) ([]string, error)
}

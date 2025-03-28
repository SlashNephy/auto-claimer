package query

import (
	"context"

	"github.com/SlashNephy/auto-claimer/domain/entity"
	"github.com/SlashNephy/auto-claimer/domain/hoyoverse"
)

type HonkaiStarrailCodeStore interface {
	ListAvailableHonkaiStarrailCodes(ctx context.Context) ([]*hoyoverse.Code, error)
}

type RedeemedCodeStore interface {
	ListRedeemedCodes(ctx context.Context, account entity.Account) ([]string, error)
}

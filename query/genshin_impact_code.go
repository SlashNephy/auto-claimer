package query

import (
	"context"

	"github.com/SlashNephy/auto-claimer/domain/hoyoverse"
	"github.com/samber/do/v2"
)

type GenshinImpactQueryImpl struct {
	store     GenshinImpactStore
	codeStore GenshinImpactCodeStore
}

func NewGenshinImpactQuery(i do.Injector) (*GenshinImpactQueryImpl, error) {
	return &GenshinImpactQueryImpl{
		store:     do.MustInvoke[GenshinImpactStore](i),
		codeStore: do.MustInvoke[GenshinImpactCodeStore](i),
	}, nil
}

func (q *GenshinImpactQueryImpl) ListGenshinImpactGameAccounts(ctx context.Context) ([]*hoyoverse.GameAccount, error) {
	return q.store.ListGenshinImpactGameAccounts(ctx)
}

func (q *GenshinImpactQueryImpl) ListAvailableGenshinImpactCodes(ctx context.Context) ([]*hoyoverse.Code, error) {
	return q.codeStore.ListAvailableGenshinImpactCodes(ctx)
}

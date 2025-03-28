package query

import (
	"context"

	"github.com/SlashNephy/auto-claimer/domain/hoyoverse"
	"github.com/samber/do/v2"
)

type GenshinImpactCodeQueryImpl struct {
	store GenshinImpactCodeStore
}

func NewGenshinImpactCodeQuery(i do.Injector) (*GenshinImpactCodeQueryImpl, error) {
	return &GenshinImpactCodeQueryImpl{
		store: do.MustInvoke[GenshinImpactCodeStore](i),
	}, nil
}

func (q *GenshinImpactCodeQueryImpl) ListAvailableCodes(ctx context.Context) ([]*hoyoverse.Code, error) {
	return q.store.ListAvailableGenshinImpactCodes(ctx)
}

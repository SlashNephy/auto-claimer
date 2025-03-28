package query

import (
	"context"

	"github.com/SlashNephy/auto-claimer/domain/hoyoverse"
	"github.com/samber/do/v2"
)

type HonkaiStarrailCodeQueryImpl struct {
	store HonkaiStarrailCodeStore
}

func NewHonkaiStarrailCodeQuery(i do.Injector) (*HonkaiStarrailCodeQueryImpl, error) {
	return &HonkaiStarrailCodeQueryImpl{
		store: do.MustInvoke[HonkaiStarrailCodeStore](i),
	}, nil
}

func (q *HonkaiStarrailCodeQueryImpl) ListAvailableCodes(ctx context.Context) ([]*hoyoverse.Code, error) {
	return q.store.ListAvailableHonkaiStarrailCodes(ctx)
}

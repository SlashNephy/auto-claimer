package query

import (
	"context"

	"github.com/SlashNephy/auto-claimer/domain/hoyoverse"
	"github.com/samber/do/v2"
)

type HonkaiStarrailQueryImpl struct {
	store     HonkaiStarrailStore
	codeStore HonkaiStarrailCodeStore
}

func NewHonkaiStarrailQuery(i do.Injector) (*HonkaiStarrailQueryImpl, error) {
	return &HonkaiStarrailQueryImpl{
		store:     do.MustInvoke[HonkaiStarrailStore](i),
		codeStore: do.MustInvoke[HonkaiStarrailCodeStore](i),
	}, nil
}

func (q *HonkaiStarrailQueryImpl) ListHonkaiStarrailGameAccounts(ctx context.Context) ([]*hoyoverse.GameAccount, error) {
	return q.store.ListHonkaiStarrailGameAccounts(ctx)
}

func (q *HonkaiStarrailQueryImpl) ListAvailableHonkaiStarrailCodes(ctx context.Context) ([]*hoyoverse.Code, error) {
	return q.codeStore.ListAvailableHonkaiStarrailCodes(ctx)
}

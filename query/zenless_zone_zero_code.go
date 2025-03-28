package query

import (
	"context"

	"github.com/SlashNephy/auto-claimer/domain/hoyoverse"
	"github.com/samber/do/v2"
)

type ZenlessZoneZeroQueryImpl struct {
	store     ZenlessZoneZeroStore
	codeStore ZenlessZoneZeroCodeStore
}

func NewZenlessZoneZeroQuery(i do.Injector) (*ZenlessZoneZeroQueryImpl, error) {
	return &ZenlessZoneZeroQueryImpl{
		store:     do.MustInvoke[ZenlessZoneZeroStore](i),
		codeStore: do.MustInvoke[ZenlessZoneZeroCodeStore](i),
	}, nil
}

func (q *ZenlessZoneZeroQueryImpl) ListZenlessZoneZeroGameAccounts(ctx context.Context) ([]*hoyoverse.GameAccount, error) {
	return q.store.ListZenlessZoneZeroGameAccounts(ctx)
}

func (q *ZenlessZoneZeroQueryImpl) ListAvailableZenlessZoneZeroCodes(ctx context.Context) ([]*hoyoverse.Code, error) {
	return q.codeStore.ListAvailableZenlessZoneZeroCodes(ctx)
}

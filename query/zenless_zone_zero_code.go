package query

import (
	"context"

	"github.com/SlashNephy/auto-claimer/domain/hoyoverse"
	"github.com/samber/do/v2"
)

type ZenlessZoneZeroCodeQueryImpl struct {
	store ZenlessZoneZeroCodeStore
}

func NewZenlessZoneZeroCodeQuery(i do.Injector) (*ZenlessZoneZeroCodeQueryImpl, error) {
	return &ZenlessZoneZeroCodeQueryImpl{
		store: do.MustInvoke[ZenlessZoneZeroCodeStore](i),
	}, nil
}

func (q *ZenlessZoneZeroCodeQueryImpl) ListAvailableCodes(ctx context.Context) ([]*hoyoverse.Code, error) {
	return q.store.ListAvailableZenlessZoneZeroCodes(ctx)
}

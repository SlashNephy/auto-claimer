package query

import (
	"context"

	"github.com/SlashNephy/auto-claimer/domain/entity"
	"github.com/samber/do/v2"
)

type RedeemedCodeQueryImpl struct {
	store RedeemedCodeStore
}

func NewRedeemedCodeQuery(i do.Injector) (*RedeemedCodeQueryImpl, error) {
	return &RedeemedCodeQueryImpl{
		store: do.MustInvoke[RedeemedCodeStore](i),
	}, nil
}

func (q *RedeemedCodeQueryImpl) ListRedeemedCodes(ctx context.Context, account entity.Account) ([]string, error) {
	return q.store.ListRedeemedCodes(ctx, account)
}

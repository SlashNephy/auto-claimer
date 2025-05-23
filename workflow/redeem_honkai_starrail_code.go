package workflow

import (
	"context"
	"errors"

	"github.com/SlashNephy/auto-claimer/domain/entity"
	"github.com/samber/do/v2"
)

func NewRedeemHonkaiStarrailCodeWorkflow(i do.Injector) (RedeemHonkaiStarrailCodeWorkflow, error) {
	redeemer := do.MustInvoke[RedeemHonkaiStarrailCodeStore](i)
	marker := do.MustInvoke[MarkCodeAsRedeemedStore](i)

	return NewWorkflowFunc(func(ctx context.Context, command *RedeemHonkaiStarrailCodeCommand) (*HonkaiStarrailCodeRedeemedEvent, error) {
		alreadyRedeemed := false
		if err := redeemer.RedeemHonkaiStarrailCode(ctx, command.Account, command.Code); err != nil {
			if errors.Is(err, entity.ErrCodeAlreadyRedeemed) || errors.Is(err, entity.ErrCodeExpired) {
				// すでに交換済みのコードであったり、期限切れだったコードの場合にも MarkCodeAsRedeemed されるようにする
				alreadyRedeemed = true
			} else {
				return nil, err
			}
		}

		if err := marker.MarkCodeAsRedeemed(ctx, command.Account, command.Code); err != nil {
			return nil, err
		}

		if alreadyRedeemed {
			return nil, entity.ErrCodeAlreadyRedeemed
		}

		return &HonkaiStarrailCodeRedeemedEvent{
			RedeemedCode: command.Code,
		}, nil
	}), nil
}

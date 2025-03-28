package workflow

import (
	"context"
	"errors"

	"github.com/SlashNephy/auto-claimer/domain/entity"
	"github.com/samber/do/v2"
)

func NewRedeemZenlessZoneZeroCodeWorkflow(i do.Injector) (RedeemZenlessZoneZeroCodeWorkflow, error) {
	redeemer := do.MustInvoke[RedeemZenlessZoneZeroCodeStore](i)
	marker := do.MustInvoke[MarkCodeAsRedeemedStore](i)

	return NewWorkflowFunc(func(ctx context.Context, command *RedeemZenlessZoneZeroCodeCommand) (*ZenlessZoneZeroCodeRedeemedEvent, error) {
		alreadyRedeemed := false
		if err := redeemer.RedeemZenlessZoneZeroCode(ctx, command.Account, command.Code); err != nil {
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

		return &ZenlessZoneZeroCodeRedeemedEvent{
			RedeemedCode: command.Code,
		}, nil
	}), nil
}

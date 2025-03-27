package pipeline

import (
	"context"
	"errors"
	"log/slog"
	"slices"
	"time"

	"github.com/SlashNephy/auto-claimer/domain/entity"
	"github.com/SlashNephy/auto-claimer/domain/hoyoverse"
	"github.com/SlashNephy/auto-claimer/query"
	"github.com/SlashNephy/auto-claimer/workflow"
	"github.com/samber/do/v2"
)

func NewRedeemHonkaiStarrailCodesPipeline(i do.Injector) (RedeemHonkaiStarrailCodesPipeline, error) {
	query := do.MustInvoke[query.HonkaiStarrailCodeQuery](i)
	login := do.MustInvoke[workflow.LoginHoYoverseAccountWorkflow](i)
	redeemer := do.MustInvoke[workflow.RedeemHonkaiStarrailCodeWorkflow](i)
	notifier := do.MustInvoke[workflow.NotifyHoYoverseCodeRedeemedWorkflow](i)

	return NewPipelineFunc(func(ctx context.Context, input *RedeemHonkaiStarrailCodesInput) (*RedeemHonkaiStarrailCodesOutput, error) {
		availableCodes, err := query.ListAvailableCodes(ctx)
		if err != nil {
			return nil, err
		}

		loggedIn, err := login.Do(ctx, &workflow.LoginHoYoverseAccountCommand{
			Email:    input.HoYoverseEmail,
			Password: input.HoYoverseEmail,
		})
		if err != nil {
			return nil, err
		}

		redeem := func(account *hoyoverse.GameAccount, code *hoyoverse.Code) error {
			redeemed, err := redeemer.Do(ctx, &workflow.RedeemHonkaiStarrailCodeCommand{
				Account: account,
				Code:    code,
			})
			if err != nil {
				if errors.Is(err, entity.ErrCodeAlreadyRedeemed) {
					slog.WarnContext(
						ctx,
						"code already redeemed",
						slog.String("code", code.Code),
						slog.String("account_id", account.UID),
						slog.String("account_name", account.Nickname),
						slog.String("account_region", account.Region),
					)
					return nil
				}
				if errors.Is(err, entity.ErrCodeExpired) {
					slog.WarnContext(
						ctx,
						"code expired",
						slog.String("code", code.Code),
						slog.String("account_id", account.UID),
						slog.String("account_name", account.Nickname),
						slog.String("account_region", account.Region),
					)
					return nil
				}
				return err
			}

			slog.InfoContext(
				ctx,
				"code redeemed",
				slog.String("code", redeemed.RedeemedCode.Code),
				slog.Any("rewards", redeemed.RedeemedCode.Rewards),
				slog.String("account_id", account.UID),
				slog.String("account_name", account.Nickname),
				slog.String("account_region", account.Region),
			)

			if input.DiscordWebhookURL != nil {
				_, err = notifier.Do(ctx, &workflow.NotifyHoYoverseCodeRedeemedCommand{
					DiscordWebhookURL: *input.DiscordWebhookURL,
					RedeemedCode:      redeemed.RedeemedCode,
					Account:           account,
				})
				if err != nil {
					return err
				}
			}

			return nil
		}

		for _, account := range loggedIn.Accounts {
			redeemedCodes, err := query.ListRedeemedCodes(ctx, account)
			if err != nil {
				return nil, err
			}

			for _, code := range availableCodes {
				if slices.Contains(redeemedCodes, code.Code) {
					slog.InfoContext(
						ctx,
						"skipped redeemed code",
						slog.String("code", code.Code),
						slog.String("account_id", account.UID),
						slog.String("account_name", account.Nickname),
						slog.String("account_region", account.Region),
					)
					continue
				}

				if err = redeem(account, code); err != nil {
					return nil, err
				}

				time.Sleep(10 * time.Second)
			}
		}

		return &RedeemHonkaiStarrailCodesOutput{}, nil
	}), nil
}

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
	"github.com/samber/lo"
)

func NewRedeemZenlessZoneZeroCodesPipeline(i do.Injector) (RedeemZenlessZoneZeroCodesPipeline, error) {
	redeemedCodeQuery := do.MustInvoke[query.RedeemedCodeQuery](i)
	zenlessZoneZeroQuery := do.MustInvoke[query.ZenlessZoneZeroQuery](i)
	redeemWorkflow := do.MustInvoke[workflow.RedeemZenlessZoneZeroCodeWorkflow](i)
	notifyWorkflow := do.MustInvoke[workflow.NotifyHoYoverseCodeRedeemedWorkflow](i)

	return NewPipelineFunc(func(ctx context.Context, input *RedeemZenlessZoneZeroCodesInput) (*RedeemZenlessZoneZeroCodesOutput, error) {
		availableCodes, err := zenlessZoneZeroQuery.ListAvailableZenlessZoneZeroCodes(ctx)
		if err != nil {
			return nil, err
		}

		accounts, err := zenlessZoneZeroQuery.ListZenlessZoneZeroGameAccounts(ctx)
		if err != nil {
			return nil, err
		}

		slog.InfoContext(ctx, "Zenless Zone Zero game accounts",
			slog.Any("accounts", lo.Map(accounts, func(account *hoyoverse.GameAccount, _ int) string {
				return account.String()
			})),
		)

		redeem := func(account *hoyoverse.GameAccount, code *hoyoverse.Code) error {
			redeemed, err := redeemWorkflow.Do(ctx, &workflow.RedeemZenlessZoneZeroCodeCommand{
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
				_, err = notifyWorkflow.Do(ctx, &workflow.NotifyHoYoverseCodeRedeemedCommand{
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

		for _, account := range accounts {
			redeemedCodes, err := redeemedCodeQuery.ListRedeemedCodes(ctx, account)
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

		return &RedeemZenlessZoneZeroCodesOutput{}, nil
	}), nil
}

package pipeline

import (
	"context"
	"log/slog"
	"slices"

	"github.com/SlashNephy/auto-claimer/domain/entity"
	"github.com/SlashNephy/auto-claimer/workflow"
	"github.com/samber/do/v2"
)

type BatchRedeemCodesConfig struct {
	Games             []entity.Game `env:"GAMES"`
	HoYoverseEmail    *string       `env:"HOYOVERSE_EMAIL"`
	HoYoversePassword *string       `env:"HOYOVERSE_PASSWORD"`
	DiscordWebhookURL *string       `env:"DISCORD_WEBHOOK_URL"`
}

func NewBatchRedeemCodesPipeline(i do.Injector) (BatchRedeemCodesPipeline, error) {
	config := do.MustInvoke[*BatchRedeemCodesConfig](i)
	loginHoYoverseAccountWorkflow := do.MustInvoke[workflow.LoginHoYoverseAccountWorkflow](i)
	redeemHonkaiStarrailCodesPipeline := do.MustInvoke[RedeemHonkaiStarrailCodesPipeline](i)
	redeemGenshinImpactCodesPipeline := do.MustInvoke[RedeemGenshinImpactCodesPipeline](i)
	redeemZenlessZoneZeroCodesPipeline := do.MustInvoke[RedeemZenlessZoneZeroCodesPipeline](i)

	return NewPipelineFunc(func(ctx context.Context, input *BatchRedeemCodesInput) (*BatchRedeemCodesOutput, error) {
		if len(config.Games) == 0 {
			slog.InfoContext(ctx, "no games specified")
			return &BatchRedeemCodesOutput{}, nil
		}

		slog.InfoContext(ctx, "redeeming games", slog.Any("games", config.Games))

		if config.HoYoverseEmail != nil && config.HoYoversePassword != nil {
			_, err := loginHoYoverseAccountWorkflow.Do(ctx, &workflow.LoginHoYoverseAccountCommand{
				Email:    *config.HoYoverseEmail,
				Password: *config.HoYoversePassword,
			})
			if err != nil {
				return nil, err
			}

			if slices.Contains(config.Games, entity.GameHonkaiStarrail) {
				_, err := redeemHonkaiStarrailCodesPipeline.Do(ctx, &RedeemHonkaiStarrailCodesInput{
					DiscordWebhookURL: config.DiscordWebhookURL,
				})
				if err != nil {
					slog.ErrorContext(ctx, "failed to redeem Honkai Starrail codes", slog.String("error", err.Error()))
				}
			}

			if slices.Contains(config.Games, entity.GameGenshinImpact) {
				_, err := redeemGenshinImpactCodesPipeline.Do(ctx, &RedeemGenshinImpactCodesInput{
					DiscordWebhookURL: config.DiscordWebhookURL,
				})
				if err != nil {
					slog.ErrorContext(ctx, "failed to redeem Genshin Impact codes", slog.String("error", err.Error()))
				}
			}

			if slices.Contains(config.Games, entity.GameZenlessZoneZero) {
				_, err := redeemZenlessZoneZeroCodesPipeline.Do(ctx, &RedeemZenlessZoneZeroCodesInput{
					DiscordWebhookURL: config.DiscordWebhookURL,
				})
				if err != nil {
					slog.ErrorContext(ctx, "failed to redeem Zenless Zone Zero codes", slog.String("error", err.Error()))
				}
			}
		}

		return &BatchRedeemCodesOutput{}, nil
	}), nil
}

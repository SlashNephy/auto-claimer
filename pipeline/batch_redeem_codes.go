package pipeline

import (
	"context"

	"github.com/samber/do/v2"
)

type BatchRedeemCodesConfig struct {
	HoYoverseEmail    *string `env:"HOYOVERSE_EMAIL"`
	HoYoversePassword *string `env:"HOYOVERSE_PASSWORD"`
	DiscordWebhookURL *string `env:"DISCORD_WEBHOOK_URL"`
}

func NewBatchRedeemCodesPipeline(i do.Injector) (BatchRedeemCodesPipeline, error) {
	config := do.MustInvoke[*BatchRedeemCodesConfig](i)
	redeemHonkaiStarrailCodes := do.MustInvoke[RedeemHonkaiStarrailCodesPipeline](i)

	return NewPipelineFunc(func(ctx context.Context, input *BatchRedeemCodesInput) (*BatchRedeemCodesOutput, error) {
		if config.HoYoverseEmail != nil && config.HoYoversePassword != nil {
			_, err := redeemHonkaiStarrailCodes.Do(ctx, &RedeemHonkaiStarrailCodesInput{
				HoYoverseEmail:    *config.HoYoverseEmail,
				HoYoversePassword: *config.HoYoversePassword,
				DiscordWebhookURL: config.DiscordWebhookURL,
			})
			if err != nil {
				return nil, err
			}
		}

		return &BatchRedeemCodesOutput{}, nil
	}), nil
}

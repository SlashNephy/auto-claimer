package pipeline

import (
	"context"
)

type Pipeline[Input, Output any] interface {
	Do(ctx context.Context, input Input) (Output, error)
}

type PipelineFunc[Input, Output any] func(ctx context.Context, input Input) (Output, error)

func NewPipelineFunc[Input, Output any](f PipelineFunc[Input, Output]) PipelineFunc[Input, Output] {
	return f
}

func (f PipelineFunc[Input, Output]) Do(ctx context.Context, input Input) (Output, error) {
	return f(ctx, input)
}

type (
	RedeemHonkaiStarrailCodesPipeline Pipeline[*RedeemHonkaiStarrailCodesInput, *RedeemHonkaiStarrailCodesOutput]
	RedeemHonkaiStarrailCodesInput    struct {
		HoYoverseEmail    string
		HoYoversePassword string
		DiscordWebhookURL *string
	}
	RedeemHonkaiStarrailCodesOutput struct{}
)

type (
	BatchRedeemCodesPipeline Pipeline[*BatchRedeemCodesInput, *BatchRedeemCodesOutput]
	BatchRedeemCodesInput    struct{}
	BatchRedeemCodesOutput   struct{}
)

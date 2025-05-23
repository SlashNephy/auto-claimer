package workflow

import (
	"context"

	"github.com/SlashNephy/auto-claimer/domain/hoyoverse"
	"github.com/disgoorg/disgo/discord"
)

type Workflow[Command, Event any] interface {
	Do(ctx context.Context, command Command) (Event, error)
}

type WorkflowFunc[Command, Event any] func(ctx context.Context, command Command) (Event, error)

func NewWorkflowFunc[Command, Event any](f WorkflowFunc[Command, Event]) WorkflowFunc[Command, Event] {
	return f
}

func (f WorkflowFunc[Command, Event]) Do(ctx context.Context, command Command) (Event, error) {
	return f(ctx, command)
}

type (
	LoginHoYoverseAccountWorkflow Workflow[*LoginHoYoverseAccountCommand, *HoYoverseAccountLoggedInEvent]
	LoginHoYoverseAccountCommand  struct {
		Email         string
		Password      string
		HoYoverseUUID *string
		MiHoYoUUID    *string
	}
	HoYoverseAccountLoggedInEvent struct{}
)

type (
	RedeemHonkaiStarrailCodeWorkflow Workflow[*RedeemHonkaiStarrailCodeCommand, *HonkaiStarrailCodeRedeemedEvent]
	RedeemHonkaiStarrailCodeCommand  struct {
		Account *hoyoverse.GameAccount
		Code    *hoyoverse.Code
	}
	HonkaiStarrailCodeRedeemedEvent struct {
		RedeemedCode *hoyoverse.Code
	}
)

type (
	RedeemGenshinImpactCodeWorkflow Workflow[*RedeemGenshinImpactCodeCommand, *GenshinImpactCodeRedeemedEvent]
	RedeemGenshinImpactCodeCommand  struct {
		Account *hoyoverse.GameAccount
		Code    *hoyoverse.Code
	}
	GenshinImpactCodeRedeemedEvent struct {
		RedeemedCode *hoyoverse.Code
	}
)

type (
	RedeemZenlessZoneZeroCodeWorkflow Workflow[*RedeemZenlessZoneZeroCodeCommand, *ZenlessZoneZeroCodeRedeemedEvent]
	RedeemZenlessZoneZeroCodeCommand  struct {
		Account *hoyoverse.GameAccount
		Code    *hoyoverse.Code
	}
	ZenlessZoneZeroCodeRedeemedEvent struct {
		RedeemedCode *hoyoverse.Code
	}
)

type (
	NotifyHoYoverseCodeRedeemedWorkflow Workflow[*NotifyHoYoverseCodeRedeemedCommand, *HoYoverseCodeRedeemedNotifiedEvent]
	NotifyHoYoverseCodeRedeemedCommand  struct {
		DiscordWebhookURL string
		RedeemedCode      *hoyoverse.Code
		Account           *hoyoverse.GameAccount
	}
	HoYoverseCodeRedeemedNotifiedEvent struct {
		DiscordMessage *discord.Message
	}
)

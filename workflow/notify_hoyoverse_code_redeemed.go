package workflow

import (
	"context"
	"strings"

	"github.com/SlashNephy/auto-claimer/locale"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/webhook"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/samber/do/v2"
)

func NewNotifyHoYoverseCodeRedeemedWorkflow(do.Injector) (NotifyHoYoverseCodeRedeemedWorkflow, error) {
	return NewWorkflowFunc(func(ctx context.Context, command *NotifyHoYoverseCodeRedeemedCommand) (*HoYoverseCodeRedeemedNotifiedEvent, error) {
		client, err := webhook.NewWithURL(command.DiscordWebhookURL)
		if err != nil {
			return nil, err
		}

		localizer := locale.NewLocalizer(command.Account.Language)

		defer client.Close(ctx)
		message, err := client.CreateMessage(
			discord.NewWebhookMessageCreateBuilder().
				SetUsername("auto-claimer").
				AddEmbeds(
					discord.NewEmbedBuilder().
						SetTitle(
							localizer.MustLocalize(&i18n.LocalizeConfig{
								DefaultMessage: &i18n.Message{
									ID:    "HoYoverseCodeRedeemedNotification.Title",
									Other: "`{{.Code}}` redeemed!",
								},
								TemplateData: map[string]any{
									"Code": command.RedeemedCode.Code,
								},
							}),
						).
						SetAuthor(
							localizer.MustLocalize(&i18n.LocalizeConfig{
								DefaultMessage: &i18n.Message{
									ID:    "HoYoverseCodeRedeemedNotification.Author",
									Other: "{{.Nickname}} (UID:{{.UID}})",
								},
								TemplateData: map[string]any{
									"Nickname": command.Account.Nickname,
									"UID":      command.Account.UID,
								},
							}),
							"", "",
						).
						SetFields(
							discord.EmbedField{
								Name: localizer.MustLocalizeMessage(&i18n.Message{
									ID:    "HoYoverseCodeRedeemedNotification.RewardsField",
									Other: "Rewards",
								}),
								Value: strings.Join(command.RedeemedCode.Rewards, "\n"),
							},
						).
						SetFooter(
							localizer.MustLocalizeMessage(command.RedeemedCode.Game.LocalizeMessage()),
							"",
						).
						Build(),
				).
				Build(),
		)
		if err != nil {
			return nil, err
		}

		return &HoYoverseCodeRedeemedNotifiedEvent{
			DiscordMessage: message,
		}, nil
	}), nil
}

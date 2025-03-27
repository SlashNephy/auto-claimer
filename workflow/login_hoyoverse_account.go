package workflow

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/SlashNephy/auto-claimer/domain/hoyoverse"
	"github.com/samber/do/v2"
	"github.com/samber/lo"
)

func NewLoginHoYoverseAccountWorkflow(i do.Injector) (LoginHoYoverseAccountWorkflow, error) {
	store := do.MustInvoke[LoginHoYoverseAccountStore](i)

	return NewWorkflowFunc(func(ctx context.Context, command *LoginHoYoverseAccountCommand) (*HoYoverseAccountLoggedInEvent, error) {
		if err := store.Login(ctx, command.Email, command.Password); err != nil {
			return nil, fmt.Errorf("failed to login HoYoverse account: %w", err)
		}

		accounts, err := store.ListGameAccounts(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to list HoYoverse game accounts: %w", err)
		}

		slog.InfoContext(ctx, "HoYoverse game accounts",
			slog.Any("accounts", lo.Map(accounts, func(account *hoyoverse.GameAccount, _ int) string {
				return account.String()
			})),
		)

		return &HoYoverseAccountLoggedInEvent{
			Accounts: accounts,
		}, nil
	}), nil
}

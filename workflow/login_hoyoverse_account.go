package workflow

import (
	"context"
	"fmt"

	"github.com/samber/do/v2"
)

func NewLoginHoYoverseAccountWorkflow(i do.Injector) (LoginHoYoverseAccountWorkflow, error) {
	store := do.MustInvoke[LoginHoYoverseAccountStore](i)

	return NewWorkflowFunc(func(ctx context.Context, command *LoginHoYoverseAccountCommand) (*HoYoverseAccountLoggedInEvent, error) {
		if err := store.Login(ctx, command.Email, command.Password); err != nil {
			return nil, fmt.Errorf("failed to login HoYoverse account: %w", err)
		}

		return &HoYoverseAccountLoggedInEvent{}, nil
	}), nil
}

package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/SlashNephy/auto-claimer/domain/entity"
	"github.com/SlashNephy/auto-claimer/repository/schema"
	"github.com/samber/do/v2"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

type RedeemedCodeRepository struct {
	db *gorm.DB
}

func NewRedeemedCodeRepository(i do.Injector) (*RedeemedCodeRepository, error) {
	db := do.MustInvoke[*gorm.DB](i)
	if err := db.AutoMigrate(&schema.RedeemedCode{}); err != nil {
		return nil, fmt.Errorf("failed to migrate RedeemedCode schema: %w", err)
	}

	return &RedeemedCodeRepository{
		db: db,
	}, nil
}

func (r *RedeemedCodeRepository) ListRedeemedCodes(ctx context.Context, account entity.Account) ([]string, error) {
	var redeemedCodes []*schema.RedeemedCode
	result := r.db.Where("game = ? AND account_id = ?", account.GetGame(), account.GetID()).Find(&redeemedCodes)
	if err := result.Error; err != nil {
		return nil, err
	}

	return lo.Map(redeemedCodes, func(code *schema.RedeemedCode, _ int) string {
		return code.Code
	}), nil
}

func (r *RedeemedCodeRepository) MarkCodeAsRedeemed(ctx context.Context, account entity.Account, code entity.Code) error {
	result := r.db.Create(&schema.RedeemedCode{
		Game:      code.GetGame(),
		Code:      code.GetCode(),
		AccountID: account.GetID(),
	})
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil
		}
		return err
	}

	return nil
}

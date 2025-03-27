package schema

import "github.com/SlashNephy/auto-claimer/domain/entity"

type RedeemedCode struct {
	Game      entity.Game `gorm:"uniqueIndex:idx_game_code_account"`
	Code      string      `gorm:"uniqueIndex:idx_game_code_account"`
	AccountID string      `gorm:"uniqueIndex:idx_game_code_account"`
}

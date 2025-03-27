package database

import (
	"github.com/glebarez/sqlite"
	"github.com/samber/do/v2"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	DSN string `env:"DSN" envDefault:"auto-claimer.db"`
}

func NewDatabase(i do.Injector) (*gorm.DB, error) {
	config := do.MustInvoke[*DatabaseConfig](i)

	return gorm.Open(
		sqlite.Open(config.DSN),
		&gorm.Config{
			TranslateError: true,
		},
	)
}

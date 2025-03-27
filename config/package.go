package config

import (
	"github.com/SlashNephy/auto-claimer/database"
	"github.com/SlashNephy/auto-claimer/pipeline"
	"github.com/samber/do/v2"
)

var Package = do.Package(
	do.Lazy(LoadConfig),

	// Config 構造体のフィールドを provide する
	// do.InvokeStruct はリフレクションを使用するため避けたい
	do.Lazy(func(i do.Injector) (*database.DatabaseConfig, error) {
		config, err := do.Invoke[*Config](i)
		if err != nil {
			return nil, err
		}
		return &config.Database, nil
	}),
	do.Lazy(func(i do.Injector) (*pipeline.BatchRedeemCodesConfig, error) {
		config, err := do.Invoke[*Config](i)
		if err != nil {
			return nil, err
		}
		return &config.Batch, nil
	}),
)

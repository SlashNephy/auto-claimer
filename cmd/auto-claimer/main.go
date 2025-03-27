package main

import (
	"context"
	"log/slog"

	"github.com/SlashNephy/auto-claimer/pipeline"
	"github.com/samber/do/v2"
)

func main() {
	ctx := context.Background()
	injector := do.New(Package)

	batch := do.MustInvoke[pipeline.BatchRedeemCodesPipeline](injector)
	if _, err := batch.Do(ctx, &pipeline.BatchRedeemCodesInput{}); err != nil {
		slog.ErrorContext(ctx, "failed to execute batch", slog.String("error", err.Error()))
	}
}

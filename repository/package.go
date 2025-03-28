package repository

import (
	"github.com/SlashNephy/auto-claimer/query"
	"github.com/SlashNephy/auto-claimer/workflow"
	"github.com/samber/do/v2"
)

var Package = do.Package(
	do.Lazy(NewHonkaiStarrailRepository),
	do.Bind[*HonkaiStarrailRepository, workflow.LoginHoYoverseAccountStore](),
	do.Bind[*HonkaiStarrailRepository, workflow.RedeemHonkaiStarrailCodeStore](),

	do.Lazy(NewEnneadRepository),
	do.Bind[*EnneadRepository, query.HonkaiStarrailCodeStore](),
	do.Bind[*EnneadRepository, query.GenshinImpactCodeStore](),
	do.Bind[*EnneadRepository, query.ZenlessZoneZeroCodeStore](),

	do.Lazy(NewRedeemedCodeRepository),
	do.Bind[*RedeemedCodeRepository, query.RedeemedCodeStore](),
	do.Bind[*RedeemedCodeRepository, workflow.MarkCodeAsRedeemedStore](),
)

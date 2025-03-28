package repository

import (
	"github.com/SlashNephy/auto-claimer/query"
	"github.com/SlashNephy/auto-claimer/workflow"
	"github.com/samber/do/v2"
)

var Package = do.Package(
	do.Lazy(NewHoYoverseRepository),
	do.Bind[*HoYoverseRepository, query.HonkaiStarrailStore](),
	do.Bind[*HoYoverseRepository, query.GenshinImpactStore](),
	do.Bind[*HoYoverseRepository, query.ZenlessZoneZeroStore](),
	do.Bind[*HoYoverseRepository, workflow.LoginHoYoverseAccountStore](),
	do.Bind[*HoYoverseRepository, workflow.RedeemHonkaiStarrailCodeStore](),
	do.Bind[*HoYoverseRepository, workflow.RedeemGenshinImpactCodeStore](),
	do.Bind[*HoYoverseRepository, workflow.RedeemZenlessZoneZeroCodeStore](),

	do.Lazy(NewEnneadRepository),
	do.Bind[*EnneadRepository, query.HonkaiStarrailCodeStore](),
	do.Bind[*EnneadRepository, query.GenshinImpactCodeStore](),
	do.Bind[*EnneadRepository, query.ZenlessZoneZeroCodeStore](),

	do.Lazy(NewRedeemedCodeRepository),
	do.Bind[*RedeemedCodeRepository, query.RedeemedCodeStore](),
	do.Bind[*RedeemedCodeRepository, workflow.MarkCodeAsRedeemedStore](),
)

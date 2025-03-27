package repository

import (
	"github.com/SlashNephy/auto-claimer/query"
	"github.com/SlashNephy/auto-claimer/workflow"
	"github.com/samber/do/v2"
)

var Package = do.Package(
	do.Lazy(NewHonkaiStarrailRepository),
	do.Bind[*HonkaiStarrailRepository, query.HonkaiStarrailCodeStore](),
	do.Bind[*HonkaiStarrailRepository, workflow.LoginHoYoverseAccountStore](),
	do.Bind[*HonkaiStarrailRepository, workflow.RedeemHonkaiStarrailCodeStore](),

	do.Lazy(NewRedeemedCodeRepository),
	do.Bind[*RedeemedCodeRepository, query.RedeemedCodeStore](),
	do.Bind[*RedeemedCodeRepository, workflow.MarkCodeAsRedeemedStore](),
)

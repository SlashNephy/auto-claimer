package query

import "github.com/samber/do/v2"

var Package = do.Package(
	do.Lazy(NewRedeemedCodeQuery),
	do.Bind[*RedeemedCodeQueryImpl, RedeemedCodeQuery](),
	do.Lazy(NewHonkaiStarrailCodeQuery),
	do.Bind[*HonkaiStarrailCodeQueryImpl, HonkaiStarrailCodeQuery](),
	do.Lazy(NewGenshinImpactCodeQuery),
	do.Bind[*GenshinImpactCodeQueryImpl, GenshinImpactCodeQuery](),
	do.Lazy(NewZenlessZoneZeroCodeQuery),
	do.Bind[*ZenlessZoneZeroCodeQueryImpl, ZenlessZoneZeroCodeQuery](),
)

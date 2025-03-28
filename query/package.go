package query

import "github.com/samber/do/v2"

var Package = do.Package(
	do.Lazy(NewRedeemedCodeQuery),
	do.Bind[*RedeemedCodeQueryImpl, RedeemedCodeQuery](),
	do.Lazy(NewHonkaiStarrailQuery),
	do.Bind[*HonkaiStarrailQueryImpl, HonkaiStarrailQuery](),
	do.Lazy(NewGenshinImpactQuery),
	do.Bind[*GenshinImpactQueryImpl, GenshinImpactQuery](),
	do.Lazy(NewZenlessZoneZeroQuery),
	do.Bind[*ZenlessZoneZeroQueryImpl, ZenlessZoneZeroQuery](),
)

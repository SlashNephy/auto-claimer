package pipeline

import "github.com/samber/do/v2"

var Package = do.Package(
	do.Lazy(NewRedeemHonkaiStarrailCodesPipeline),
	do.Lazy(NewRedeemGenshinImpactCodesPipeline),
	do.Lazy(NewRedeemZenlessZoneZeroCodesPipeline),
	do.Lazy(NewBatchRedeemCodesPipeline),
)

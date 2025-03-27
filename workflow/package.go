package workflow

import "github.com/samber/do/v2"

var Package = do.Package(
	do.Lazy(NewLoginHoYoverseAccountWorkflow),
	do.Lazy(NewRedeemHonkaiStarrailCodeWorkflow),
	do.Lazy(NewNotifyHoYoverseCodeRedeemedWorkflow),
)

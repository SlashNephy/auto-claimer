package query

import "github.com/samber/do/v2"

var Package = do.Package(
	do.Lazy(NewHonkaiStarrailCodeQuery),
	do.Bind[*HonkaiStarrailCodeQueryImpl, HonkaiStarrailCodeQuery](),
)

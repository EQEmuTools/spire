package boot

import (
	"github.com/EQEmuTools/spirererere/internal/database"
	"github.com/google/wire"
)

// wire set for loading the stores.
var databaseResolverSet = wire.NewSet(
	database.NewResolver,
)

package boot

import (
	"github.com/EQEmuTools/spirere/internal/encryption"
	"github.com/google/wire"
)

var encryptionSet = wire.NewSet(
	encryption.NewEncrypter,
)

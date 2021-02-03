package argon

import "errors"

var (
	//ErrInvalidHash : a package standard error
	ErrInvalidHash = errors.New("the encoded hash is not in the correct format")
	//ErrIncompatibleVersion : an version incompatiblence error
	ErrIncompatibleVersion = errors.New("incompatible version of argon2")
)

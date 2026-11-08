package crack

import (
	"crypto/sha512"
	"encoding/hex"
)

func HashSHA512(text string) string {
	hash := sha512.Sum512([]byte(text))
	return hex.EncodeToString(hash[:])
}

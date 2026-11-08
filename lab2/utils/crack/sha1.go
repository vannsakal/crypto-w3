package crack

import (
	"crypto/sha1"
	"encoding/hex"
)

func HashSHA1(text string) string {
	hash := sha1.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

package auth

import (
	"crypto/rand"
	"encoding/hex"
)

func MakeRefreshToken() string {
	refTokenBytes := make([]byte, 32)
	_, err := rand.Read(refTokenBytes)
	if err != nil {
		return err.Error()
	}

	return hex.EncodeToString(refTokenBytes)
}

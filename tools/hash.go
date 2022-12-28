package tools

import (
	"crypto/sha256"
	"fmt"
)

type Hash string

func Digest(content []byte) Hash {
	return Hash(fmt.Sprintf("sha256:%x", hash(content)))
}

// For the moment, only SHA256 is supported for the simplicity
func hash(content []byte) [sha256.Size]byte {
	return sha256.Sum256(content)
}

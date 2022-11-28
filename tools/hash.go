package tools

import "crypto/sha256"

type Hash string

func Digest(content []byte) Hash {
	return Hash(hash(content))
}

func hash(content []byte) []byte {
	hash := sha256.Sum256(content)
	return hash[:]
}

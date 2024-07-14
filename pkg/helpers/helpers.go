package helpers

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

func Sha256Digest(size int32) string {
	b := make([]byte, size)
	id := sha256.New()
	_, _ = rand.Read(b)
	id.Write(b)
	newID := fmt.Sprintf("%x", id.Sum(nil))
	return newID
}

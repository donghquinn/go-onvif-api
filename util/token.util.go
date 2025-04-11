package util

import (
	"crypto/rand"
	"fmt"
)

func CreateToken() string {
	b := make([]byte, 4)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

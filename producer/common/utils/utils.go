package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

func KeyGen(s string) string {
	h := sha1.Sum([]byte(s)) 
	return fmt.Sprintf("word_count:%s", hex.EncodeToString(h[: 8])) 
}
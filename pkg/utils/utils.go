package utils

import (
	"crypto/sha256"
	"encoding/base64"
)

func EncodeURL(url string) string {
	hash := sha256.Sum256([]byte(url))
	return base64.URLEncoding.EncodeToString(hash[:])[:8]
}

func DecodeURL(shortURL string) string {
	hash, err := base64.URLEncoding.DecodeString(shortURL + "==")
	if err != nil {
		return ""
	}
	return string(hash)
}

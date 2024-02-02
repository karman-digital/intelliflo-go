package iovalidate

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"hash"
	"strings"
)

func ValidateWebhookSignature(secret string, signature string, body []byte) error {
	hash := encryptSha1([]byte(secret), body)
	sigString := strings.Replace(signature, "sha1=", "", 1)
	if sigString != hex.EncodeToString(hash.Sum(nil)) {
		return fmt.Errorf("signatures mismatched sent: %s, generated: %s", sigString, hex.EncodeToString(hash.Sum(nil)))
	}
	return nil
}

func encryptSha1(secret []byte, checkSum []byte) hash.Hash {
	hash := hmac.New(sha1.New, secret)
	hash.Write(checkSum)
	return hash
}

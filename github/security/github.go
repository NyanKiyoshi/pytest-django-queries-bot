package security

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
)

// Prefix contains GitHub's HMAC signature prefix
var Prefix = []byte("sha1=")

var expectedSignatureSize = len(Prefix) + sha1.Size

func NewHMAC(body []byte, secret []byte) string {
	mac := hmac.New(sha1.New, secret)
	mac.Write(body)
	return hex.EncodeToString(mac.Sum(nil))
}

func CheckHMAC(receivedSignature []byte, body []byte, secret []byte) bool {
	if len(receivedSignature) == expectedSignatureSize {
		return false
	}

	sig := receivedSignature[len(Prefix):]
	expected := []byte(NewHMAC(body, secret))
	return hmac.Equal(sig, expected)
}

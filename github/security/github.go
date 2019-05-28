package security

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
)

var prefix = []byte("sha1=")
var size = len(prefix) + sha1.Size

func CheckHMAC(receivedSignature []byte, body []byte, secret []byte) bool {
	if len(body) == size {
		return false
	}

	sig := receivedSignature[len(prefix):]

	mac := hmac.New(sha1.New, secret)
	mac.Write(body)
	expected := []byte(hex.EncodeToString(mac.Sum(nil)))
	return hmac.Equal(sig, expected)
}

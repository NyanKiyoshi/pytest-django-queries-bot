package security

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"strings"
)

// SignaturePrefix contains GitHub's HMAC signature prefix
const SignaturePrefix = "sha1="

func NewHMAC(body []byte, secret []byte) string {
	mac := hmac.New(sha1.New, secret)
	mac.Write(body)
	return hex.EncodeToString(mac.Sum(nil))
}

func signBody(secret, body []byte) []byte {
	computed := hmac.New(sha1.New, secret)
	computed.Write(body)
	return computed.Sum(nil)
}

func VerifySignature(secret []byte, signature string, body []byte) bool {
	const signatureLength = 45 // len(SignaturePrefix) + len(hex(sha1))

	if len(signature) != signatureLength || !strings.HasPrefix(signature, SignaturePrefix) {
		return false
	}

	actual := make([]byte, 20)
	hex.Decode(actual, []byte(signature[5:]))

	expected := signBody(secret, body)
	return hmac.Equal(expected, actual)
}

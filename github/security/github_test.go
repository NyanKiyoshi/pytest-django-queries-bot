package security

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var message = "Hello World"
var signature = "sha1=858da8837b87f04b052c0f6e954c3f7bbe081164"

var hmacflags = []struct {
	secret string
	result bool
}{
	{"secret", true},
	{"invalid", false},
	{"sha1=858da8837b87f04b052c0f6e954c3f7bbe081165", false},
	{"858da8837b87f04b052c0f6e954c3f7bbe081164", false},
}

func TestCheckHMAC(t *testing.T) {
	var ok bool
	for _, tt := range hmacflags {
		t.Run(tt.secret, func(t *testing.T) {
			ok = VerifySignature([]byte(tt.secret), signature, []byte(message))
			assert.Equal(t, tt.result, ok, "The signatures should have matched")
		})
	}
}

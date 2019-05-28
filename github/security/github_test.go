package security

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var message = []byte("Hello World")
var signature = []byte("sha1=858da8837b87f04b052c0f6e954c3f7bbe081164")

var hmacflags = []struct {
	secret string
	result bool
}{
	{"secret", true},
	{"invalid", false},
}
func TestCheckHMAC(t *testing.T) {
	for _, tt := range hmacflags {
		t.Run(tt.secret, func(t *testing.T) {
			ok := CheckHMAC(signature, message, []byte(tt.secret))
			assert.Equal(t, tt.result, ok, "The signatures should have matched")
		})
	}
}

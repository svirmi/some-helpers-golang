package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDebugJWT(t *testing.T) {
	var hf HelpersFunctions = HelpersFunctions{}
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"

	header, payload, err := hf.DebugJWT(tokenString)

	expectedHeader := "{\"alg\":\"HS256\",\"typ\":\"JWT\"}"
	expectedPayload := "{\"iat\":1516239022,\"name\":\"John Doe\",\"sub\":\"1234567890\"}"

	assert.Nil(t, err)
	assert.Equal(t, expectedHeader, header)
	assert.Equal(t, expectedPayload, payload)
}

func TestDebugJWTWithInvalidToken(t *testing.T) {
	var hf HelpersFunctions = HelpersFunctions{}
	tokenString := "xxxxx.yyyyy.zzzzz"

	header, payload, err := hf.DebugJWT(tokenString)

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "error parsing token")
	assert.Empty(t, header)
	assert.Empty(t, payload)
}

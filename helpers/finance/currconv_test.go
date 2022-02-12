package finance

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFinanceFunctions(t *testing.T) {
	ff := NewFinanceFunctions("SomeApiUrl", "SomeApiKey")

	assert.Equal(t, "SomeApiUrl", ff.ApiUrl)
	assert.Equal(t, "SomeApiKey", ff.ApiKey)
}

func TestNewFinanceFunctionsWithEmptyApiUrl(t *testing.T) {
	ff := NewFinanceFunctions("", "SomeApiKey")

	assert.NotEmpty(t, ff.ApiUrl)
	assert.Contains(t, ff.ApiUrl, "fcsapi.com")
}

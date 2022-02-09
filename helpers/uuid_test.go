package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// creates a instance of the structure to be used as a receiver
var hf HelpersFunctions = HelpersFunctions{}

func TestNewUuidWithHyphen(t *testing.T) {
	uuidWithHyphen := hf.NewUuid(false)

	assert.Len(t, uuidWithHyphen, 36)
	assert.Contains(t, uuidWithHyphen, "-")
}

func TestNewUuidWithoutHyphen(t *testing.T) {
	uuidWithHyphen := hf.NewUuid(true)

	assert.Len(t, uuidWithHyphen, 32)
	assert.NotContains(t, uuidWithHyphen, "-")
}

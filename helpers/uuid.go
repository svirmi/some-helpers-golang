package helpers

import (
	"strings"

	"github.com/google/uuid"
)

func (hf *HelpersFunctions) NewUuid(withoutHyphen bool) string {
	uuidWithHyphen := uuid.New()

	if withoutHyphen {
		return strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	}

	return uuidWithHyphen.String()
}

package helpers

import (
	"encoding/json"
	"fmt"

	"github.com/golang-jwt/jwt"
)

func (hf *HelpersFunctions) DebugJWT(tokenString string) (string, string, error) {
	parser := jwt.Parser{}
	token, _, err := parser.ParseUnverified(tokenString, jwt.MapClaims{})

	if err != nil {
		return "", "", fmt.Errorf("error parsing token: %s", err.Error())
	}

	header, _ := json.Marshal(token.Header)
	payload, _ := json.Marshal(token.Claims)

	return string(header), string(payload), nil
}

package helpers

type Interface interface {
	NewUuid(withoutHyphen bool) string
	DebugJWT(tokenString string) (string, string, error)
}

type HelpersFunctions struct{}

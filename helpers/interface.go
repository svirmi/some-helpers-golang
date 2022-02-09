package helpers

type Interface interface {
	NewUuid(withoutHyphen bool) string
}

type HelpersFunctions struct{}

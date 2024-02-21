package instances

import "net/http"

type IsRequirement interface {
	Expectation(r http.Response) bool
	Report() string
}

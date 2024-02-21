package instances

type IsRequirement interface {
	Expectation(interface{}) bool
	Report() string
}

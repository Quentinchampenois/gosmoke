package instances

type Requirement struct {
	ID         uint
	Name       string
	InstanceID uint
	ContainsID uint
	Contains   Contains
}

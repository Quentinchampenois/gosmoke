package instances

import "gorm.io/gorm"

type Instance struct {
	gorm.Model
	Name         string
	URL          string
	Requirements []Requirement
}

package instances

import (
	"fmt"
	"gorm.io/gorm"
	"strings"
)

type Contains struct {
	gorm.Model
	InstanceID uint
	Name       string
	Expected   string
	IsMet      bool
}

func (c *Contains) Expectation(el interface{}) bool {
	c.IsMet = strings.Contains(el.(string), c.Expected)
	return c.IsMet
}

func (c *Contains) Report() string {
	if c.IsMet {
		return fmt.Sprintf("'%s' is valid", c.Name)
	} else {
		return fmt.Sprintf("'%s' is invalid", c.Name)
	}
}

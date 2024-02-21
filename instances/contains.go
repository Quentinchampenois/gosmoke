package instances

import (
	"fmt"
	"gorm.io/gorm"
	"io"
	"net/http"
	"strings"
)

type Contains struct {
	gorm.Model
	InstanceID uint
	Name       string
	Expected   string
	IsMet      bool
}

func (c *Contains) Expectation(el http.Response) bool {
	content, err := io.ReadAll(el.Body)
	if err != nil {
		c.IsMet = false
		return c.IsMet
	}

	fmt.Println(strings.Contains(string(content), c.Expected))
	c.IsMet = strings.Contains(string(content), c.Expected)
	return c.IsMet
}

func (c *Contains) Report() string {
	if c.IsMet {
		return fmt.Sprintf("'%s' is valid", c.Name)
	} else {
		return fmt.Sprintf("'%s' is invalid", c.Name)
	}
}

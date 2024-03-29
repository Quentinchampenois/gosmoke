package instances

import (
	"fmt"
	"gorm.io/gorm"
	"net/http"
)

type Request struct {
	gorm.Model
	InstanceID uint
	Name       string
	StatusCode int
	IsMet      bool
}

func (r *Request) Expectation(el http.Response) bool {
	r.IsMet = el.StatusCode == r.StatusCode
	return r.IsMet
}

func (r *Request) Report() string {
	if r.IsMet {
		return fmt.Sprintf("'%s' is valid", r.Name)
	} else {
		return fmt.Sprintf("'%s' is invalid", r.Name)
	}
}

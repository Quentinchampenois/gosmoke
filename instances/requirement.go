package instances

import "net/http"

type Requirement struct {
	ID         uint
	Name       string
	IsMet      bool
	InstanceID uint
	ContainsID uint
	Contains   Contains
	RequestID  uint
	Request    Request
}

func (r *Requirement) Expectation(el http.Response) bool {
	if r.ContainsID != 0 {
		return r.Contains.Expectation(el)
	} else {
		return r.Request.Expectation(el)
	}
}

func (r *Requirement) Report() string {
	if r.ContainsID != 0 {
		return r.Contains.Report()
	} else {
		return r.Request.Report()
	}
}

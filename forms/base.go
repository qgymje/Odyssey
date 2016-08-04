package forms

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

type Base struct {
	Valid *validation.Validation
	Msg   *errmsg
}

func newBase() *Base {
	return &Base{
		Valid: &validation.Validation{},
		Msg:   newErrmsg(),
	}
}

type errmsg struct {
	msg map[string][]string
}

func newErrmsg() *errmsg {
	return &errmsg{
		msg: map[string][]string{},
	}
}

func (s *errmsg) formatBindError(err error) {
	s.setError("params", err.Error())
}

func (s *errmsg) formatBindError2(errs []*gin.Error) {
	for _, e := range errs {
		s.setError("binding", e.Err.Error())
	}
}

func (s *errmsg) setError(key, msg string) {
	s.msg[key] = append(s.msg[key], msg)
}

// maybe can implement the Error() string method
func (s *errmsg) ErrorMsg() map[string][]string {
	return s.msg
}

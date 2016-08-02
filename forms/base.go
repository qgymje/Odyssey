package forms

import "github.com/astaxie/beego/validation"

type Base struct {
	*validation.Validation
	*errmsg
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

func (s *errmsg) setError(key, msg string) {
	s.msg[key] = append(s.msg[key], msg)
}

func (s *errmsg) ErrorMsg() map[string][]string {
	return s.msg
}

package forms

type errmsg struct {
	msg map[string][]string
}

func newErrmsg() errmsg {
	return errmsg{
		msg: map[string][]string{},
	}
}

func (s errmsg) formatBindError(msg string) {

}

func (s errmsg) setError(key, msg string) {
	s.msg[key] = append(s.msg[key], msg)
}

func (s errmsg) ErrorMsg() map[string][]string {
	return s.msg
}

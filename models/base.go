package models

type Base struct {
	_exists  bool
	_created bool
}

func (b Base) Exists() bool {
	return b._exists
}

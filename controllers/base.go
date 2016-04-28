package controllers

type Base struct{}

func (b *Base) Authorization() error {
	return nil
}

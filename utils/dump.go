package utils

import (
	"fmt"

	"github.com/kr/pretty"
)

func Dump(x ...interface{}) {
	fmt.Printf("%# v", pretty.Formatter(x))
}

func Sdump(x ...interface{}) string {
	return fmt.Sprintf("%# v", pretty.Formatter(x))
}

package utils

import (
	"fmt"

	"github.com/kr/pretty"
)

func Dump(x ...interface{}) {
	fmt.Printf("%# v", pretty.Formatter(x))
}

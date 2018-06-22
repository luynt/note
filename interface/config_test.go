package _interface

import (
	"fmt"
	"testing"
	config2 "github.com/luynt/note/interface/config"
)

type Custom string

func (t Custom) String() string  {
	return string(t)
}

func Test(t *testing.T) {
	cfn := config2.NewConfig()
		fmt.Println(cfn.String())

	var x  Custom="xsdasdasd"
	config2.Benmark(x)
}

package config

import "fmt"

type Config interface {
	String() string
}

func NewConfig() Config {
	return &convert{}
}

type convert struct {
}

func (c *convert) String() string {
	return "This is convert"
}

func Benmark(cfn Config) {
	fmt.Println(cfn.String())

}

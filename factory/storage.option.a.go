package factory

import (
	"fmt"
)

// EnumOptionA option for factory
const EnumOptionA = "OptionA"

type option struct{}

func init() {
	addOption(&option{})
	fmt.Println(EnumOptionA + " added to factory")
}

func (o *option) evalOption(params ...interface{}) bool {
	if params[0] == EnumOptionA {
		return true
	}
	return false
}

func (o *option) getOption() IOption {
	return &option{}
}

// Do implements IOption method
func (o *option) Do() {
	fmt.Println("Do: Option A")
}

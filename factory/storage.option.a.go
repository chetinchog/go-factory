package factory

import (
	"fmt"
)

// EnumOptionA option for factory
const EnumOptionA = "OptionA"

type optionA struct{}

func init() {
	addOption(&optionA{})
	fmt.Println(EnumOptionA + " added to factory")
}

func (o *optionA) evalOption(params ...interface{}) bool {
	if params[0] == EnumOptionA {
		return true
	}
	return false
}

// Do implements IOption method
func (o *optionA) Do() {
	fmt.Println("Do: Option A")
}

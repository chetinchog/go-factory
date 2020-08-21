package options

import (
	"fmt"
	"gofp/factory"
)

// EnumOptionA option for factory
const EnumOptionA = "OptionA"

type optionA struct{}

// Add itself to the factory
func init() {
	factory.AddOption(&optionA{})
	fmt.Println(EnumOptionA + " added to factory")
}

// EvalOption returns true when params correspond to this option
func (o *optionA) EvalOption(params ...interface{}) bool {
	if params[0] == EnumOptionA {
		return true
	}
	return false
}

// Do implements IOption method
func (o *optionA) Do() {
	fmt.Println("Do: Option A")
}

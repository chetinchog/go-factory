package options

import (
	"fmt"
	"gofp/factory"
)

// EnumOptionB option for factory
const EnumOptionB = "OptionB"

type optionB struct{}

// Add itself to the factory
func init() {
	factory.AddOption(&optionB{})
	fmt.Println(EnumOptionB + " added to factory")
}

// EvalOption returns true when params correspond to this option
func (o *optionB) EvalOption(params ...interface{}) bool {
	if params[0] == EnumOptionB {
		return true
	}
	return false
}

// Do implements IOption method
func (o *optionB) Do() {
	fmt.Println("Do: Option B")
}

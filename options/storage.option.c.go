package options

import (
	"fmt"
	"gofp/factory"
)

// EnumOptionC option for factory
const EnumOptionC = "OptionC"

type optionC struct{}

// Add itself to the factory
func init() {
	factory.AddOption(&optionC{})
	fmt.Println(EnumOptionC + " added to factory")
}

// EvalOption returns true when params correspond to this option
func (o *optionC) EvalOption(params ...interface{}) bool {
	if params[0] == EnumOptionC {
		return true
	}
	return false
}

// Do implements IOption method
func (o *optionC) Do() {
	fmt.Println("Do: Option C")
}

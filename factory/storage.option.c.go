package factory

import (
	"fmt"
)

// EnumOptionC option for factory
const EnumOptionC = "OptionC"

type optionC struct{}

func init() {
	addOption(&optionC{})
	fmt.Println(EnumOptionC + " added to factory")
}

func (o *optionC) evalOption(params ...interface{}) bool {
	if params[0] == EnumOptionC {
		return true
	}
	return false
}

// Do implements IOption method
func (o *optionC) Do() {
	fmt.Println("Do: Option C")
}

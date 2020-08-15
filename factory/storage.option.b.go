package factory

import (
	"fmt"
)

// EnumOptionB option for factory
const EnumOptionB = "OptionB"

type optionB struct{}

func init() {
	addOption(&optionB{})
	fmt.Println(EnumOptionB + " added to factory")
}

func (o *optionB) evalOption(params ...interface{}) bool {
	if params[0] == EnumOptionB {
		return true
	}
	return false
}

// Do implements IOption method
func (o *optionB) Do() {
	fmt.Println("Do: Option B")
}

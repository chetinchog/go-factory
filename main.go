package main

import (
	"factory/factory"
	"fmt"
)

func main() {
	fmt.Println()
	fmt.Println("Go Factory")

	option, _ := factory.GetFactory().GetOption(factory.EnumOptionA)
	option.Do()

	option, _ = factory.GetFactory().GetOption(factory.EnumOptionB)
	option.Do()

	option, _ = factory.GetFactory().GetOption(factory.EnumOptionC)
	option.Do()
}

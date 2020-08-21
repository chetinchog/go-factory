package main

import (
	"fmt"
	"gofp/factory"
	"gofp/options"
)

func main() {
	fmt.Println()
	fmt.Println("Go Factory")

	option, _ := factory.GetFactory().GetOption(options.EnumOptionA)
	option.Do()

	option, _ = factory.GetFactory().GetOption(options.EnumOptionB)
	option.Do()

	option, _ = factory.GetFactory().GetOption(options.EnumOptionC)
	option.Do()
}

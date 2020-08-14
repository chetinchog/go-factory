package main

import (
	"factory/factory"
	"fmt"
)

func main() {
	fmt.Println()
	fmt.Println("Go Factory")

	option, err := factory.GetFactory().GetOption(factory.EnumOptionA)
	if err != nil {
		fmt.Printf("Error: %v", err)
	} else {
		option.Do()
	}
}

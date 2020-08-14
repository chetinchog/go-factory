package main

import (
	"factory/factory"
	"fmt"
)

func main() {
	fmt.Println("Go Factory")

	_, err := factory.GetFactory().GetOption(factory.EnumOptionNil)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
}

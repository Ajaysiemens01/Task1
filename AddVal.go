package main

import (
	"fmt"
)

func (inputString *str) AddValue(toAdd interface{}) (string, error) {
	fmt.Println("-----------------------------------------------------------------")

	fmt.Println("Input String : ", inputString.Input)

	// Check the type of toAdd and handle accordingly
	var stringAfterAdd string
	switch val := toAdd.(type) {
	case int: // If it's an int
		stringAfterAdd = inputString.Input + fmt.Sprintf("%d", val)
	case rune: // If it's a rune
		stringAfterAdd = inputString.Input + string(val)
	case string:
		// If it's already a string
		stringAfterAdd = inputString.Input + val
	default:
		return "", fmt.Errorf("unsupported type: %T", val)
	}

	fmt.Printf("After adding value: %s\n", stringAfterAdd)
	return stringAfterAdd, nil

}

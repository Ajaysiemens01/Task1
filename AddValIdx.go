package main

import (
	"fmt"
)

func (inputString *str) AddValueAtIndex(toAdd interface{}, position int) (string, error) {

	fmt.Println("-----------------------------------------------------------------")

	fmt.Println("Input String : ", inputString.Input)
	if position < 0 || position > len(inputString.Input) {
		return "", fmt.Errorf("invalid position: %d", position)
	}

	var stringAfterAdd string
	switch v := toAdd.(type) {
	case string: // If it's a string
		stringAfterAdd = inputString.Input[:position] + v + inputString.Input[position:]
	case rune: // If it's a rune
		stringAfterAdd = inputString.Input[:position] + string(v) + inputString.Input[position:]
	case int: // If it's an int
		stringAfterAdd = inputString.Input[:position] + fmt.Sprintf("%d", v) + inputString.Input[position:]
	default:
		return "", fmt.Errorf("unsupported type: %T", toAdd)
	}
	fmt.Printf("After adding value \"%s\" at position \"%d\": %s\n", toAdd, position, stringAfterAdd)

	return stringAfterAdd, nil
}

package main

import (
	"fmt"
)

func (inputString *str) RemoveEmptyValues() (string, error) {

	fmt.Println("-----------------------------------------------------------------")

	fmt.Println("Input String : ", inputString.Input)
	inputStringArray := []rune(inputString.Input)
	pos := 0

	// Iterate through the string to remove the EmptyValues
	for current := 0; current < len(inputStringArray); current++ {
		// Iterate through the string to check if the EmptyValues is not present add the character to the new string

		if inputStringArray[current] != ' ' {
			inputStringArray[pos] = inputStringArray[current]
			pos++
		}

	}

	stringAfterRemovingEmptyValues := string(inputStringArray[0:pos])

	if stringAfterRemovingEmptyValues == inputString.Input {
		return "", fmt.Errorf("no empty values found")
	} else {
		fmt.Println("After removing EmptyValues: ", stringAfterRemovingEmptyValues)
	}

	//we can also use strings.ReplaceAll(input, string(to_delete), "")
	return stringAfterRemovingEmptyValues, nil
}

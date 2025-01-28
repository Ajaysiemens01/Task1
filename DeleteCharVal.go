package main

import (
	"fmt"
)

func (inputString *str) DeleteCharValue(toDelete rune) (string, error) {

	fmt.Println("-----------------------------------------------------------------")

	fmt.Println("Input String : ", inputString.Input)
	inputStringArray := []rune(inputString.Input)
	pos := 0

	// Iterate through the string to shift the characters to the left and leave the character to delete

	for position := 0; position < len(inputStringArray); position++ {
		if inputStringArray[position] == toDelete {
			continue
		}
		inputStringArray[pos] = inputStringArray[position]
		pos++
	}

	stringAfterDelete := string(inputStringArray[0:pos])
	if stringAfterDelete == inputString.Input {
		return "", fmt.Errorf("no character \"%c\" found", toDelete)
	} else {
		fmt.Printf("After deleting character \"%c\": %s\n", toDelete, stringAfterDelete)
		return stringAfterDelete, nil
	}

	//we can also use strings.ReplaceAll(input, string(to_delete), "")

}

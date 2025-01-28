package main

import (
	"fmt"
)

func (inputString *str) RemoveDuplicates() (string, error) {

	fmt.Println("-----------------------------------------------------------------")

	fmt.Println("Input String : ", inputString.Input)
	inputStringArray := []rune(inputString.Input)
	pos := 0

	// Iterate through the string to remove the duplicates
	for curr := 0; curr < len(inputStringArray); curr++ {
		prev := 0
		// If the character is space, add it to the new string
		if inputStringArray[curr] == ' ' {
			inputStringArray[pos] = inputStringArray[curr]
			pos++
			continue
		}
		// Iterate through the string to check if the character is already present
		for prev < curr {
			if inputStringArray[curr] == inputStringArray[prev] {
				break
			}
			prev++
		}
		if prev == curr {
			inputStringArray[pos] = inputStringArray[curr]
			pos++
		}
	}
	stringAfterRemovingDuplicates := string(inputStringArray[0:pos])

	if stringAfterRemovingDuplicates == inputString.Input {
		return "", fmt.Errorf("no duplicates found")
	} else {
		fmt.Println("After removing Duplicates: ", stringAfterRemovingDuplicates)
	}

	//we can also use strings.ReplaceAll(input, string(to_delete), "")
	return stringAfterRemovingDuplicates, nil
}

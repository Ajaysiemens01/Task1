package main

import "fmt"

type strOperations interface {
	RemoveDuplicates() (string, error)
	AddValueAtIndex(toAdd interface{}, pos int) (string, error)
	DeleteCharValue(toDelete rune) (string, error)
	RemoveEmptyValues() (string, error)
	AddValue(toAdd interface{}) (string, error)
}

type str struct {
	Input string
}

func main() {

	fmt.Println("Test Case 1 :")
	inputString1 := str{
		Input: "hello ",
	}

	TestRemovingDuplicates(&inputString1)
	TestAddingValueAtIndex(&inputString1, ".", 6)
	TestDeletingCharValue(&inputString1, 'l')
	TestRemoveEmptyValues(&inputString1)
	TestAddingValue(&inputString1, "world")

	fmt.Println()
	fmt.Println()
	fmt.Println("Test Case 2 :")
	inputString2 := str{
		Input: "ajay kuma",
	}

	TestRemovingDuplicates(&inputString2)
	TestAddingValueAtIndex(&inputString2, "r", 9)
	TestDeletingCharValue(&inputString2, 'a')
	TestRemoveEmptyValues(&inputString2)
	TestAddingValue(&inputString2, 'r')

	fmt.Println()
	fmt.Println()
	fmt.Println("Test Case 3 :")
	inputString3 := str{
		Input: "a bc de",
	}

	TestRemovingDuplicates(&inputString3)
	TestAddingValueAtIndex(&inputString3, "R", 10)
	TestDeletingCharValue(&inputString3, 'f')
	TestRemoveEmptyValues(&inputString3)
	TestAddingValue(&inputString3, 4)

	fmt.Println()
	fmt.Println()

	fmt.Println("Test Case 4 :")
	inputString4 := str{
		Input: "123",
	}
	TestAddingValue(&inputString4, 4.01)
	TestAddingValue(&inputString4, "45")
	TestAddingValue(&inputString4, '4')

}

package main

import (
	"fmt"
)

// Function to test the RemoveDuplicates function
func TestRemovingDuplicates(s strOperations) {
	_, err := s.RemoveDuplicates()
	if err != nil {
		fmt.Println(err)
	}
}


// Function to test the AddValueAtIndex function
func TestAddingValueAtIndex(s strOperations, toAdd interface{}, position int) {
	_, err := s.AddValueAtIndex(toAdd, position)
	if err != nil {
		fmt.Println(err)
	}
}

//  Function to test the DeleteValue function
func TestDeletingCharValue(s strOperations, toDelete rune) {
	_, err := s.DeleteCharValue(toDelete)
	if err != nil {
		fmt.Println(err)
	}
}

// Function to test the RemoveEmptyValues function
func TestRemoveEmptyValues(s strOperations) {
	_, err := s.RemoveEmptyValues()
	if err != nil {
		fmt.Println(err)
	}
}

// Function to test the AddValue function
func TestAddingValue(s strOperations, toAdd interface{}) {
	_, err := s.AddValue(toAdd)
	if err != nil {
		fmt.Println(err)
	}
}

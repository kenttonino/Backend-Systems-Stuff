package data_structures

import (
	"data-structure-in-go/libs"
	"fmt"
)

/*
	Reference: https://www.geeksforgeeks.org/dsa/introduction-to-arrays-data-structure-and-algorithm-tutorials/

	Array:
		- Collection of items stored in contiguous memory locations.

	Array Element
		- Elements stored in an array.

	Array Index
		- Used to access an element.
		- Indexes starts in 0.
*/

func Array0001() {
	arrayOfInt := []int{1, 2, 3, 4, 5}
	arrayOfCharacters := "abcde"

	for i := range 5 {
		fmt.Print(arrayOfInt[i])
		libs.NewTab(1)
	}

	libs.NewLine(2)

	for i := range 5 {
		fmt.Print(arrayOfCharacters[i])
		libs.NewTab(1)
	}
}

package libs

import "fmt"

func NewLine(num int) {
	for range num {
		fmt.Println()
	}
}

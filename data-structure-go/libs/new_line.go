package libs

import "fmt"

func NewLine(lines int) {
	for range lines {
		fmt.Println()
	}
}

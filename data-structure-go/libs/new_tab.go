package libs

import "fmt"

func NewTab(tabs int) {
	for range tabs {
		fmt.Print("\t")
	}
}

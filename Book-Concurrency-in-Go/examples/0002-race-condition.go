package examples

import (
	"fmt"
	"time"
)

func Ex0002RaceCondition() {
	var data int

	go func() {
		data++
	}()

	time.Sleep(1 * time.Second)

	if data == 0 {
		fmt.Printf("Data: %v", data)
	}
}

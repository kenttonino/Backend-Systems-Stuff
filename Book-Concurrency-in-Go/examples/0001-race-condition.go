package examples

import "fmt"

/*
Race Condition:
  - Occurs when two or more operations must execute in the correct order, but
    the program has not been written so that this order is guaranteed to be
    maintained.

Data Race
  - Where one concurrent operation attempts to read a variable while at some
    undetermined time another concurrent operation is attempting to write to
    the same variable.
*/
func Ex0001RaceCondition() {
	var data int

	// go (goroutine) is a keyword to run a function concurrently.
	go func() {
		data++
	}()

	if data == 0 {
		fmt.Printf("Data: %v", data)
	}
}

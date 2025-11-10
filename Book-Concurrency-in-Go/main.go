package main

import (
	"book-concurrency-in-go/examples"
	"book-concurrency-in-go/libs"
)

func main() {
	libs.NewLine(1)
	examples.Ex0001RaceCondition()
	libs.NewLine(2)
}

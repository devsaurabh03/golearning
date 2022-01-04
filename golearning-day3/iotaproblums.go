package golearningday3

import "fmt"

func PrintIOTAExamples() {

	const number = 100

	const (
		UTC1 = number
		UTC2 = number - iota
		UTC3
	)
	// constant seriues generator from negative to positive
	fmt.Println(UTC1, UTC2, UTC3)
}

package golearningday6

import "fmt"

func RangeLoops(words []string) {

	for index, value := range words {

		fmt.Printf("%d -> %s \n", index, value)
	}

	fmt.Println()

}

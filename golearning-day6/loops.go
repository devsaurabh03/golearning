package golearningday6

import "fmt"

func LoopPrograms() {

	var sum int

	for i := 0; i < 1000000; i++ {

		if i > 10000 {
			break
		}

		sum = sum + 1

	}

	fmt.Println(sum)

}

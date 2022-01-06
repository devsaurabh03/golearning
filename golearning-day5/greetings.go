package golearningday5

import (
	"fmt"
	"time"
)

func Greetfortheday() {

	hour := time.Now().Hour()

	fmt.Print(hour)

	switch {
	case hour > 12 && hour < 6:
		{
			fmt.Print("good afternoon")
		}
	default:
		fmt.Print("good day")
	}

}

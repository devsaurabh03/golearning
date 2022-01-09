package golearningday7

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func GuesssTheNumber() {

	//take input from user

	// match the with the randon genertor

	// return the result

	var args []string = os.Args

	if args[1] != "" {

		fmt.Println("please pass argument and try again")
		return
	}

	time.Now()

	guess := rand.Int63n(10)

	for i := int64(0); i < guess; i++ {

		if guess == rand.Int63() {

			fmt.Print("please move ahead")
		}

	}

}

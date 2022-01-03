package main

import (
	"fmt"

	golearning "github.com/golearning/snow/golearning-day1"
	golearningday2 "github.com/golearning/snow/golearning-day2"
)

func main() {

	fmt.Println(golearning.GetExample())
	dir, path := golearningday2.SeparatePath("virat/kohli/boy")

	fmt.Println(dir)
	fmt.Println(path)

}

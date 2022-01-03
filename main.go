package main

import (
	"fmt"
	"unicode/utf8"

	golearningday2 "github.com/golearning/snow/golearning-day2"
)

func main() {

	//fmt.Println(golearning.GetExample())
	//dir, path := golearningday2.SeparatePath("virat/kohli/boy")

	//fmt.Println(dir)
	//fmt.Println(path)

	fmt.Println(golearningday2.GetOsArgsInfo())

	s := "hello world"

	s1 := `hello world
	
	`

	fmt.Println(s)

	fmt.Println(s1)

	fmt.Println(utf8.RuneCountInString(s))

}

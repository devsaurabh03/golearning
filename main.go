package main

import (
	"fmt"

	golearningday5 "github.com/golearning/snow/golearning-day5"
)

func main() {

	//fmt.Println(golearning.GetExample())
	//dir, path := golearningday2.SeparatePath("virat/kohli/boy")

	//fmt.Println(dir)
	//fmt.Println(path)

	/*fmt.Println(golearningday2.GetOsArgsInfo())

	var temp []string = golearningday2.GetOsArgsInfo()

	fmt.Println(temp[1])

	s := "hello world"

	s1 := `hello world

	`

	fmt.Println(s)

	fmt.Println(s1)

	fmt.Println(utf8.RuneCountInString(s))

	golearningday3.PrintIOTAExamples()*/

	//fmt.Println(golearningday4.PasswordChecker())

	//fmt.Println(strconv.ParseBool(os.Args[1]))

	/*_, myerr := strconv.ParseInt(os.Args[1], 2, 64)

	if myerr != nil {

		fmt.Println("something error has occured")
	}

	fmt.Println("printing bool value", myvalue)
	fmt.Println("printing err value", myerror)*/

	fmt.Println(golearningday5.Greetfortheday())

}

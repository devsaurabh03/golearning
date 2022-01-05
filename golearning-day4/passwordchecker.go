package golearningday4

import golearningday2 "github.com/golearning/snow/golearning-day2"

func PasswordChecker() (value bool) {

	var args []string = golearningday2.GetOsArgsInfo()

	if args[1] == "user" && args[2] == "password" {
		value = true
	} else {
		value = false
	}

	return value
}

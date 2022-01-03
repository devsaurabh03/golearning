package golearningday2

import (
	"path"
)

func SeparatePath(input string) (dir string, file string) {

	return path.Split(input)
}

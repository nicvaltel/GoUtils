package filter

import (
	"GoUtils/maybe"
	"GoUtils/types"
	"strconv"
)

func TakeFirst[A any](predicate types.Predicate[A], xs []A) maybe.Maybe[A] {
	for _, x := range xs {
		if predicate(x) {
			return maybe.Pure[A](x)
		}
	}
	return maybe.Nothing[A]{}
}

func StringIsFloat(str string) bool {
	_, err := strconv.ParseFloat(str, 64)

	if err != nil {
		return false
	} else {
		return true
	}
}

func StringIsInt(str string) bool {
	_, err := strconv.ParseInt(str, 10, 64)

	if err != nil {
		return false
	} else {
		return true
	}
}

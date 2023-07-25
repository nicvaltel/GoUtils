package filter

import (
	"strconv"

	"github.com/nicvaltel/GoUtils/maybe"
	"github.com/nicvaltel/GoUtils/types"
)

func Filter[A any](predicate types.Predicate[A], xs []A) []A {
	ys := make([]A, len(xs))
	i := 0
	for _, x := range xs {
		if predicate(x) {
			ys[i] = x
			i++
		}
	}
	return ys[0:i]
}

// take n . filter
func TakeN[A any](n uint, predicate types.Predicate[A], xs []A) []A {
	k := 0
	var ys []A = make([]A, n)
	for _, x := range xs {
		if k >= int(n) {
			return ys
		}
		if predicate(x) {
			ys[k] = x
			k++
		}
	}
	return ys[0:k]
}

// safeHead . filter
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

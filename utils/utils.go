package utils

import (
	"fmt"
	"os"
)

type Pair[A, B any] struct {
	Fst A
	Snd B
}

func MkPair[A, B any](a A, b B) Pair[A, B] {
	return Pair[A, B]{Fst: a, Snd: b}
}

func ErrorExit(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func IIF[A any](b bool, ifTrue A, ifFalse A) A {
	if b {
		return ifTrue
	} else {
		return ifFalse
	}
}

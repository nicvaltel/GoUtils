package utils

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
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

func Carry[A, B, Z any](f func(_ A, _ B) Z, a A) (g func(_ B) Z) {
	return func(b B) Z { return f(a, b) }
}

func Carry3[A, B, C, Z any](f func(_ A, _ B, _ C) Z, a A) (g func(_ B, _ C) Z) {
	return func(b B, c C) Z { return f(a, b, c) }
}

func Carry4[A, B, C, D, Z any](f func(_ A, _ B, _ C, _ D) Z, a A) (g func(_ B, _ C, _ D) Z) {
	return func(b B, c C, d D) Z { return f(a, b, c, d) }
}

func Carry5[A, B, C, D, E, Z any](f func(_ A, _ B, _ C, _ D, _ E) Z, a A) (g func(_ B, _ C, _ D, _ E) Z) {
	return func(b B, c C, d D, e E) Z { return f(a, b, c, d, e) }
}

func AstPrinter(filename string) {
	fset := token.NewFileSet()

	node, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		fmt.Println("astPrinter error:", err)
		return
	}

	ast.Print(fset, node)
}

package maybe

import "errors"

type Maybe[A any] interface {
	IsJust() bool
	FromJust() (A, error)
	OrElse(other A) A
}

type Just[A any] struct {
	Val A
}

type Nothing[A any] struct {
}

func (_ Just[A]) IsJust() bool {
	return true
}

func (_ Nothing[A]) IsJust() bool {
	return false
}

func (j Just[A]) FromJust() (A, error) {
	return j.Val, nil
}

var NoElementError = errors.New("Trying to Get() from Nothing")

func (_ Nothing[A]) FromJust() (A, error) {
	return *new(A), NoElementError
}

func (j Just[A]) OrElse(_ A) A {
	return j.Val
}

func (_ Nothing[A]) OrElse(a A) A {
	return a
}

func Pure[A any](a A) Maybe[A] {
	return Just[A]{Val: a}
}

func NewNothing[A any]() Maybe[A] {
	return Nothing[A]{}
}

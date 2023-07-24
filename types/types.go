package types

type Func[A, B any] func(a A) B
type Predicate[A any] Func[A, bool]

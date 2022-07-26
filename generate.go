package contalgo

// Range make a []T filled with values in the `[first, last)` sequence
func Range[T Numeric](first, last T) []T {
	a := make([]T, 0, int(last-first))
	for v := first; v < last; v++ {
		a = append(a, v)
	}
	return a
}

// Generate fill each element of `a`` with `gen()``
func Generate[T Numeric](a []T, gen func() T) {
	for i := range a {
		a[i] = gen()
	}
}

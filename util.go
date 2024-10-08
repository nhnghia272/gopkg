package gopkg

func CopyArray[E comparable](src []E) []E {
	dst := make([]E, len(src))
	copy(dst, src)
	return dst
}

func CopyArray2D[E comparable](s [][]E) [][]E {
	c := make([][]E, len(s))
	for i, v := range s {
		c[i] = make([]E, len(v))
		copy(c[i], v)
	}
	return c
}

func UniqueFunc[E1, E2 comparable](s []E1, f func(E1) E2) []E1 {
	s2 := make([]E1, 0)
	m2 := make(map[E2]E1)
	for _, v := range s {
		k := f(v)
		if _, ok := m2[k]; !ok {
			m2[k] = v
			s2 = append(s2, v)
		}
	}
	return s2
}

func MapFunc[E1, E2 comparable](s []E1, f func(E1) E2) []E2 {
	s2 := make([]E2, len(s))
	for i, v := range s {
		s2[i] = f(v)
	}
	return s2
}

func FilterFunc[E comparable](s []E, f func(E) bool) []E {
	s2 := make([]E, 0)
	for _, v := range s {
		if f(v) {
			s2 = append(s2, v)
		}
	}
	return s2
}

func ReduceFunc[E1, E2 comparable](s []E1, a E2, f func(a E2, e E1, i int) E2) E2 {
	for i, v := range s {
		a = f(a, v, i)
	}
	return a
}

func SomeFunc[E comparable](s []E, f func(E) bool) bool {
	for _, v := range s {
		if f(v) {
			return true
		}
	}
	return false
}

func EveryFunc[E comparable](s []E, f func(E) bool) bool {
	for _, v := range s {
		if !f(v) {
			return false
		}
	}
	return true
}

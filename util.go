package gopkg

import (
	"encoding/json"
)

func Convert[E1, E2 any](src E1, des E2) error {
	bytes, err := json.Marshal(src)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, des)
}

func Transform[E1, E2 any](src E1) (E2, error) {
	var des E2
	return des, Convert(src, &des)
}

func LoopFunc[E any](s []E, f func(e E)) {
	for _, v := range s {
		f(v)
	}
}

func LoopParallelFunc[E any](s []E, f func(e E)) error {
	as := Async()
	for _, v := range s {
		as.Go(func() { f(v) })
	}
	return as.Wait()
}

func LoopWithIndexFunc[E any](s []E, f func(e E, i int)) {
	for i, v := range s {
		f(v, i)
	}
}

func LoopWithIndexParallelFunc[E any](s []E, f func(e E, i int)) error {
	as := Async()
	for i, v := range s {
		as.Go(func() { f(v, i) })
	}
	return as.Wait()
}

func UniqueFunc[E any, K comparable](s []E, f func(e E) K) []E {
	s2 := make([]E, 0)
	m2 := make(map[K]E)
	for _, v := range s {
		k := f(v)
		if _, ok := m2[k]; !ok {
			m2[k] = v
			s2 = append(s2, v)
		}
	}
	return s2
}

func SliceToMapFunc[E any, K comparable](s []E, f func(e E) K) map[K]E {
	m := make(map[K]E)
	for _, v := range s {
		k := f(v)
		m[k] = v
	}
	return m
}

func GroupFunc[E any, K comparable](s []E, f func(e E) K) map[K][]E {
	m := make(map[K][]E)
	for _, v := range s {
		k := f(v)
		m[k] = append(m[k], v)
	}
	return m
}

func MapFunc[E1, E2 any](s []E1, f func(e E1) E2) []E2 {
	s2 := make([]E2, len(s))
	for i, v := range s {
		s2[i] = f(v)
	}
	return s2
}

func MapParallelFunc[E1, E2 any](s []E1, f func(e E1) E2) ([]E2, error) {
	s2 := make([]E2, len(s))
	as := Async()
	for i, v := range s {
		as.Go(func() { s2[i] = f(v) })
	}
	if err := as.Wait(); err != nil {
		return nil, err
	}
	return s2, nil
}

func FilterFunc[E any](s []E, f func(e E) bool) []E {
	s2 := make([]E, 0)
	for _, v := range s {
		if f(v) {
			s2 = append(s2, v)
		}
	}
	return s2
}

func FilterParallelFunc[E any](s []E, f func(e E) bool) ([]E, error) {
	ok := make([]bool, len(s))
	as := Async()
	for i, v := range s {
		as.Go(func() { ok[i] = f(v) })
	}
	if err := as.Wait(); err != nil {
		return nil, err
	}
	s2 := make([]E, 0)
	for i := range ok {
		if ok[i] {
			s2 = append(s2, s[i])
		}
	}
	return s2, nil
}

func FindFunc[E any](s []E, f func(e E) bool) (E, bool) {
	for _, v := range s {
		if f(v) {
			return v, true
		}
	}
	return *new(E), false
}

func ReduceFunc[E1, E2 any](s []E1, t E2, f func(t E2, e E1) E2) E2 {
	for _, v := range s {
		t = f(t, v)
	}
	return t
}

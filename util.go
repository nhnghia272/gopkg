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

func LoopFunc[E any](s []E, f func(i int, e E)) {
	for i, v := range s {
		f(i, v)
	}
}

func LoopParallelFunc[E any](s []E, f func(i int, e E)) error {
	as := Async()
	for i, v := range s {
		as.Go(func() { f(i, v) })
	}
	return as.Wait()
}

func UniqueFunc[E1 any, E2 comparable](s []E1, f func(e E1) E2) []E1 {
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

func SliceToMapFunc[E1 any, E2 comparable](s []E1, f func(e E1) E2) map[E2]E1 {
	m := make(map[E2]E1)
	for _, v := range s {
		k := f(v)
		m[k] = v
	}
	return m
}

func GroupFunc[E1 any, E2 comparable](s []E1, f func(e E1) E2) map[E2][]E1 {
	m := make(map[E2][]E1)
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

func ReduceFunc[E1, E2 any](s []E1, a E2, f func(a E2, e E1) E2) E2 {
	for _, v := range s {
		a = f(a, v)
	}
	return a
}

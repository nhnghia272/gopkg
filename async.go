package gopkg

import (
	"fmt"
	"sync"
)

type async struct {
	wg  sync.WaitGroup
	one sync.Once
	err error
}

func Async() *async {
	return &async{}
}

func (s *async) Wait() error {
	s.wg.Wait()
	return s.err
}

func (s *async) Go(fn func()) *async {
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		defer func() {
			if r := recover(); r != nil {
				s.one.Do(func() {
					s.err = fmt.Errorf("%v", r)
					Debug(r)
				})
			}
		}()
		fn()
	}()
	return s
}

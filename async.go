package gopkg

import (
	"fmt"
	"os"
	"runtime/debug"
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
					os.Stderr.WriteString(fmt.Sprintf("panic: %v\n%s\n", r, debug.Stack()))
				})
			}
		}()
		fn()
	}()
	return s
}

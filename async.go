package gopkg

import (
	"fmt"
	"os"
	"runtime/debug"
	"sync"
)

type async struct {
	wg   sync.WaitGroup
	one  sync.Once
	err  error
	errs []error
}

func Async() *async {
	return &async{}
}

func (s *async) Wait() error {
	s.wg.Wait()
	return s.err
}

func (s *async) Waits() []error {
	s.wg.Wait()
	return s.errs
}

func (s *async) Go(fn func()) *async {
	s.wg.Add(1)

	i := len(s.errs)
	s.errs = append(s.errs, nil)

	go func() {
		defer s.wg.Done()
		defer func() {
			if r := recover(); r != nil {
				s.errs[i] = fmt.Errorf("%v", r)
				s.one.Do(func() {
					s.err = s.errs[i]
					os.Stderr.WriteString(fmt.Sprintf("panic: %v\n%s\n", r, debug.Stack()))
				})
			}
		}()
		fn()
	}()

	return s
}

# Golang Pkg

## Install Dependencies
```
go get github.com/nhnghia272/gopkg
```

## Example
```go
package main

import (
	"fmt"
	"time"

	"github.com/nhnghia272/gopkg"
)

func main() {
	// Initialize a new Async
	as := gopkg.Async()
	cache := gopkg.NewCacheShard[string](1)

	// Add goroutine
	as.Go(func() {
		password := gopkg.Random(20)
		fmt.Println("RandomPassword:", password)

		cache.Set("test", password, time.Second)
		fmt.Println(cache.Get("test"))
	})

	arr := []int{1, 2, 3, 4, 5, 2}
	fmt.Println("Init Slice:", arr)

	gopkg.LoopFunc(arr, func(e int) { fmt.Println("LoopFunc:", e) })

	gopkg.LoopParallelFunc(arr, func(e int) { fmt.Println("LoopParallelFunc:", e) })

	gopkg.LoopWithIndexFunc(arr, func(e int, i int) { fmt.Println("LoopWithIndexFunc:", e, i) })

	gopkg.LoopWithIndexParallelFunc(arr, func(e int, i int) { fmt.Println("LoopWithIndexParallelFunc:", e, i) })

	fmt.Println("UniqueFunc:", gopkg.UniqueFunc(arr, func(e int) int { return e }))

	fmt.Println("GroupFunc:", gopkg.GroupFunc(arr, func(e int) int { return e }))

	fmt.Println("SliceToMapFunc:", gopkg.SliceToMapFunc(arr, func(e int) int { return e }))

	fmt.Println("MapFunc:", gopkg.MapFunc(arr, func(e int) int { return e * 2 }))

	mp, err := gopkg.MapParallelFunc(arr, func(e int) int { return e * 3 })
	fmt.Println("MapParallelFunc:", mp, err)

	fmt.Println("FilterFunc:", gopkg.FilterFunc(arr, func(e int) bool { return e%2 == 0 }))

	fp, err := gopkg.FilterParallelFunc(arr, func(e int) bool { return e%2 == 1 })
	fmt.Println("FilterParallelFunc:", fp, err)

	item, found := gopkg.FindFunc(arr, func(e int) bool { return e == 3 })
	fmt.Println("FindFunc:", item, found)

	fmt.Println("ReduceFunc:", gopkg.ReduceFunc(arr, 0, func(a int, e int) int { return a + e }))

	// Wait goroutine finish
	if err := as.Wait(); err != nil {
		fmt.Println(err)
	}
}
```
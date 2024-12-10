# Golang Pkg

## Download from Github
From your project directory:
```
go get github.com/nhnghia272/gopkg
```

## Example
```go
package main

import (
	"log"

	"github.com/nhnghia272/gopkg"
)

func main() {
	// Initialize a new Async
	as := gopkg.Async()

	// Add goroutine
	as.Go(func() {
		password := gopkg.Random(20)
		log.Println("RandomPassword:", password)
	})

	arr := []int{1, 2, 3, 4, 5}

	log.Println(gopkg.UniqueFunc(arr, func(e int, i int) int {
		return e
	}))

	log.Println(gopkg.MapFunc(arr, func(e int, i int) int {
		return e * 2
	}))

	log.Println(gopkg.FilterFunc(arr, func(e int, i int) bool {
		return e%2 == 0
	}))

	log.Println(gopkg.FindFunc(arr, func(e int, i int) bool {
		return e == 3
	}))

	log.Println(gopkg.ReduceFunc(arr, 0, func(a int, e int, i int) int {
		return a + e
	}))

	log.Println(gopkg.SomeFunc(arr, func(e int, i int) bool {
		return e == 3
	}))

	log.Println(gopkg.EveryFunc(arr, func(e int, i int) bool {
		return e%2 == 0
	}))

	// Wait goroutine finish
	if err := as.Wait(); err != nil {
		log.Println(err)
	}
}
```
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
	v, ok := gopkg.FindFunc(arr, func(v int) bool {
		return v == 6
	})
	log.Println(v, ok)

	// Wait goroutine finish
	if err := as.Wait(); err != nil {
		log.Println(err)
	}
}

```
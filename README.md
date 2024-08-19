# Golang Pkg

## Download from Github
From your project directory:
```
go get github.com/nhnghia272/gopkg
```

Enviroment
```
LOG_LEVEL=info
```

## Example
```go
package main

import "github.com/nhnghia272/gopkg"

func main() {
	// Initialize a new Async
	as := gopkg.Async()

	// Add goroutine
	as.Go(func() {
		password := gopkg.Random(20)
		gopkg.Info("RandomPassword:", password)
	})

	// Wait goroutine finish
	if err := as.Wait(); err != nil {
		gopkg.Error(err)
	}
}
```

## Logrus
[github.com/sirupsen/logrus](github.com/sirupsen/logrus)
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
	"github.com/nhnghia272/gopkg"
)

type Person struct {
	Name string `json:"name" validate:"required"`
	Age  uint64 `json:"age" validate:"omitempty"`
}

func main() {
	// Initialize a new Async
	as := gopkg.NewAsync()

	// Add goroutine
	as.Go(func() {
		password := gopkg.Random(20)
		gopkg.Info("RandomPassword:", password)

		hashpassword := gopkg.HashPassword(password, 14)
		gopkg.Info("HashPassword:", hashpassword)

		person := Person{Name: "", Age: 10}
		gopkg.Info("Validate:", gopkg.NewValidator(nil).Validate(person))
	})

	// Wait goroutine finish
	if err := as.Wait(); err != nil {
		gopkg.Error(err)
	}
}
```

## Logrus
[github.com/sirupsen/logrus](github.com/sirupsen/logrus)

## Mongo
[go.mongodb.org/mongo-driver](go.mongodb.org/mongo-driver)

## Validator
[https://github.com/go-playground/validator](https://github.com/go-playground/validator)

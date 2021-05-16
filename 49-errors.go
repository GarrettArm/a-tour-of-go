package main

import (
	"fmt"
	"time"
)

/*
  error interface is like fmt.Stringer interface,
  with an Error() method instead of String()

  type error interface {
  	Error() string
  }

  i, err := someMethod()
  if err != nil {
  	do thing
  	return
  }
  otherwise, do other thing
*/

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

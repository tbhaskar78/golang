/****************************************

* File Name : error.go

* Creation Date : 27-08-2020

* Last Modified : Thursday 27 August 2020 10:55:08 PM

* Created By :  Bhaskar Tallamraju

*****************************************/
package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
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


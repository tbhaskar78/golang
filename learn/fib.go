/****************************************

* File Name : fib.go

* Creation Date : 27-08-2020

* Last Modified : Thursday 27 August 2020 11:07:01 PM

* Created By :  Bhaskar Tallamraju

*****************************************/
package main

import (
    "fmt"
)

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
    prev := 0
	cur := 1
	num := 0
    fmt.Println(prev)
	fmt.Println(cur)
	return func () int {
	    num = cur + prev
		prev = cur
		cur = num
		return cur
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

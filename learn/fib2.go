/****************************************

* File Name : fib2.go

* Creation Date : 27-08-2020

* Last Modified : Thursday 27 August 2020 11:03:42 PM

* Created By :  Bhaskar Tallamraju

*****************************************/
package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
    /* buffered channel of 5, so we should close one end of channel to let the reader know
     * channel need not be closed otherwise. close(channel) only to communicate to reader
     */
	c := make(chan int, 5)
	go fibonacci(cap(c), c)  /* send capacity of c as the first arg */
	for i := range c {
		fmt.Println(i)
	}
}


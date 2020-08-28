/****************************************

* File Name : channel.go

* Creation Date : 27-08-2020

* Last Modified : Thursday 27 August 2020 11:00:33 PM

* Created By :  Bhaskar Tallamraju

*****************************************/
package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
    /* push sum in the channel */
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

    /* create a channel, they are like pipes, data flow is in the direction of the arrow */
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)

    chn := make(chan int, 2)  // buffered channel 
    chn <- 1
    chn <- 2
    //chn <- 3 this will cause an error, as only 2 ints can be buffered
    fmt.Println(<-chn)
    fmt.Println(<-chn)
    //fmt.Println(<-chn) this will cause an error too, as no more data is available now at this point
}


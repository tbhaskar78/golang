/****************************************

* File Name : qarray.go

* Creation Date : 03-09-2020

* Last Modified : Thursday 03 September 2020 10:20:31 AM

* Created By :  Bhaskar Tallamraju

* Description: queue implementation using arrays
*
* Queue: its a data structure which works based on principle of first in first out (FIFO).
* In computing world, queue data structure can be applied in many applications where order is important
* such as holding messages. enqueue and dequeue are the operations that are provided for insertion 
* of an element into the stack and the removal of an element from the queue.
*
*   two indexes, readq follows writeq. After reading, check if readq equals writeq and then 
*   re-initialize the indexes. 
*****************************************/
package main

import (
    "fmt"
)

// global variables
// max elements that the queue can hold
const MAX_ELEMENTS int = 10
// the array ds to hold the queue
var queue [MAX_ELEMENTS]int
// readq and writeq to manage the indexes for queueing and dequeueing
var readq int = 0
var writeq int = -1

// enqueue function to queue elements in FIFO order
func enqueue(data int) {
   if writeq >= (MAX_ELEMENTS-1) {
      fmt.Printf ("Queue full, cannot insert %v\n", data)
   } else {
      writeq++
      queue[writeq] = data
   }
}

// dequeue function to queue elements in FIFO order
func dequeue() {
   if writeq == -1 {
      fmt.Printf ("Queue empty\n")
   } else {
      fmt.Printf ("data dequeued %v\n", queue[readq])
      queue[readq] = 0  /* erase the value */
      if readq == writeq {
          readq = 0
          writeq = -1
      } else {
          readq++   // only post increment in golang
      }
   }
}

func main() {

    // enqueue, check queue full after MAX_ELEMENTS
    for i := 0; i < MAX_ELEMENTS+2; i++ {
        enqueue(i+1)
    }
    fmt.Printf("After enqueing all elements, queue : %v\n", queue)

    // dequeue, check queue empty after dequeuing all elements
    for i := 0; i < MAX_ELEMENTS+4; i++ {
        dequeue()
    }
    fmt.Printf("After dequeing all elements, queue : %v\n", queue)

    return
}

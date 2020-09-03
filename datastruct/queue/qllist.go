/****************************************

* File Name : qllist.go

* Creation Date : 03-09-2020

* Last Modified : Thursday 03 September 2020 10:16:26 AM

* Created By :  Bhaskar Tallamraju

* Description: queue implementation using linkedlist
*
* Queue: its a data structure which works based on principle of first in first out (FIFO).
* In computing world, queue data structure can be applied in many applications where order is important
* such as holding messages. enqueue and dequeue are the operations that are provided for insertion 
* of an element into the stack and the removal of an element from the queue.
*
*   head -> |6|next| -> |5|next| -> |7|next| -> |newData|next|->NULL
*                                               q ->
*
*   head ptr points to the first element inserted and q points to the latest element added. 
****************************************/
package main

import (
    "fmt"
)

// global variables
// struct to hold each element in the queue
type queue struct {
   data int
   next *queue
}

// max elements that the queue can hold
const MAX_ELEMENTS int = 10
// the head pointer to the queue
var head *queue = nil
// the counter to keep track of elements
var top int = -1;

// enqueue function to queue elements in FIFO order
func enqueue(q **queue, data int) {
   if top >= MAX_ELEMENTS-1 {
      fmt.Printf ("Queue full, cannot insert %v\n", data);
   } else {
      top++   //always post increment in golang
      if *q == nil {
         *q = &queue{data: data, next: nil, }
         head = *q
      } else {
         (*q).next = &queue{data: data, next: nil, }
         *q = (*q).next
      }
   }
}

// dequeue function to queue elements in FIFO order
func dequeue(q **queue) {
   if head == nil {
      fmt.Printf ("Queue empty\n");
   } else {
      fmt.Printf ("data dequeued %v\n", head.data)
      head = head.next
      top--   //always post decrement in golang
   }
}

func main() {
   var q *queue = nil

   // enqueue, check queue full after MAX_ELEMENTS
   for i := 0; i < MAX_ELEMENTS+4; i++ {
      enqueue(&q, i+1);
   }

   // dequeue, check queue empty after dequeuing all elements
   for i := 0; i < MAX_ELEMENTS+2; i++ {
      dequeue(&q);
   }
   return
}

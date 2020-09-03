/****************************************

* File Name : stkarray.go

* Creation Date : 03-09-2020

* Last Modified : Thursday 03 September 2020 08:56:48 AM

* Created By :  Bhaskar Tallamraju

* Description: stack implementation using arrays
*
* Stack: its a data structure which works based on principle of last in first out (LIFO).
* In computing world, stack data structure can be applied in many applications such as 
* parsing syntax of expressions and solving search problem. Push and pop are the operations 
* that are provided for insertion of an element into the stack and the removal of an element 
* from the stack.
*****************************************/
package main

import (
    "fmt"
)

const MAX_ELEMENTS int = 10
var stack [MAX_ELEMENTS]int
var top int = -1

func push(data int) {
   if top >= (MAX_ELEMENTS-1) {
       fmt.Printf ("Stack overflow, cannot insert %v\n", data)
   } else {
      top++
      stack[top] = data
   }
}

func pop() {
   if top == -1 {
      fmt.Printf ("Stack underflow, no more elements in the stack\n")
   } else {
      fmt.Printf ("data popped %d\n", stack[top])
      stack[top] = 0  /* erase the value */
      top--
   }
}

func main() {
   for i := 0; i < MAX_ELEMENTS+2; i++ {
      push(i+1)
   }
   fmt.Printf("After pushing elements 1-10, stack: %v\n", stack)

   for i := 0; i < MAX_ELEMENTS+4; i++ {
      pop()
   }
   fmt.Printf("After popping elements 1-10, stack: %v\n", stack)

   return
}

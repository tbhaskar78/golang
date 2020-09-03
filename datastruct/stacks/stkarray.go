/****************************************

* File Name : stkarray.go

* Creation Date : 03-09-2020

* Last Modified : Thursday 03 September 2020 10:21:12 AM

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

//global variables

// indicate the max stack size
const MAX_ELEMENTS int = 10
// the array ds to hold the stack
var stack [MAX_ELEMENTS]int
// the index counter
var top int = -1

/* push the numbers to the list, top of stack pointing to the last element */
func push(data int) {
   if top >= (MAX_ELEMENTS-1) {
       fmt.Printf ("Stack overflow, cannot insert %v\n", data)
   } else {
      top++
      stack[top] = data
   }
}

/* pop the list, starting from top of stack */
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
    // push elements from 1 to index, test if stack overflows after reaching MAX_STACK
    for i := 0; i < MAX_ELEMENTS+2; i++ {
        push(i+1)
    }
    fmt.Printf("After pushing elements 1-10, stack: %v\n", stack)

    // pop elements in reverse, test if stack underflows after popping all elements
    for i := 0; i < MAX_ELEMENTS+4; i++ {
        pop()
    }
    fmt.Printf("After popping elements 1-10, stack: %v\n", stack)

    return
}

/****************************************

* File Name : stkllist.go

* Creation Date : 02-09-2020

* Last Modified : Thursday 03 September 2020 10:19:46 AM

* Created By :  Bhaskar Tallamraju

* Description: stack implementation using linkedlist
*
* Stack: its a data structure which works based on principle of last in first out (LIFO).
* In computing world, stack data structure can be applied in many applications such as 
* parsing syntax of expressions and solving search problem. Push and pop are the operations 
* that are provided for insertion of an element into the stack and the removal of an element 
* from the stack.
*  Example : 
*   topstack -> |newElement|next| -> |5|next| -> |7|next| -> |firstelement|next|->NULL
*
*   topstack ptr points to the new element pushed and gets popped from there too. 
*****************************************/
package main

import (
    "fmt"
)

// global variables
// the node for each element in the stack
type stack struct {
    data int
    next *stack
}

// global variable keeping track of number of elements in the stack
var top int = -1

// global variable to indicate the max stack size
var MAX_STACK int = 10

/* push the numbers to the list, top of stack pointing to the last element */
func push(topstack **stack, data int) {
    if top >= (MAX_STACK-1) {
        fmt.Printf ("Stack overflow\n")
    } else {
        top++
        fmt.Println(data)
        *topstack = &stack{data: data, next: *topstack,}
    }
}

/* pop the list, starting from top of stack */
func pop(topstack **stack) {
    if top == -1 {
        fmt.Printf ("Stack underflow\n")
    } else {
        fmt.Printf ("data popped %d\n", (*topstack).data)
        *topstack = (*topstack).next
        top--
    }
}

func main() {
    var topstack *stack = nil

    // push elements from 1 to index, test if stack overflows after reaching MAX_STACK
    for i := 0; i<MAX_STACK+4; i++ {
        push(&topstack, i+1)
    }

    // pop elements in reverse, test if stack underflows after popping all elements
    for i := 0; i<MAX_STACK+2; i++ {
        pop(&topstack)
    }
}

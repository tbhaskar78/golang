/****************************************

* File Name : pointer.go

* Creation Date : 17-08-2020

* Last Modified : Monday 17 August 2020 05:03:43 PM

* Created By :  Bhaskar Tallamraju

*****************************************/
package main

import (
    "fmt"
)

func main() {
   i := 20
   var p *int = &i
   fmt.Printf("Pointer is %d\n", *p)
   *p = 21
   fmt.Printf("Pointer is %d\n", i)
   var q *int = nil
   /* the following line will cause a segmentation */
   /*fmt.Printf("Pointer %d\n", *q) */
}


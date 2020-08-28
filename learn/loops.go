/****************************************

* File Name : loops.go

* Creation Date : 17-08-2020

* Last Modified : Monday 17 August 2020 05:02:58 PM

* Created By :  Bhaskar Tallamraju

*****************************************/
package main

import (
    "fmt"
    "runtime"
    "time"
)


func main() {

    for i := 0; i < 10; i++ {
        fmt.Println(i)
    }

    /* there is no while loop in GO, for is while in golang */

    sum := -1
    for sum < 10 {
        sum += 1
    }
    fmt.Printf("Sum is %v\n", sum)

    /* infinite loop */
    for {
        sum += 1;
        if sum == 99 {
            break
        }
    }
    fmt.Printf("Sum is %v\n", sum)

    /* if can start with a short init like in for */
    if v := 99; v < 100 {
        fmt.Printf("I am checking v is less than 100 \n")
    } else {
        fmt.Printf("v is more than 100 \n")
    }


    /* switch case */
    switch os := runtime.GOOS; os {
        case "linux":
            fmt.Printf("I am linux\n")
        default:
            fmt.Printf("I am any other OS\n")
    }


    /* switch case with evaluation */
    today := time.Now().Weekday()
    switch time.Sunday {
        case today + 0:
            fmt.Printf("Today is Sunday \n")
        case today + 1:
            fmt.Printf("1 Days after sunday \n")
        case today + 2:
            fmt.Printf("2 Days after sunday \n")
        default:
            fmt.Printf("too far out !\n")
    }

    /* switch case with no condition */
    hour := time.Now()
    switch  {
        case hour.Hour() < 12:
            fmt.Printf("Good morn!\n")
        case hour.Hour() < 17:
            fmt.Printf("Good afternoon!\n")
        default:
            fmt.Printf("Good night!\n")
    }

    /* range is a for loop over slice or map; it returns 2 values, index into the slice
       and value at the index */
    pow := []int{1, 2, 4, 8, 16, 32, 64 , 128} /* create a slice  */
    fmt.Printf("range over pow \n")
    for i,v := range pow {  /* now range over it, returns i and v */
        fmt.Printf("2^%d = %d \n", i, v)
    }
    for i,_ := range pow {  /* now range over it, only consider i not v */
        fmt.Printf("2^%d = ? \n", i)
    }
    for _,v := range pow {  /* now range over it, only consider i not v */
        fmt.Printf("2^? = %d \n", v)
    }

    /* defer function : defers till nearby functions complete */
    defer fmt.Printf("I will go last\n")
    fmt.Printf("I will go first\n")
    fmt.Printf("I will go second\n")


    /* defer is pushed onto stack and is popped in LIFO manner */
    for j := 0; j < 10; j++ {
        defer fmt.Println(j)
    }

}

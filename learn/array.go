/****************************************

* File Name : array.go

* Creation Date : 15-08-2020

* Last Modified : Saturday 15 August 2020 11:00:17 PM

* Created By :  Bhaskar Tallamraju

*****************************************/
package main

import (
    "fmt"
)

func main() {

    var arr [3]int
    arr[0] = 10
    arr[1] = 100
    arr[2] = 1000
    defer fmt.Println(arr)
    fmt.Println(len(arr))

    arr1 := [2]string{"Hello", "World"}
    fmt.Println(arr1)

}

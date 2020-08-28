/****************************************

* File Name : swapString.go

* Creation Date : 17-08-2020

* Last Modified : Monday 17 August 2020 05:04:32 PM

* Created By :  Bhaskar Tallamraju

*****************************************/
package main

import (
    "fmt"
    )

func swap(x, y string) (string, string) {
    return y, x
}

func main() {
    a, b := swap("Hello", "World");
    fmt.Println(a, b)
}

/****************************************

* File Name : struct.go

* Creation Date : 15-08-2020

* Last Modified : Saturday 15 August 2020 10:18:03 PM

* Created By :  Bhaskar Tallamraju

*****************************************/
package main

import (
    "fmt"
)

type st struct {
    x int
    y float64
}

func main() {
    str := st{}
    var stt st
    str.x = 10
    str.y = 10.9876
    fmt.Println(str)
    fmt.Println(stt)

}

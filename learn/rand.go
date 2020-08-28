/****************************************

* File Name : rand.go

* Creation Date : 17-08-2020

* Last Modified : Monday 17 August 2020 05:04:05 PM

* Created By :  Bhaskar Tallamraju

*****************************************/
package main

import (
    "fmt"
    "math"
    "math/rand"
    "time"
    )

func main() {
    seed := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(seed)
    fmt.Println(r1.Intn(100))
    fmt.Println(math.Pi)
}

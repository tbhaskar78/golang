/****************************************

* File Name : map.go

* Creation Date : 17-08-2020

* Last Modified : Monday 17 August 2020 04:58:24 PM

* Created By :  Bhaskar Tallamraju

*****************************************/
package main

import (
    "fmt"
)
type Coord struct {
    x, y float64
}
func main() {

    m := make(map[int]Coord)
    m[0] = Coord{2, 3,}
    m[1] = Coord{3.141, 2.63,}

    fmt.Println(m[0])
    mstr := make(map[string]Coord)
    mstr["India"] = Coord{9.999, 8.63,}
    fmt.Println(mstr["India"])

    m2 := map[string]Coord{
            "Bell Labs": {40.68433, -74.39967},
            "Google":    {37.42202, -122.08408},
    }
    fmt.Println("The value at key Google is :", m2["Google"])
    delete(m2, "Google")
    fmt.Println("The value at key Google after delete is :", m2["Google"])

    v, ok := m2["Google"]
    fmt.Println("The value at key Google to check if key is present :", v, "Present?", ok)

}

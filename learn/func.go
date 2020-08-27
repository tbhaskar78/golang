/****************************************

* File Name : func.go

* Creation Date : 27-08-2020

* Last Modified : Thursday 27 August 2020 10:11:29 PM

* Created By :  Bhaskar Tallamraju

*****************************************/
package main

import (
    "fmt"
    "math"
)

type Vertex struct {
        X, Y float64
}

/* passing function as a parameter to compute */
func compute(fn func(x, y int) int) int {
    return fn( 3, 4)
}

/* creating a closure, each closure is bound to its own variables that it accesses */
func adder() func(int) int {
    sum := 0
    return func (x int) int {
        sum += x
        return sum
    }
}

/* writing a method for a type */
func (v Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/* use pointer to scale */
func (v *Vertex) ptrScale(f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}

/* use normal variable to scale: this will work on local variable and does not change the values in Vertex */
func (v Vertex) Scale(f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}

func main() {

    mult := func(x, y int) int {
        return (x*y)
    }

    fmt.Printf("\n============ Printing values by passing function as param ============== \n")
    fmt.Println(compute(mult))
    fmt.Println(mult(4, 5))


    pos := adder()
    neg := adder()    /* these 2 will have their own sum variable references */
    fmt.Printf("\n============ Printing 2 closures ============== \n")
    for i := 0; i < 10; i++ {
        fmt.Println(pos(i), neg(-2*i))
    }


    /* print with pointer for scaling and one without, notice the difference */
    fmt.Printf("\n============ Printing methods for types ============== \n")
    v := Vertex{3, 4,}
    v.Scale(10)
    fmt.Println(v.Abs())
    v.ptrScale(10)
    fmt.Println(v.Abs())


}

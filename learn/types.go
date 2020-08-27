/****************************************

* File Name : types.go

* Creation Date : 17-08-2020

* Last Modified : Monday 17 August 2020 05:05:07 PM

* Created By :  Bhaskar Tallamraju

*****************************************/
package main

import (
	"fmt"
	"math/cmplx"
    "math"
)

var (
	ToBe bool       = false
	MaxInt uint64     = 1<<64 - 1
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
    /* print global vars and their type */
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

    /* print local variables and their type */
    var a, b, c = 10, true, "no"
    f64 := 10.99999
    i32 := 1000
	fmt.Printf("Type: a %T \n", a);
	fmt.Printf("Type: b %T \n", b);
	fmt.Printf("Type: c %T \n", c);
	fmt.Printf("Type: f64 %T \n", f64);

    /* convert between types */
    a = int(math.Ceil(f64))  /* float to int, use ceil  */
    fmt.Printf("int ceil val of f64 %v\n",  a)
    a = int(math.Floor(f64))  /* float to int, use floor  */
    fmt.Printf("int floor val of f64 %v\n",  a)
    f64 = float64(i32)  /* int to float */
    fmt.Printf("float64 val of i32 %v\n",  f64)

    /* constants cannot be declared using  := . It can be string, bool, numerical etc */
    const World string = "Hello World"
    const integer  = 100
    const f64_2  = 100.8767
    const bTruth = false

}

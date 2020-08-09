package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
    var a, b, c = 10, true, "no"
    d := 10.99999
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	fmt.Printf("Type: a %T \n", a);
	fmt.Printf("Type: b %T \n", b);
	fmt.Printf("Type: c %T \n", c);
	fmt.Printf("Type: d %T \n", d);
    fmt.Println(a, b, c, d)
}

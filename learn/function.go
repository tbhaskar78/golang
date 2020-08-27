/****************************************

* File Name : function.go

* Creation Date : 17-08-2020

* Last Modified : Monday 17 August 2020 05:01:40 PM

* Created By :  Bhaskar Tallamraju

*****************************************/
package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}


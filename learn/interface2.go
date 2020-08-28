/****************************************

* File Name : interface2.go

* Creation Date : 27-08-2020

* Last Modified : Thursday 27 August 2020 10:40:43 PM

* Created By :  Bhaskar Tallamraju

*****************************************/
package main

import (
    "fmt"
    "math"
)

/*  define an interface I, any interface written should define M() */
type I interface {
    M()
}

/*  Type T which is a struct */
type T struct {
    S string
}

/* type T defines M() */
func (t *T) M() {
    if (t == nil) {
        fmt.Printf("<nil>\n")
        return
    }
    fmt.Println(t.S)
}

/* type F is a float64 */
type F float64

/* type F defines M() */
func (f F) M() {
    fmt.Println(f)
}

func main() {
    var i I

    /* struct T is nil here , not initialized, handle it in the interface implemented */
    var tt *T
    i = tt
    describe(i)
    i.M()

    /* struct needs to be a pointer, whch defines M(). T is initialized to "Hello" */
    i = &T{"Hello"}
    describe(i)
    i.M()


    /* float64 directly used in the interface */
    i = F(math.Pi)
    describe(i)
    i.M()

    /* empty interfaces */
    var ii interface{}
    describe2(ii)
    ii = 42
    describe2(ii)
    ii = "hello"

    /* switch case based on type interface */
    do(21)
    do("hello")
    do(true)
}

func describe(i I) {
    fmt.Printf("(%v, %T)\n", i, i)
}
func describe2(i interface{}) {
    fmt.Printf("(%v, %T)\n", i, i)
}

func do(i interface{}) {
    switch v := i.(type) {
        case int:
            fmt.Printf("Twice %v is %v\n", v, v*2)
        case string:
            fmt.Printf("%q is %v bytes long\n", v, len(v))
        case bool:
            fmt.Printf("%v is true ?\n", v)
        default:
            fmt.Printf("I don't know about type %T!\n", v)
    }
}


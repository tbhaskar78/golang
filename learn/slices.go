/****************************************

* File Name : slices.go

* Creation Date : 15-08-2020

* Last Modified : Monday 17 August 2020 04:00:39 PM

* Created By :  Bhaskar Tallamraju

*****************************************/
package main

import (
    "fmt"
    "strings"
)

func main() {

    /* define an array */
    arr := [4]string{ "John", "George", "Adam", "Robin" }
    fmt.Printf("1. print original array \n\t")
    fmt.Printf("arr: %v\n", arr)

    /* now find a slice of array, it is like reference to a portion of an array 
       slice required a lower and upper bound of an array */
    a := arr[0:2]
    b := arr[1:3]
    fmt.Printf("2. checking upper and lower bounds a[0:2] and b[1:3] \n\t")

    fmt.Printf("a: %v \n\tb: %v\n", a, b)

    /* skipping the lower and upper bound defaults it to 0 and length for upper bound */
    a1 := arr[:]
    a2 := arr[0:]
    a3 := arr[0:len(arr)]
    a4 := arr[:len(arr)]
    fmt.Printf("3. All of these are same (without bounds mentioned) arr[:], arr[:10], arr[0:], arr[0:10] \n\t")
    fmt.Printf("a1: %v \n\ta2: %v\n\ta3: %v\n\ta4: %v\n", a1, a2, a3, a4)

    /* change one of the indices in the slice, should reflect in all slices/references and the original array */
    a[1] = "Bill"
    fmt.Printf("4. replacing slice index a[1] with Bill, changes the original array also and all who referece it \n\t")
    fmt.Printf("arr: %v \n\tb: %v\n", arr, b)

    /* slice literal, without the length*/
    slice := []bool{true, false, true, true, false}
    fmt.Printf("5. slice literal, without mentioning array length slice := []bool{true, false, etc...} \n\t")
    fmt.Println(slice)

    /* struct init with slices */
    slice2 := []struct{
                  x, y int
              }{
                  {2, 3},
                  {3, 4},
                  {4, 5},
                  {6, 7},
              }
    fmt.Printf("6. slice literal, without mentioning array length slice := []struct{x, y int} { {1,2}, {3,4},} \n\t")
    fmt.Println(slice2)

    /* length and capacity of slice 
       len is the len of the slice
       capacity is the length of the underlying layer starting from index in slice */
    fmt.Printf("7. checking len and cap of a[0:2] and b[1:3] \n\t")
    fmt.Printf("len of b %v and len of a %v\n\t", len(b), len(a))
    fmt.Printf("cap of b %v and cap of a %v\n", cap(b), cap(a))

    /* empty slice is nil not NULL/null */
    var try []int
    if try == nil {
        fmt.Printf("\ttry is nil cap %v len %v\n", len(try), cap(try))
    }
    /* extending or receding a  slice */
    a = a[:0]
    fmt.Printf("\tlen of a %v after receding it to a[:0]\n\t", len(a))
    a = a[:4]
    fmt.Printf("len of a %v after extending it to a[:4] elements are %v\n\t", len(a), a)
    a = a[2:]
    fmt.Printf("len of a %v after dropping first 2  a[2:] elements are  %v\n", len(a), a)


    /* create a slice with make */
    am := make([]int, 0, 5)
    fmt.Printf("len of am %v  cap of am %v\n", len(am), cap(am))
    am = am[:5]
    fmt.Printf("len of am %v  cap of am %v\n", len(am), cap(am))
    bm := am[2:5]
    fmt.Printf("len of bm %v  cap of bm %v\n", len(bm), cap(bm))

    /* slice of slices of strings and create an identity matrix of 3x3 */
    ident := [][]string{
            []string{"0", "0", "0"},
            []string{"0", "0", "0"},
            []string{"0", "0", "0"},
          }
    ident[0][0] = "1"
    ident[1][1] = "1"
    ident[2][2] = "1"
    for i := 0; i < len(ident); i++ {
        fmt.Printf("%s\n", strings.Join(ident[i], " "))
    }

    /* append elemets to a slice */
    var snil []int
    fmt.Printf("cap %v len %v \n", cap(snil), len(snil))
    snil = append(snil, 1, 2, 3, 4)
    fmt.Printf("cap %v len %v and snil is %v\n", cap(snil), len(snil), snil)

}


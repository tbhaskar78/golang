/****************************************

* File Name : nakedReturn.go

* Creation Date : 17-08-2020

* Last Modified : Monday 17 August 2020 05:03:22 PM

* Created By :  Bhaskar Tallamraju

*****************************************/
package main

import "fmt"

func split(sum int) (x, y int) {
    x = (sum * 4)/9
    y = sum - x
    return     /* naked return, returns the named return values */
}

func main(){
    fmt.Println(split(27))
}

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

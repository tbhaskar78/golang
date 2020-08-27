/****************************************

 * File Name : pic.go

 * Creation Date : 17-08-2020

 * Last Modified : Monday 17 August 2020 04:44:38 PM

 * Created By :  Bhaskar Tallamraju

 *****************************************/
package main

import (
        "golang.org/x/tour/pic"
       )



func Pic(dx, dy int) [][]uint8 {
    // Allocate two-dimensioanl array.
   a := make([][]uint8, dy)
   for i := 0; i < dy; i++ {
       a[i] = make([]uint8, dx)
   }

   // Do something.
   for i := 0; i < dy; i++ {
      for j := 0; j < dx; j++ {
          a[i][j] =  uint8(i^j)
      }
   }
   return a
}

func main() {
    pic.Show(Pic)
}

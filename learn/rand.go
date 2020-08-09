package main

import (
    "fmt"
    "math"
    "math/rand"
    "time"
    )

func main() {
    seed := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(seed)
    fmt.Println(r1.Intn(100))
    fmt.Println(math.Pi)
}

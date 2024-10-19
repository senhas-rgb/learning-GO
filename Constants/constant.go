package main

import (  //importing multiple libs
  "math"
  "fmt"
)

const s string = "this is a constant"

func main () {
  fmt.Println(s)

  const n = 500000000

  const d = 3e20 / n
  fmt.Println(d)

  fmt.Println(int64(d))

  fmt.Println(math.Sin(n))
}

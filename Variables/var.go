package main

import "fmt"

var test string = "this is a string outside main func()"

func main () {
  var a = "Init"
  fmt.Println(a)

  var b, c int = 1, 2
  fmt.Println(b, c)

  var d = true
  fmt.Println(d)

  var e int
  fmt.Println(e)

  f := "apple" // only used within functions. declares and initializes
  fmt.Println(f)

  fmt.Println(test)
}

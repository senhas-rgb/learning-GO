package main

import "fmt"

func main() {
  if 7%2 == 0{
    fmt.Println("7 is even")
  } else {
    fmt.Println("7 is odd")
  }

  if num := 9; num < 0 {
    fmt.Println(num, "is negative\n")
  } else if num < 10 {
    fmt.Println(num, "has 1 digit\n")
  } else {
    fmt.Println(num, "has multiple digits\n")
  }
}

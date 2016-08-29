package main

import (
  "fmt"
)

func main() {
  const pi = 3.14
  fmt.Printf("Value of pi is %f\n", pi)
  fmt.Printf("Value of pi is %1.1f\n", pi)

  fmt.Println("Value of pi is", pi)

  const (
    a = 5
    b = 10
  )

  fmt.Println("a =", a, "b =", b)
}

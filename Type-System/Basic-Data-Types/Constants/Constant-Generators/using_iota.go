package main

import (
  "fmt"
)

func main() {
  const (
    p = 2
    q
    r
    s = "aString"
    t
  )

  fmt.Println("p =", p, "q =", q, "r =", r, "s =", s, "t =", t)

  type Month int

  const (
    Jan Month = iota
    Feb
    Mar
    Apr
    May
    Jun
    Jul
    Aug
    Sep
    Oct
    Nov
    Dec
  )

  fmt.Println("Jan =", Jan, "Feb =", Feb, "Nov =", Nov, "Dec =", Dec)
}

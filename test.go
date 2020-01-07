package main

import (
  "fmt"
)

type Square struct {
  name string
}

func main() {

  pages := 10
  per := 3

  fmt.Println(pages / per)
}

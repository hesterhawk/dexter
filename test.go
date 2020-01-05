package main

import (
  "fmt"
)

type Shape interface {
  area() float64
}

type Square struct {
  x1, y1, x2, y2 float64
}

func (s *Square) area() float64 {
  l := s.x1 + s.y1 + s.x2 + s.y2
  return l * l
}

func main() {

  s := Square{0, 0, 5, 5}
  fmt.Println(s.area())
}

package main

import (
  "fmt"
  "sort"
)

type Square struct {
  name string
}

func main() {

  var s []Square

  s = append(s, Square{"aaaa"})
  s = append(s, Square{"cccc"})
  s = append(s, Square{"ffff"})
  s = append(s, Square{"bbbb"})
  s = append(s, Square{"mmmm"})

  sort.Slice(s[:], func(i, j int) bool {
    return s[i].name < s[j].name
  })

  fmt.Println(s)
}

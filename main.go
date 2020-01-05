package main

import (
  "fmt"
  //"log"
  "generator"
)

func main() {

  // temp
  var langs = []string {"pl", "en"}

  var drafts = generator.GetSortedDrafts(langs)

  fmt.Println(drafts)
  fmt.Println("done.")
}

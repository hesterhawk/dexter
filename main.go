package main

import (
  "fmt"
  //"log"
  "helper"
  "strings"
  "generator"  
)

func main() {

  // temp
  var langs = []string {"pl", "en"}

  var drafts = generator.GetAllSortedDrafts(langs)

  //generator.GeneratePubs(drafts)

  helper.ReadFile(drafts[0])

  fmt.Printf("%q\n", strings.Split("fdsas #title fdfdsafdsafdsa", "#title "))

  view.Make("{{.Count}} items are made of {{.Material}}")

  fmt.Println("done.")
}

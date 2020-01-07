package main

import (
  "fmt"
  //"log"
  "generator"
)


var CC = "DSADSADSA"

func main() {

  // ---- config
  var langs = []string {"pl", "en"}

  var perPage = 3

  // ----

  var drafts = generator.GetAllSortedDrafts(langs)

  for _, d := range drafts {
    fmt.Println(d)
  }

  if len(drafts) > 0 {

    generator.GenerateDraftsLists(drafts, perPage)
  }
}

package main

import (
  "generator"  
  "bootstrap"
)

func main() {

  // ---- config
  var langs = []string {"pl", "en"}

  var perPage = 3
  // ----

  bootstrap.Init()

  var drafts = generator.GetAllSortedDrafts(langs)

  if len(drafts) > 0 {

    generator.GenerateAllDraftsLists(drafts, perPage)
  }
}

package main

import (
  "generator"  
  "bootstrap"  
)

func main() {

  bootstrap.Init()

  var config = bootstrap.Config()

  var drafts = generator.GetAllSortedDrafts(config.Json.Pubs.Default.Langs)

  if len(drafts) > 0 {

    generator.GenerateAllDraftsLists(drafts, config.Json.Pubs.Default.PerPage)
  }
}

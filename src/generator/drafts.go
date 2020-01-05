package generator

import (
  "io/ioutil"
  "sort"
  "log"
  "os"
)

// test
func GetAllSortedDrafts(langs []string) []string {

  var drafts []string

  path, _ := os.Getwd()

  for _, lang := range langs {

    var dir = path + "/drafts/" + lang

    var err = os.Chdir(dir)
    CheckError(err)

    files, err := ioutil.ReadDir(dir)
    CheckError(err)

    for _, entry := range files {
      drafts = append(drafts, dir + "/" + entry.Name())
    }
  }

  sort.Strings(drafts)

  return drafts
}

func CheckError(e error) {
  if e != nil {
    log.Fatalf("[!!] Error with opening folder: %s", e)
  }
}

package generator

import (
  "io/ioutil"
  "sort"
  "os"
  "fmt"
  "bytes"
  "text/template"
)

type Drafts struct {
  Drafts []Draft
  CurrentPage int
}

func GetSortedDrafts(langs []string) []string {

  var drafts []string

  path, _ := os.Getwd()

  for _, lang := range langs {

    var dir = path + "/drafts/" + lang

    var err = os.Chdir(dir)
    CheckFatal(6, err)

    files, err := ioutil.ReadDir(dir)
    CheckFatal(7, err)

    for _, entry := range files {
      drafts = append(drafts, dir + "/" + entry.Name())
    }
  }

  sort.Strings(drafts)

  for _, d := range drafts {

    
  }

  return drafts
}

func (ds *Drafts) render() {

  /*
    --- Gather draft content
  */
  draftList, err := ioutil.ReadFile("pubs.html")
  CheckFatal(8, err)

  /*
    --- Merge two of them
  */
  pub, _ := template.New("drafts").Parse(string([]byte(draftList)))

  var draftListContent bytes.Buffer
  pub.Execute(&draftListContent, ds)

  fmt.Println(draftListContent.String())
}

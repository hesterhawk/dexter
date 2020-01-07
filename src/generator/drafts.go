package generator

import (
  "io/ioutil"
  "sort"
  "os"
  "fmt"
  "log"  
  "bytes"
  "helper"
  "strconv"
  "exception"
  "text/template"  
)

type Drafts struct {
  AllPages int
  CurrentPage int  
  Drafts []Draft
}

/*
  --- Get all sorted Drafts - by all languages
*/
func GetAllSortedDrafts(langs []string) []Draft {

  var drafts []Draft

  path, _ := os.Getwd()

  for _, lang := range langs {

    var dir = path + "/drafts/" + lang

    var err = os.Chdir(dir)
    exception.CheckFatal(6, err)

    files, err := ioutil.ReadDir(dir)
    exception.CheckFatal(7, err)

    for _, entry := range files {

      d := Draft{File: entry.Name(), FilePath: dir + "/" + entry.Name(), Lang: lang}
      d.SetHeader()

      drafts = append(drafts, d)
    }
  }

  /*
    --- Sorting articles by file name (so it's a date)
  */
  sort.Slice(drafts[:], func(i, j int) bool {
    return drafts[i].File > drafts[j].File
  })

  return drafts
}

func GenerateDraftsLists(drafts []Draft, perPage int) {

	currentPage := 1
	lenDrafts := len(drafts)

  allPages := lenDrafts / perPage

  if (lenDrafts % perPage) > 0 {
    allPages++
  }

  for from := 0 ; from < lenDrafts ; from += perPage {

		to := 0

		if lenDrafts < (from + perPage) {	

			to = lenDrafts
		} else {

			to = from + perPage
		}

		currentPage++

		dl := Drafts{
      AllPages: allPages,
			CurrentPage: currentPage,
			Drafts: drafts[from:to],
    }
    
    dl.Render()
	}

}

// TODO
func (ds *Drafts) Render() {

  path, _ := os.Getwd()

  /*
    --- Gather draft content
  */

  // TODO - path
  draftList, err := ioutil.ReadFile(path + "/../../templates/pubs.html")
  exception.CheckFatal(8, err)

  /*
    --- Merge two of them
  */
  pub, _ := template.New("drafts").Funcs(template.FuncMap{"N": helper.TmplRange}).Parse(string([]byte(draftList)))

  var draftListContent bytes.Buffer
  pub.Execute(&draftListContent, ds)

  /*
    --- Create Draft Lists file
  */
  helper.CreateFile(path + "/../../bin/page-" + strconv.Itoa(ds.CurrentPage) + ".html", draftListContent.String())

  // DEBUG
  fmt.Println(ds.AllPages)
  log.Fatalf("%s", "dd")
}

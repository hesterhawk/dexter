package generator

import (
  "io/ioutil"
  "sort"
  "os"
  "fmt"
  "bytes"
  "helper"
  "strconv"
  "bootstrap"
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
  
  var config = bootstrap.Config()

  /*
    Gathering drafts from dir and get the header info from the draft file
  */
  for _, lang := range langs {

    var dir = config.AppPath + config.Json.Main.Dirs.Drafts + "/" + lang

    var err = os.Chdir(dir)
    exception.CheckFatal(6, err)

    files, err := ioutil.ReadDir(dir)
    exception.CheckFatal(7, err)

    for _, entry := range files {

      d := Draft{File: entry.Name(), FilePath: dir + "/" + entry.Name(), Lang: lang}

      /*
        Set draft header info to struct
      */
      d.SetHeader()

      /*
        Render draft to the html file
      */
      d.Render()

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

func GenerateAllDraftsLists(drafts []Draft, perPage int) {

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
			CurrentPage: currentPage - 1,
			Drafts: drafts[from:to],
    }
    
    /*
      Render draft list to the html file
    */
    dl.Render()
	}

}


func (ds *Drafts) Render() {

  var config = bootstrap.Config()  

  /*
    --- Gather draft content
  */
  draftList, err := ioutil.ReadFile(config.AppPath + config.Json.Main.Dirs.Templates + "/pubs.html")
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
  helper.CreateFile(config.AppPath + config.Json.Main.Dirs.Bin + "/page-" + strconv.Itoa(ds.CurrentPage) + ".html", draftListContent.String())

  // Info
  fmt.Println("[list] page-" + strconv.Itoa(ds.CurrentPage) + ".html rendered..")
}

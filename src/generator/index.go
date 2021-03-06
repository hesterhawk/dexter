package generator

import (
	"fmt"
	"bytes"
	"helper"	
	"bootstrap"
	"exception"
  "io/ioutil"  
  "text/template"
)

type Index struct {	
	FeaturedDraft Draft
	Drafts []Draft
	DraftsCount int
	DraftsPerPage int
}

func GenerateIndexPage(drafts []Draft) {

	var config = bootstrap.Config()

	var perPage = config.Json.Index.Default.PerPage

	var draftsLen = len(drafts)

	if draftsLen > perPage {
		drafts = drafts[:perPage]
	}

	var index = Index{Drafts:drafts, FeaturedDraft: drafts[0], DraftsCount: draftsLen, DraftsPerPage: perPage}

	index.Render()
}

func (in *Index) Render() {

	var config = bootstrap.Config()

	/*
    --- Gather pub template content
  */
  tmpl, err := ioutil.ReadFile(config.AppPath + config.Json.Main.Dirs.Templates + "/index.html")
  exception.CheckFatal(3, err)

	tmplContent := string([]byte(tmpl))
	
	/*
    --- Merge two of them
  */
  index, _ := template.New("index").Parse(tmplContent)

  var indexContent bytes.Buffer
  index.Execute(&indexContent, in)

  /*
    --- Create pub file
  */
  helper.CreateFile(config.AppPath + config.Json.Main.Dirs.Bin + "/index.html", indexContent.String())

  // Info
  fmt.Println("[index] index.html rendered..")
}
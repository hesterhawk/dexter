package generator

import (
  "os"
  "fmt"
  "bufio"
  "bytes"
  "helper"
  "strconv"  
  "bootstrap" 
  "exception"
  "io/ioutil"  
  "text/template"
)


type Draft struct {
  Title string
  Description string
  CreatedDate string
  Author string
  File string
  FilePath string
  Template string
  Image string
  Content string
  Lang string
  Visible bool
  Update bool
}

func (d* Draft) SetHeader() {

  var result []string

  file, err := os.Open(d.File)
  exception.CheckFatal(1, err)

  defer file.Close()

  scanner := bufio.NewScanner(file)

  for cnt := 0 ; cnt < 8 ; cnt++ {

    scanner.Scan()

    if cnt > 0 {

      result = append(result, scanner.Text())
    }
  }

  d.Title = result[0]
  d.Description = result[1]
  d.CreatedDate = result[2]
  d.Author = result[3]
  d.Image = result[4]
  d.Visible, _ = strconv.ParseBool(result[5])
  d.Update, _ = strconv.ParseBool(result[6])
}


func (d* Draft) Render() {

  var config = bootstrap.Config()

  /*
    --- Gather draft content
  */
  draft, err := ioutil.ReadFile(d.FilePath)
  exception.CheckFatal(2, err)

  d.Content = string([]byte(draft))

  /*
    --- Gather pub template content
  */
  tmpl, err := ioutil.ReadFile(config.AppPath + config.Json.Main.Dirs.Templates + "/pub.html")
  exception.CheckFatal(3, err)

  tmplContent := string([]byte(tmpl))

  /*
    --- Merge two of them
  */
  pub, _ := template.New(d.File).Parse(tmplContent)

  var pubContent bytes.Buffer
  pub.Execute(&pubContent, d)

  /*
    --- Create pub file
  */
  helper.CreateFile(config.AppPath + config.Json.Main.Dirs.Bin + "/" + d.File, pubContent.String())

  // Info
  fmt.Println("[pub] " + d.File + " rendered..")
}

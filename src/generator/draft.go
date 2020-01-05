package generator

import (
  "os"
  "bufio"
  "bytes"
  "strconv"
  "io/ioutil"
  "text/template"
)


type Draft struct {
  Title string
  Description string
  CreatedDate string
  Author string
  File string
  Template string
  Image string
  Content string
  Visible bool
  Update bool
}


func (d* Draft) setHeader() {

  var result []string

  file, err := os.Open(d.File)
  CheckFatal(1, err)

  defer file.Close()

  scanner := bufio.NewScanner(file)

  for cnt := 0 ; cnt < 7 ; cnt++ {

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


func (d* Draft) render(path string) {

  /*
    --- Gather draft content
  */
  draft, err := ioutil.ReadFile(d.File)
  CheckFatal(2, err)

  d.Content = string([]byte(draft))

  /*
    --- Gather pub template content
  */
  tmpl, err := ioutil.ReadFile(d.Template)
  CheckFatal(3, err)

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
  f, err := os.Create(path + "/" + d.File)
  CheckFatal(4, err)

  _, err = f.WriteString(pubContent.String())
  CheckFatal(5, err)

  defer f.Close()
}

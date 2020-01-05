package helper

import (
  "text/template"
  "os"
)

type Inventory struct {
  Material string
  Count    uint
}

func Make(text string) {

  sweaters := Inventory{"wool", 17}
  tmpl, err := template.New("test").Parse(text)

  if err != nil { panic(err) }

  err = tmpl.Execute(os.Stdout, sweaters)

  if err != nil { panic(err) }
}

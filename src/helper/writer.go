package helper

import (
  "os"
  "exception"
)

func CreateFile(path string, content string) {

  f, err := os.Create(path)
  exception.CheckFatal(4, err)

  _, err = f.WriteString(content)
  exception.CheckFatal(5, err)

  defer f.Close()
}
package generator

import "log"

var codes = map[int]string{
  1: "Draft file problem",
  2: "Draft content problem",
  3: "Pub template problem",
  4: "Cannot create file",
  5: "Cannot write to file",
  6: "Drafts folder problem [1]",
  7: "Drafts folder problem [2]",
  8: "Drafts list template problem",
}

func CheckFatal(code int, e error) {

  if e != nil {
    log.Fatalf("[!!] %s: %s", codes[code], e)
  }
}

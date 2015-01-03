package main

import (
	"fmt"
	"github.com/zakuro9715/jingu/parser"
)

func main() {
  var src string
  fmt.Scanln(&src)
  p := parser.Parser{}
  p.Init(src)
  _, errs := p.Parse()

  for _, err := range errs {
    fmt.Println(err)
  }
}

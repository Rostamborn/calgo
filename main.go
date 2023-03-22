package main

import (
	"fmt"
	"log"
	"os"

	"github.com/calgo/parser"
  "github.com/calgo/gui"
)

func main() {
  f, err := os.Open("/home/arman/Coding/golang/calgo/UK_Holidays.ics")
  if err != nil {
    log.Fatal("couldn't open file")
  }
  defer f.Close()

  var p parser.Parser
  pr := p.NewParser(f)

  pr.Parse()

  items := pr.Events  

  for _, v := range items {
    fmt.Println(v)
  }
  gui.Init()
}

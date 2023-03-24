package main

import (
	"fmt"
	"log"

	"github.com/calgo/parser"
  "github.com/calgo/gui"
)

func main() {
  items, err := parser.ReadICS("UK_Holidays.ics")
  if err != nil {
    log.Fatal("couldnt Read ics")
  }

  for _, v := range items {
    fmt.Println(v)
  }
  err, ok := parser.WriteToICS(items, "test.ics")
  if err != nil {
    fmt.Println(err.Error())
    fmt.Println(ok)
  }
  gui.Init()
}

package main

import (
	"fmt"
	"log"

	"github.com/calgo/gui"
	"github.com/calgo/parser"
)

func main() {
	events, err := parser.ReadICS("test.ics")
	if err != nil {
		log.Fatal("couldnt Read ics")
	}

	for _, v := range events {
		fmt.Println(v)
	}
	// err, ok := parser.WriteToICS(items, "test.ics")
	// if err != nil {
	//   fmt.Println(err.Error())
	//   fmt.Println(ok)
	// }
	gui.Initialize(events)
}

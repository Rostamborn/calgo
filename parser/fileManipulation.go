package parser

import (
	"errors"
	"os"
)

func ReadICS(path string) ([]Event, error){
  f, err := os.Open("/home/arman/Coding/golang/calgo/UK_Holidays.ics")
  if err != nil {
    return nil, errors.New("couldn't open file")
  }
  defer f.Close()

  var p Parser
  pr := p.NewParser(f)
  pr.Parse()

  return pr.Events, nil
}

package parser

import (
	"bufio"
	"errors"
	"io"
	"log"
	"strings"
)

func contains(slice []string, pattern string) bool {
  for _, v := range slice {
    if v == pattern {
      return true
    }
  }
  return false
}

type Line struct {
  Key string
  Value string
}

type Event struct {
  Uid string
  Summary string
  DTStart string
  DTEnd string
}

type Parser struct {
  scanner *bufio.Scanner
  Events []Event
}

// Parser constructor
func (p *Parser) NewParser(f io.Reader) *Parser {
  return &Parser{
    scanner: bufio.NewScanner(f),
    Events: make([]Event, 0),
  }
}

// parses each line given to it and returns key value pair
// based on Line structure
func (p *Parser) ParseLine() (*Line, error, bool) {
  isEOF := !p.scanner.Scan()
  l := p.scanner.Text()
  for !isEOF && !strings.Contains(l, ":") {
    isEOF = !p.scanner.Scan()
    l = p.scanner.Text()
  }
  if !isEOF  {
    keyValuePair := strings.Split(l, ":")       
    return &Line{Key: keyValuePair[0], Value: keyValuePair[1]} , nil, false
  }
  return nil, nil, isEOF
}

func (p *Parser) ParseEvent() (*Event, error){
  var l *Line
  var err error
  // a varibale to indicate End Of File
  var isEOF = false
  // a variable to indicate whether we pared END: or not
  var hitEND = false

  var event Event 
  for !isEOF {
    l, err, isEOF = p.ParseLine()
    if err != nil {
      return nil, err
    }

    switch l.Key {
    case "UID":
      event.Uid = l.Value
    case "SUMMARY":
      event.Summary = l.Value
    case "DTSTART":
      event.DTStart = l.Value
    case "DTSTART;VALUE=DATE":
      event.DTStart = l.Value
    case "DTEND":
      event.DTEnd = l.Value
    case "DTEnd;VALUE=DATE":
      event.DTEnd = l.Value
    case "BEGIN":
      return nil, errors.New("Cannot begin new event without ending the previous")
    // we do this trick to end parsing if we reach END:
    case "END":
      isEOF = true
      hitEND = true
    default:
      continue
    }
  }
  if hitEND{
  return &event, nil
  } else {
    return nil, errors.New("Didn't reach END:VEVENT")
  }
}

func (p *Parser) Parse() error{
    // isEOF: a varibale to indicate End Of File
  l, err, isEOF := p.ParseLine()
  for !isEOF {
    if l.Key == "BEGIN" && l.Value == "VEVENT"{
      event, err := p.ParseEvent()
      if err != nil {
        log.Fatal("error while parsing lines")
      } else {
        p.Events = append(p.Events, *event)
      }
    }
    l, err, isEOF = p.ParseLine()
  }
  return err
}

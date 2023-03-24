package parser

import (
	"fmt"
	"os"
)

// writes to a file with "ics" extension. if the file
// doesn't exist it will create one with the path name
func WriteToICS(events []Event, path string) (error, bool) {
  f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
  if err != nil {
    return err, false
  }
  defer f.Close()

  for _, event := range events {
    // this is the format that will be written to a given path
    fmt.Fprintf(f, "BEGIN:VEVENT\nUID:%s\nSUMMARY:%s\nDTSTART:%s\nDTEND%s\nEND:VEVENT\n",
    event.Uid,
    event.Summary,
    event.DTStart,
    event.DTEnd,
    )
  }
  return nil, true
}

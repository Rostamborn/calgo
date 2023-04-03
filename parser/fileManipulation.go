package parser

import (
	"errors"
	"os"
)

func ReadICS(path string) ([]Event, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.New("couldn't open file")
	}
	defer f.Close()

	var p Parser
	pr := p.NewParser(f)
	e := pr.Parse()

	return pr.Events, e
}

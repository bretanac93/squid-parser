package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	s "strings"
)

// LogEntry Doc
type LogEntry struct {
	Timestamp  string
	Elapsed    string
	ClientIP   string
	StatusCode string
	Bytes      string
	Method     string
	URL        string
	User       string
	PeerHost   string
	Mime       string
}

// TransformFile Doc
func TransformFile(rd io.Reader) (r []LogEntry, e error) {
	scanner := bufio.NewScanner(rd)

	for scanner.Scan() {
		line := scanner.Text()
		var pl LogEntry
		pl, e = Parse(line)
		if e == nil {
			r = append(r, pl)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return
}

// Parse Doc
func Parse(l string) (g LogEntry, e error) {
	c := s.Fields(l)
	if len(c) == 10 {
		g = LogEntry{
			Timestamp:  c[0],
			Elapsed:    c[1],
			ClientIP:   c[2],
			StatusCode: c[3],
			Bytes:      c[4],
			Method:     c[5],
			URL:        c[6],
			User:       c[7],
			PeerHost:   c[8],
			Mime:       c[9],
		}
	} else {
		e = fmt.Errorf("%s amount of fields ≠ 10", l)
	}
	return
}

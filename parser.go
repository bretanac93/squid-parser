package squidparser

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
func TransformFile(e io.Reader) (r []LogEntry) {
	scanner := bufio.NewScanner(e)

	for scanner.Scan() {
		line := scanner.Text()
		r = append(r, Parse(line))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return
}

// Parse Doc
func Parse(l string) (e LogEntry) {
	c := s.Fields(l)
	e = LogEntry{
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

	return
}

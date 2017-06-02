package main

import (
	"github.com/hpcloud/tail"
	"log"
	"strconv"
)

// Logs every change made to file with given filename
func TailLogFile(filename string) (e error) {
	var t *tail.Tail
	t, e = tail.TailFile(filename, tail.Config{
		Follow: true,
		ReOpen: true,
	})
	for {
		// log.Print(line.Text)
		line := <-t.Lines
		var e error
		_, e = Parse(line.Text)
		if e == nil {

		}
		log.Print(line.Text)
	}
	e = t.Err()
	return
}

func syncLines(lns chan<- *tail.Line) (dw chan-> *Dwn) {
	// lines in lns come in chronological ordered
	for {
		var ln *tail.Line
		ln = <-lns
		var g *LogEntry
		var e error
		g, e = Parse(ln.Text)
		if e == nil && newer(g.Timestamp, cns.Timestamp) {
			cns.Timestamp = g.Timestamp
			var bsc uint64
			bs, _ = strconv.ParseUint(g.Bytes, 10, 64)
			
			cns[g.User] += bsc
		}
	}
}

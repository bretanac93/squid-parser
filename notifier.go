package main

import (
	"github.com/hpcloud/tail"
	"log"
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

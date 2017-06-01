package squidparser

import (
	"github.com/hpcloud/tail"
	"log"
)

// TailLogFile Doc
func TailLogFile(filename string) {
	t, err := tail.TailFile(filename, tail.Config{
		Follow: true,
		ReOpen: true})

	if err != nil {
		log.Fatal(err)
	}

	for line := range t.Lines {
		// log.Print(line.Text)
		p := Parse(line.Text)
	}
}

package main

import (
	"flag"
	"log"
)

func main() {
	var fn *string
	fn = flag.String("f", "", "Filename to watch")
	flag.Parse()
	var e error
	if *fn != "" {
		e = TailLogFile(*fn)
	}
	if e != nil {
		log.Fatal(e.Error())
	}
}

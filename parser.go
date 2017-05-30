package main

import "fmt"
import s "strings"

type logEntry struct {
	timestamp  string
	elapsed    string
	clientIP   string
	statusCode string
	bytes      string
	method     string
	url        string
	user       string
	peerHost   string
	mime       string
}

func parse(l string) (e logEntry) {
	c := s.Fields(l)
	e = logEntry{
		timestamp:  c[0],
		elapsed:    c[1],
		clientIP:   c[2],
		statusCode: c[3],
		bytes:      c[4],
		method:     c[5],
		url:        c[6],
		user:       c[7],
		peerHost:   c[8],
		mime:       c[9],
	}

	return
}

func main() {
	logLine := "1480792667.325    388 10.2.37.199 TCP_MISS/200 804 POST http://clients1.google.com/ocsp juan.cabo HIER_DIRECT/172.217.2.142 application/ocsp-response"
	entry := parse(logLine)
	fmt.Println(entry)
}
